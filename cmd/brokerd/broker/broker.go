package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
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
)

// Broker creates and tracks request to store Cids in
// the Filecoin network.
type Broker struct {
	store  *srstore.Store
	packer packer.Packer
	piecer piecer.Piecer
}

// New creates a Broker backed by the provdied `ds`.
func New(ds datastore.TxnDatastore, packer packer.Packer, piecer piecer.Piecer) (*Broker, error) {
	store, err := srstore.New(txndswrap.Wrap(ds, "/broker-store"))
	if err != nil {
		return nil, fmt.Errorf("initializing broker request store: %s", err)
	}

	b := &Broker{
		store:  store,
		packer: packer,
		piecer: piecer,
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
		Status:    broker.StatusBatching,
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
	if err := b.packer.ReadyToPack(ctx, br); err != nil {
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
func (b *Broker) CreateStorageDeal(ctx context.Context, srb broker.BrokerRequestGroup) (broker.StorageDeal, error) {
	if !srb.Cid.Defined() {
		return broker.StorageDeal{}, ErrInvalidCid
	}
	if len(srb.GroupedStorageRequests) == 0 {
		return broker.StorageDeal{}, ErrEmptyGroup
	}

	now := time.Now()
	sd := broker.StorageDeal{
		ID:               broker.StorageDealID(uuid.New().String()),
		Cid:              srb.Cid,
		Status:           broker.StorageDealPreparing,
		BrokerRequestIDs: srb.GroupedStorageRequests,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	// Transactionally we:
	// - Move involved BrokerRequest statuses to `Preparing`.
	// - Link each BrokerRequest with the StorageDeal.
	// - Save the `StorageDeal` in the store.
	if err := b.store.CreateStorageDeal(ctx, sd); err != nil {
		return broker.StorageDeal{}, fmt.Errorf("creating storage deal: %w", err)
	}

	// Signal Piecer that there's work to do. It will eventually call us
	// through PreparedStorageDeal(...).
	if err := b.piecer.ReadyToPrepare(ctx, sd.ID, sd.Cid); err != nil {
		// TODO: same possible improvement as described in `ReadyToPack`
		// applies here.
		return broker.StorageDeal{}, fmt.Errorf("signaling piecer: %s", err)
	}

	return sd, nil
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
