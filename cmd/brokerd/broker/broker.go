package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	logger "github.com/ipfs/go-log/v2"
	"github.com/textileio/broker-core/auctioneer"
	"github.com/textileio/broker-core/broker"
	"github.com/textileio/broker-core/cmd/brokerd/srstore"
	"github.com/textileio/broker-core/dshelper/txndswrap"
	"github.com/textileio/broker-core/packer"
	"github.com/textileio/broker-core/piecer"
)

var (
	// ErrNotFound is returned when the broker request doesn't exist.
	ErrNotFound = fmt.Errorf("broker request not found")
	// ErrInvalidCid is returned when the Cid is undefined.
	ErrInvalidCid = fmt.Errorf("the cid can't be undefined")
	// ErrEmptyGroup is returned when an empty storage deal group
	// is received.
	ErrEmptyGroup = fmt.Errorf("the storage deal group is empty")

	log = logger.Logger("broker")
)

// Broker creates and tracks request to store Cids in
// the Filecoin network.
type Broker struct {
	store      *srstore.Store
	packer     packer.Packer
	piecer     piecer.Piecer
	auctioneer auctioneer.Auctioneer
	dealEpochs uint64
}

// New creates a Broker backed by the provided `ds`.
func New(
	ds txndswrap.TxnDatastore,
	packer packer.Packer,
	piecer piecer.Piecer,
	auctioneer auctioneer.Auctioneer,
	dealEpochs uint64,
) (*Broker, error) {
	if dealEpochs < broker.MinDealEpochs {
		return nil, fmt.Errorf("deal epochs is less than minimum allowed: %d", broker.MinDealEpochs)
	}
	if dealEpochs > broker.MaxDealEpochs {
		return nil, fmt.Errorf("deal epochs is greater than maximum allowed: %d", broker.MaxDealEpochs)
	}

	store, err := srstore.New(txndswrap.Wrap(ds, "/broker-store"))
	if err != nil {
		return nil, fmt.Errorf("initializing broker request store: %s", err)
	}

	b := &Broker{
		store:      store,
		packer:     packer,
		piecer:     piecer,
		auctioneer: auctioneer,
		dealEpochs: dealEpochs,
	}
	return b, nil
}

var _ broker.Broker = (*Broker)(nil)

