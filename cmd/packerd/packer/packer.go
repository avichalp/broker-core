package packer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	ipld "github.com/ipfs/go-ipld-format"
	logger "github.com/ipfs/go-log/v2"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/multiformats/go-multibase"
	"github.com/textileio/broker-core/broker"
	"github.com/textileio/broker-core/cmd/packerd/packer/store"
	"github.com/textileio/broker-core/dshelper/txndswrap"
	packeri "github.com/textileio/broker-core/packer"
)

var (
	log = logger.Logger("packer")
)

// Packer provides batching strategies to bundle multiple
// BrokerRequest into a StorageDeal.
type Packer struct {
	batchFrequency time.Duration
	batchLimit     int64

	store  *store.Store
	ipfs   *httpapi.HttpApi
	broker broker.Broker

	onceClose       sync.Once
	daemonCtx       context.Context
	daemonCancelCtx context.CancelFunc
	daemonClosed    chan struct{}
}

var _ packeri.Packer = (*Packer)(nil)

// New returns a new Packer.
func New(
	ds txndswrap.TxnDatastore,
	ipfsClient *httpapi.HttpApi,
	broker broker.Broker,
	opts ...Option) (*Packer, error) {
	cfg := defaultConfig
	for _, op := range opts {
		if err := op(&cfg); err != nil {
			return nil, fmt.Errorf("applying option: %s", err)
		}
	}

	store, err := store.New(txndswrap.Wrap(ds, "/queue"))
	if err != nil {
		return nil, fmt.Errorf("initializing store: %s", err)
	}

	ctx, cls := context.WithCancel(context.Background())
	p := &Packer{
		batchFrequency: cfg.frequency,
		batchLimit:     calcBatchLimit(cfg.sectorSize),

		store:  store,
		ipfs:   ipfsClient,
		broker: broker,

		daemonCtx:       ctx,
		daemonCancelCtx: cls,
		daemonClosed:    make(chan struct{}),
	}

	go p.daemon()

	return p, nil
}

// ReadyToPack signals the packer that there's a new BrokerRequest that can be
// considered. Packer will notify the broker async when the final StorageDeal
// that contains this BrokerRequest gets created.
func (p *Packer) ReadyToPack(ctx context.Context, id broker.BrokerRequestID, dataCid cid.Cid) error {
	if dataCid.Version() != 1 {
		return fmt.Errorf("only cidv1 is supported")
	}

	br := store.BatchableBrokerRequest{
		ID:              "", // Will be set by `store`, explicit here to signal that wasn't missed.
		BrokerRequestID: id,
		DataCid:         dataCid,
	}
	if err := p.store.Enqueue(br); err != nil {
		return fmt.Errorf("enqueueing broker request: %s", err)
	}

	return nil
}

// Close closes the packer.
func (p *Packer) Close() error {
	p.onceClose.Do(func() {
		p.daemonCancelCtx()
		<-p.daemonClosed
	})
	return nil
}

func (p *Packer) daemon() {
	defer close(p.daemonClosed)
	for {
		select {
		case <-p.daemonCtx.Done():
			log.Infof("packer closed")
			return
		case <-time.After(p.batchFrequency):
			if _, err := p.pack(p.daemonCtx); err != nil {
				log.Errorf("packing: %s", err)
			}
		}
	}
}

func (p *Packer) pack(ctx context.Context) (int64, error) {
	batchCid, includedBBRs, numCidsBatched, err := p.batchQueue(ctx)
	if err != nil {
		return 0, fmt.Errorf("batching cids: %s", err)
	}

	// Nothing batched.
	if len(includedBBRs) == 0 {
		return 0, nil
	}

	brids := make([]broker.BrokerRequestID, len(includedBBRs))
	for i, bbr := range includedBBRs {
		brids[i] = bbr.BrokerRequestID
	}
	sdID, err := p.broker.CreateStorageDeal(ctx, batchCid, brids)
	if err != nil {
		return 0, fmt.Errorf("creating storage deal: %s", err)
	}

	if err := p.store.Dequeue(includedBBRs); err != nil {
		return 0, fmt.Errorf("dequeueing batched broker requests: %s", err)
	}

	log.Infof("storage deal created: {id: %s, cid: %s, numCidsBatched: %d}", sdID, batchCid, numCidsBatched)

	return numCidsBatched, nil
}

