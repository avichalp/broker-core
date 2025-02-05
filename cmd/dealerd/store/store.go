package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/jackc/pgconn"
	"github.com/textileio/broker-core/broker"
	"github.com/textileio/broker-core/cmd/dealerd/store/internal/db"
	"github.com/textileio/broker-core/cmd/dealerd/store/migrations"
	"github.com/textileio/broker-core/storeutil"
	logger "github.com/textileio/go-log/v2"
)

const (
	// StatusDealMaking indicates that an auction data is ready for deal making.
	StatusDealMaking = AuctionDealStatus(db.StatusDealMaking)
	// StatusConfirmation indicates that a deal is pending to be confirmed.
	StatusConfirmation = AuctionDealStatus(db.StatusConfirmation)
	// StatusReportFinalized indicates that a deal result is pending to be reported.
	StatusReportFinalized = AuctionDealStatus(db.StatusReportFinalized)
	// StatusFinalized indicates that a deal result has been reported and all processing is complete.
	StatusFinalized = AuctionDealStatus(db.StatusFinalized)

	stuckSeconds = int64(300)
)

var (
	log = logger.Logger("store")

	// ErrNotFound is returned if the item isn't found in the store.
	ErrNotFound = errors.New("id not found")

	// ErrAuctionDataExists indicates that the auction data already exists.
	ErrAuctionDataExists = errors.New("auction data already exists")
)

// AuctionDealStatus is the type of action deal status.
type AuctionDealStatus db.Status

// AuctionData contains information of data to be stored in Filecoin.
type AuctionData struct {
	ID         string
	BatchID    broker.BatchID
	PayloadCid cid.Cid
	PieceCid   cid.Cid
	PieceSize  uint64
	Duration   uint64
	CreatedAt  time.Time
	CARURL     string
}

// AuctionDeal contains information to make a deal with a particular
// storage-provider. The data information is stored in the linked AuctionData.
type AuctionDeal db.AuctionDeal

// RemoteWallet contains configuration of a remote wallet.
type RemoteWallet = db.RemoteWallet

// Store provides persistent storage for Bids.
type Store struct {
	conn      *sql.DB
	db        *db.Queries
	statusMap map[string]storagemarket.StorageDealStatus
}

// New returns a *Store.
func New(postgresURI string) (*Store, error) {
	as := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})
	conn, err := storeutil.MigrateAndConnectToDB(postgresURI, as)
	if err != nil {
		return nil, err
	}
	s := &Store{conn: conn, db: db.New(conn)}
	return s, s.initStatusMap()
}

func (s *Store) useTxFromCtx(ctx context.Context, f func(*db.Queries) error) (err error) {
	return storeutil.WithCtxTx(ctx,
		func(tx *sql.Tx) error { return f(s.db.WithTx(tx)) },
		func() error { return f(s.db) })
}

// Create persist new auction data and a set of related auction deals. It
// writes the IDs back to the parameters passed in.
func (s *Store) Create(ctx context.Context, ad *AuctionData, ads []*AuctionDeal, rw *broker.RemoteWallet) error {
	err := validate(ad, ads, rw)
	if err != nil {
		return fmt.Errorf("invalid auction data: %s", err)
	}

	return storeutil.WithTx(ctx, s.conn, func(tx *sql.Tx) error {
		txn := s.db.WithTx(tx)
		if err := txn.CreateAuctionData(ctx, db.CreateAuctionDataParams{
			ID:         ad.ID,
			BatchID:    ad.BatchID,
			PayloadCid: ad.PayloadCid.String(),
			PieceCid:   ad.PieceCid.String(),
			PieceSize:  ad.PieceSize,
			Duration:   ad.Duration,
			CarUrl:     ad.CARURL,
		}); err != nil {
			if err, ok := err.(*pgconn.PgError); ok {
				if err.Code == "23505" {
					return ErrAuctionDataExists
				}
			}

			return fmt.Errorf("saving auction data in db: %s", err)
		}
		for _, aud := range ads {
			aud.ID = uuid.New().String()
			aud.AuctionDataID = ad.ID
			aud.BatchID = ad.BatchID
			aud.Status = db.StatusDealMaking
			aud.ReadyAt = time.Unix(0, 0)
			aud.MarketDealStatus = "UNKNOWN"
			if err := txn.CreateAuctionDeal(ctx, db.CreateAuctionDealParams{
				ID:                  aud.ID,
				AuctionDataID:       aud.AuctionDataID,
				BatchID:             aud.BatchID,
				StorageProviderID:   aud.StorageProviderID,
				PricePerGibPerEpoch: aud.PricePerGibPerEpoch,
				StartEpoch:          aud.StartEpoch,
				Verified:            aud.Verified,
				FastRetrieval:       aud.FastRetrieval,
				AuctionID:           aud.AuctionID,
				BidID:               aud.BidID,
				Status:              aud.Status,
				Executing:           aud.Executing,
				ErrorCause:          aud.ErrorCause,
				Retries:             aud.Retries,
				ProposalCid:         aud.ProposalCid,
				DealID:              aud.DealID,
				DealExpiration:      aud.DealExpiration,
				MarketDealStatus:    aud.MarketDealStatus,
				ReadyAt:             aud.ReadyAt,
			}); err != nil {
				return fmt.Errorf("saving auction deal in datastore: %s", err)
			}
		}
		if rw != nil {
			maddrs := make([]string, len(rw.Multiaddrs))
			for i, maddr := range rw.Multiaddrs {
				maddrs[i] = maddr.String()
			}
			if err := txn.CreateRemoteWallet(ctx, db.CreateRemoteWalletParams{
				AuctionDataID: ad.ID,
				PeerID:        rw.PeerID.String(),
				AuthToken:     rw.AuthToken,
				WalletAddr:    rw.WalletAddr.String(),
				Multiaddrs:    maddrs,
			}); err != nil {
				return fmt.Errorf("saving remote wallet config in datastore: %s", err)
			}
		}

		return nil
	})
}