// Create creates a new BrokerRequest with the provided Cid and
// Metadata configuration.
func (b *Broker) Create(ctx context.Context, c cid.Cid, meta broker.Metadata) (broker.BrokerRequest, error) {
	if !c.Defined() {
		return broker.BrokerRequest{}, ErrInvalidCid
	}

	if err := meta.Validate(); err != nil {
		return broker.BrokerRequest{}, fmt.Errorf("invalid metadata: %s", err)
	}

	now := time.Now()
	br := broker.BrokerRequest{
		ID:        broker.BrokerRequestID(uuid.New().String()),
		DataCid:   c,
		Status:    broker.RequestBatching,
		Metadata:  meta,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := b.store.SaveBrokerRequest(ctx, br); err != nil {
		return broker.BrokerRequest{}, fmt.Errorf("saving broker request in store: %s", err)
	}

	// We notify the Packer that this BrokerRequest is ready to be considered.
	// We'll receive a call to `(*Broker).CreateStorageDeal(...)` which will contain
	// this BrokerRequest, and continue with the bidding process..
	if err := b.packer.ReadyToPack(ctx, br.ID, br.DataCid); err != nil {
		// TODO: there's room for improvement here. We can mark this broker-request
		// to be retried in signaling the packer, to avoid be orphaned.
		// Under normal circumstances this shouldn't happen.
		// We can simply save BrokerRequest with a "ReadyToBatch", and have some daemon
		// making sure of notifying and then switching to "Batching"; shoudn't be a big deal.
		return broker.BrokerRequest{}, fmt.Errorf("notifying packer of ready broker request: %s", err)
	}

	return br, nil
}

// Get gets a BrokerRequest by id. If doesn't exist, it returns ErrNotFound.
func (b *Broker) Get(ctx context.Context, ID broker.BrokerRequestID) (broker.BrokerRequest, error) {
	br, err := b.store.GetBrokerRequest(ctx, ID)
	if err == srstore.ErrNotFound {
		return broker.BrokerRequest{}, ErrNotFound
	}
	if err != nil {
		return broker.BrokerRequest{}, fmt.Errorf("get broker request from store: %s", err)
	}

	return br, nil
}

// CreateStorageDeal creates a StorageDeal that contains multiple BrokerRequest. This API is most probably
// caused by Packer. When Packer batches enough pending BrokerRequests in a BrokerRequestGroup, it signals
// the Broker to create a StorageDeal. This StorageDeal should be prepared (piece-size/commP) before publishing
// it in the feed.
func (b *Broker) CreateStorageDeal(
	ctx context.Context,
	batchCid cid.Cid,
	brids []broker.BrokerRequestID) (broker.StorageDealID, error) {
	if !batchCid.Defined() {
		return "", ErrInvalidCid
	}
	if len(brids) == 0 {
		return "", ErrEmptyGroup
	}

	now := time.Now()
	sd := broker.StorageDeal{
		ID:               broker.StorageDealID(uuid.New().String()),
		Cid:              batchCid,
		Status:           broker.StorageDealPreparing,
		BrokerRequestIDs: brids,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	// Transactionally we:
	// - Move involved BrokerRequest statuses to `Preparing`.
	// - Link each BrokerRequest with the StorageDeal.
	// - Save the `StorageDeal` in the store.
	if err := b.store.CreateStorageDeal(ctx, sd); err != nil {
		return "", fmt.Errorf("creating storage deal: %w", err)
	}

	log.Debugf("creating storage deal %s created, signaling piecer...", sd.ID)
	// Signal Piecer that there's work to do. It will eventually call us
	// through PreparedStorageDeal(...).
	if err := b.piecer.ReadyToPrepare(ctx, sd.ID, sd.Cid); err != nil {
		// TODO: same possible improvement as described in `ReadyToPack`
		// applies here.
		return "", fmt.Errorf("signaling piecer: %s", err)
	}

	return sd.ID, nil
}

// StorageDealPrepared is called by Prepared to notify that the data preparation stage is done,
// and to continue with the storage deal process.
func (b *Broker) StorageDealPrepared(
	ctx context.Context,
	id broker.StorageDealID,
	po broker.DataPreparationResult,
) error {
	// TODO: include the data preparation result (piece-size and CommP) in StorageDeal data.
	// @jsign: I'll do this tomorrow.
	// var sd broker.StorageDeal // assume this variable will exist...

	log.Debugf("storage deal %s was prepared, signaling auctioneer...", id)
	// Signal the Auctioneer to create an auction. It will eventually call StorageDealAuctioned(..) to tell
	// us about who won things.
	auctionID, err := b.auctioneer.ReadyToAuction(ctx, id, po.PieceSize, b.dealEpochs)
	if err != nil {
		return fmt.Errorf("signaling auctioneer to create auction: %s", err)
	}

	log.Debugf("created auction %s", auctionID)
	return nil
}

// StorageDealAuctioned is called by the Auctioneer with the result of the StorageDeal auction.
func (b *Broker) StorageDealAuctioned(ctx context.Context, auction broker.Auction) error {
	log.Debugf("storage deal %s was auctioned, signaling dealer...", auction.DealID)

	// TODO: Signal dealer to start the deal.
	return nil
}

// GetStorageDeal gets an existing storage deal. If the storage deal doesn't exists, it returns
// ErrNotFound.
func (b *Broker) GetStorageDeal(ctx context.Context, id broker.StorageDealID) (broker.StorageDeal, error) {
	sd, err := b.store.GetStorageDeal(ctx, id)
	if err == ErrNotFound {
		return broker.StorageDeal{}, ErrNotFound
	}
	if err != nil {
		return broker.StorageDeal{}, fmt.Errorf("get storage deal from store: %s", err)
	}

	return sd, nil
}