func (p *Packer) batchQueue(ctx context.Context) (cid.Cid, []store.BatchableBrokerRequest, int64, error) {
	var bbrs []store.BatchableBrokerRequest
	var numBatchedCids int64

	batchRoot := unixfs.EmptyDirNode()
	batchRoot.SetCidBuilder(merkledag.V1CidPrefix())
	batchNodes := map[cid.Cid]*merkledag.ProtoNode{}

	it := p.store.NewIterator()
	for {
		br, ok := it.Next()
		if !ok {
			break
		}

		// 1- See if adding this extra node will go beyond the maximum DAG size.
		targetNode, err := p.ipfs.Dag().Get(ctx, br.DataCid)
		if err != nil {
			return cid.Undef, nil, 0, fmt.Errorf("getting node by cid: %s", err)
		}
		ns, err := targetNode.Stat()
		if err != nil {
			return cid.Undef, nil, 0, fmt.Errorf("getting node stat: %s", err)
		}
		rootStat, err := batchRoot.Stat()
		if err != nil {
			return cid.Undef, nil, 0, fmt.Errorf("getting batch root stats: %s", err)
		}
		currentBatchSize := int64(rootStat.CumulativeSize)

		// If adding this data to the DAG exceeds our limit size for a batch,
		// we skip it and try with the next. Skipped ones still keep original order,
		// so they naturally have priority in the next batch to get in.
		if currentBatchSize+int64(ns.CumulativeSize) > p.batchLimit {
			log.Infof("skipping cid %s which doesn't fit in sector-limit bounds: %s", br.DataCid)
			continue
		}

		// 2- We get a base32 Cid.
		base32Cid, err := br.DataCid.StringOfBase(multibase.Base32)
		if err != nil {
			return cid.Undef, nil, 0, fmt.Errorf("transforming to base32 cid: %s", err)
		}
		base32CidLen := len(base32Cid)

		// 3- Build path through 3 layers.
		layer0LinkName := "ba.." + base32Cid[base32CidLen-2:base32CidLen]
		layer1Node, err := getOrCreateLayerNode(batchRoot, layer0LinkName, batchNodes)
		if err != nil {
			return cid.Undef, nil, 0, fmt.Errorf("get/create layer0 node: %s", err)
		}
		layer1LinkName := "ba.." + base32Cid[base32CidLen-4:base32CidLen]
		layer2Node, err := getOrCreateLayerNode(layer1Node, layer1LinkName, batchNodes)
		if err != nil {
			return cid.Undef, nil, 0, fmt.Errorf("get/create layer1 node: %s", err)
		}

		// 4- Add the target Cid in the layer 3 node, and bubble up parent nodes
		//    changes up to the updated batchRoot that now includes this cid.
		_, err = layer2Node.GetNodeLink(base32Cid)
		if err != nil && err != merkledag.ErrLinkNotFound {
			return cid.Undef, nil, 0, fmt.Errorf("getting target node: %s", err)
		}
		if err == merkledag.ErrLinkNotFound {
			if err := updateNodeLink(layer2Node, base32Cid, targetNode, batchNodes); err != nil {
				return cid.Undef, nil, 0, fmt.Errorf("adding target node link: %s", err)
			}
			if err := updateNodeLink(layer1Node, layer1LinkName, layer2Node, batchNodes); err != nil {
				return cid.Undef, nil, 0, fmt.Errorf("updating layer 2 node: %s", err)
			}
			if err := updateNodeLink(batchRoot, layer0LinkName, layer1Node, batchNodes); err != nil {
				return cid.Undef, nil, 0, fmt.Errorf("updating layer 1 node: %s", err)
			}
			numBatchedCids++
		} else {
			// This "else" case is interesting. It means we already batched the same dataCid
			// in the dag. Other BrokerRequest had the same Cid!
			// That's to say, we don't do anything new in this iteration.
			// This isn't a problem, both StorageRequests are fulfilled and can find their data.

			// Let's log a warning just to be aware that this is a lucky situation and not
			// something firing often.
			log.Warnf("lucky! cid %s is already in the batch", base32Cid)
		}

		// We include the BrokerRequest as part of this batch.
		bbrs = append(bbrs, br)
	}

	toBeAddedNodes := make([]ipld.Node, 0, len(batchNodes))
	for _, node := range batchNodes {
		toBeAddedNodes = append(toBeAddedNodes, node)
	}
	if err := p.ipfs.Dag().AddMany(ctx, toBeAddedNodes); err != nil {
		return cid.Undef, nil, 0, fmt.Errorf("adding updated nodes: %s", err)
	}

	if err := p.ipfs.Pin().Add(ctx, path.IpfsPath(batchRoot.Cid()), options.Pin.Recursive(true)); err != nil {
		return cid.Undef, nil, 0, fmt.Errorf("pinning batch root: %s", err)
	}

	return batchRoot.Cid(), bbrs, numBatchedCids, nil
}