// GetNextPending returns the next auction deal in the provided status and marks
// the deal as executing. If none exists, it returns false in the second parameter.
func (s *Store) GetNextPending(ctx context.Context, status AuctionDealStatus) (ad AuctionDeal, exists bool, err error) {
	err = s.useTxFromCtx(ctx, func(q *db.Queries) error {
		params := db.NextPendingAuctionDealParams{
			Status:       db.Status(status),
			StuckSeconds: stuckSeconds,
		}
		deal, err := q.NextPendingAuctionDeal(ctx, params)
		if err == nil {
			if int64(time.Since(deal.ReadyAt).Seconds()) > stuckSeconds {
				log.Warnf("re-executing stuck auction deal %s", deal.ID)
			}
			ad = AuctionDeal(deal)
		}
		return err
	})
	if err == sql.ErrNoRows {
		return AuctionDeal{}, false, nil
	} else if err != nil {
		return
	}
	return ad, true, nil
}

// SaveAndMoveAuctionDeal persists a modified auction deal, moves it to the
// provided status and marks it as not executing.
func (s *Store) SaveAndMoveAuctionDeal(ctx context.Context, aud AuctionDeal, newStatus AuctionDealStatus) error {
	if err := areAllExpectedFieldsSet(aud); err != nil {
		return fmt.Errorf("validating that expected fields are set for new status: %s", err)
	}
	return s.useTxFromCtx(ctx, func(q *db.Queries) error {
		rows, err := q.UpdateAuctionDeal(ctx, db.UpdateAuctionDealParams{
			ID:                  aud.ID,
			AuctionDataID:       aud.AuctionDataID,
			BatchID:             aud.BatchID,
			StorageProviderID:   aud.StorageProviderID,
			PricePerGibPerEpoch: aud.PricePerGibPerEpoch,
			StartEpoch:          aud.StartEpoch,
			Verified:            aud.Verified,
			FastRetrieval:       aud.FastRetrieval,
			AuctionID:           aud.AuctionID,
			BidID:               aud.BidID,

			Status:     db.Status(newStatus),
			Executing:  false,
			ErrorCause: aud.ErrorCause,
			Retries:    aud.Retries,

			ProposalCid:      aud.ProposalCid,
			DealID:           aud.DealID,
			DealExpiration:   aud.DealExpiration,
			MarketDealStatus: aud.MarketDealStatus,
			ReadyAt:          aud.ReadyAt,
		})
		if err != nil {
			return err
		}
		if rows == 0 {
			return ErrNotFound
		}
		return nil
	})
}

// GetAuctionData returns an auction data by id.
func (s *Store) GetAuctionData(ctx context.Context, auctionDataID string) (ad AuctionData, err error) {
	err = s.useTxFromCtx(ctx, func(q *db.Queries) error {
		datum, err := q.GetAuctionData(ctx, auctionDataID)
		if err == nil {
			payloadCid, err := cid.Parse(datum.PayloadCid)
			if err != nil {
				return fmt.Errorf("parsing payload cid from db: %s", err)
			}
			pieceCid, err := cid.Parse(datum.PieceCid)
			if err != nil {
				return fmt.Errorf("parsing piece cid from db: %s", err)
			}
			ad = AuctionData{
				ID:         datum.ID,
				BatchID:    datum.BatchID,
				PayloadCid: payloadCid,
				PieceCid:   pieceCid,
				PieceSize:  datum.PieceSize,
				Duration:   datum.Duration,
				CreatedAt:  datum.CreatedAt,
				CARURL:     datum.CarUrl,
			}
		}
		return err
	})
	return
}

// IsBoostAllowed returns true if Boost proposals are allowed for the storaeg provider.
func (s *Store) IsBoostAllowed(ctx context.Context, spID string) (ok bool, err error) {
	err = s.useTxFromCtx(ctx, func(q *db.Queries) error {
		_, err := q.IsBoostAllowed(ctx, spID)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("get boost whitelist row: %s", err)
		}
		ok = true
		return nil
	})
	return
}