func getOrCreateLayerNode(
	parentNode *merkledag.ProtoNode,
	layerLinkName string,
	batchNodes map[cid.Cid]*merkledag.ProtoNode) (*merkledag.ProtoNode, error) {
	layerLink, err := parentNode.GetNodeLink(layerLinkName)
	if err != nil && err != merkledag.ErrLinkNotFound {
		return nil, fmt.Errorf("getting node by name: %s", err)
	}
	var layerNode *merkledag.ProtoNode
	if err == merkledag.ErrLinkNotFound {
		// We'll edit the paretNode to add an extra link. We can remove the current parentNode
		// from the batch nodes set, since it won't be used in the final DAG.
		delete(batchNodes, parentNode.Cid())

		layerNode = unixfs.EmptyDirNode()
		if err := parentNode.AddNodeLink(layerLinkName, layerNode); err != nil {
			return nil, fmt.Errorf("adding link: %s", err)
		}

		// Add again parentNode with the extra link.
		batchNodes[parentNode.Cid()] = parentNode
	} else {
		l1Node, ok := batchNodes[layerLink.Cid]
		if !ok {
			return nil, fmt.Errorf("proto node not found in batch nodes, this should never happen")
		}
		layerNode = l1Node
	}

	return layerNode, nil
}

func updateNodeLink(
	node *merkledag.ProtoNode,
	linkName string,
	newNode ipld.Node,
	batchNodes map[cid.Cid]*merkledag.ProtoNode) error {
	delete(batchNodes, node.Cid())
	if err := node.RemoveNodeLink(linkName); err != nil && err != merkledag.ErrLinkNotFound {
		return fmt.Errorf("removing link name: %s", err)
	}
	if err := node.AddNodeLink(linkName, newNode); err != nil {
		return fmt.Errorf("updating batch root node link: %s", err)
	}
	batchNodes[node.Cid()] = node

	return nil
}

// calcBatchLimit does a worst-case estimation of the maximum size the batch DAG can have to fit
// in the specified sector size.
// Rationale of the calculation:
// - Mul 127/128: Account fr32 padding in the sector.
// - Div 1.04: This is an estimation of CAR serialization overhead.
//             2(varint)+34(cid-size)+~1024(small-block) -> ~4% overhead.
//
// Note that this limit is a worst-case scenario for a perfectly calculated size for a DAG,
// that's to say, accounting for deduplication. TL;DR: safe boundaries.
func calcBatchLimit(sectorSize int64) int64 {
	return int64(float64(sectorSize) * float64(127) / float64(128) / 1.04)
}