// GetRemoteWallet returns the remote wallet configuration for an auction if exists.
// If the auction deal wasn't configured with a remote wallet, *it will return nil without an error*.
func (s *Store) GetRemoteWallet(ctx context.Context, auctionDataID string) (rw *RemoteWallet, err error) {
	err = s.useTxFromCtx(ctx, func(q *db.Queries) error {
		remoteWallet, err := q.GetRemoteWallet(ctx, auctionDataID)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("get remote wallet: %s", err)
		}
		rw = &remoteWallet
		return nil
	})
	return
}

// GetStatusCounts returns a summary of in which deal statuses are watched deals.
func (s *Store) GetStatusCounts(ctx context.Context) (map[storagemarket.StorageDealStatus]int64, error) {
	ads, err := s.db.GetAuctionDealsByStatus(ctx, db.StatusConfirmation)
	if err != nil {
		return nil, fmt.Errorf("getting auction deals: %s", err)
	}
	ret := map[storagemarket.StorageDealStatus]int64{}
	for _, ad := range ads {
		st, ok := s.statusMap[ad.MarketDealStatus]
		if !ok {
			return nil, fmt.Errorf("unknown status: %s", ad.MarketDealStatus)
		}
		ret[st]++
	}
	return ret, nil
}

// GetStatusForStatusID returns the string enum representation for the given storage deal status.
func (s *Store) GetStatusForStatusID(
	ctx context.Context,
	id storagemarket.StorageDealStatus,
) (status string, ok bool, err error) {
	var mds db.MarketDealStatus
	mds, err = s.db.GetMarketDealStatusForID(ctx, int64(id))
	if err == sql.ErrNoRows {
		return "", false, nil
	}
	if err != nil {
		return status, ok, fmt.Errorf("getting market deal status for id: %s", err)
	}
	return mds.Type, true, nil
}

func validate(ad *AuctionData, ads []*AuctionDeal, rw *broker.RemoteWallet) error {
	if ad.Duration <= 0 {
		return fmt.Errorf("invalid duration: %d", ad.Duration)
	}
	if ad.ID == "" {
		return errors.New("id is empty")
	}
	if ad.BatchID == "" {
		return errors.New("batch id is empty")
	}
	if !ad.PayloadCid.Defined() {
		return errors.New("invalid payload cid")
	}
	if !ad.PieceCid.Defined() {
		return errors.New("invalid piece cid")
	}
	if ad.PieceSize <= 0 {
		return errors.New("piece size is zero")
	}

	for _, auctionDeal := range ads {
		if auctionDeal.StorageProviderID == "" {
			return errors.New("storage-provider address is empty")
		}
		if auctionDeal.PricePerGibPerEpoch < 0 {
			return errors.New("price-per-gib-per-epoch is negative")
		}
		if auctionDeal.StartEpoch <= 0 {
			return errors.New("start-epoch isn't positive")
		}
	}

	if rw != nil {
		if err := rw.PeerID.Validate(); err != nil {
			return fmt.Errorf("remote wallet invalid peer-id: %s", err)
		}
		if rw.AuthToken == "" {
			return errors.New("remote wallet auth token is empty")
		}
		if rw.WalletAddr.Empty() {
			return errors.New("remote wallet wallet address is empty")
		}
	}

	return nil
}

func areAllExpectedFieldsSet(ad AuctionDeal) error {
	switch ad.Status {
	case db.StatusDealMaking:
		// Nothing to validate.
	case db.StatusConfirmation:
		_, errProposalCid := cid.Parse(ad.ProposalCid)
		_, errDealUID := uuid.Parse(ad.ProposalCid)
		if errProposalCid != nil && errDealUID != nil {
			return fmt.Errorf("proposalcid or dealUID (%s) should be set to transition to WaitingConfirmation", ad.ProposalCid)
		}
	case db.StatusReportFinalized:
		if ad.Executing {
			if ad.ErrorCause == "" { // Success
				if ad.DealID == 0 {
					return errors.New("a success status should have a defined deal-id")
				}
				if ad.DealExpiration == 0 {
					return errors.New("a success status should have a defined deal-expiration")
				}
				if ad.ErrorCause != "" {
					return fmt.Errorf("a success status can't have an error cause: %s", ad.ErrorCause)
				}
			}
			return nil
		}
		fallthrough
	default:
		return fmt.Errorf("unknown status: %s", ad.Status) // and executing?
	}
	return nil
}

// getAllPending is a method only to be used in tests.
func (s *Store) getAllPending() (ret []AuctionDeal, err error) {
	ads, err := s.db.GetAuctionDealsByStatus(context.Background(), db.StatusDealMaking)
	if err != nil {
		return nil, fmt.Errorf("get auction deals: %s", err)
	}
	for _, ad := range ads {
		if !ad.Executing {
			ret = append(ret, AuctionDeal(ad))
		}
	}
	return
}

func (s *Store) initStatusMap() error {
	s.statusMap = make(map[string]uint64)
	statuses, err := s.db.GetAllMarketDealStatuses(context.Background())
	if err != nil {
		return fmt.Errorf("querying for all market deal statuses: %s", err)
	}
	for _, status := range statuses {
		s.statusMap[status.Type] = uint64(status.ID)
	}
	return nil
}
