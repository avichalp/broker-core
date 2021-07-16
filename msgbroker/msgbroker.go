package msgbroker

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multiaddr"
	"github.com/textileio/bidbot/lib/auction"
	"github.com/textileio/broker-core/broker"
	"github.com/textileio/broker-core/dealer"
	pb "github.com/textileio/broker-core/gen/broker/v1"
	"google.golang.org/protobuf/proto"
)

// TopicHandler is function that processes a received message.
// If no error is returned, the message will be automatically acked.
// If an error is returned, the message will be automatically nacked.
type TopicHandler func(context.Context, []byte) error

// MsgBroker is a message-broker for async message communication.
type MsgBroker interface {
	// RegisterTopicHandler registers a handler to a topic, with a defined
	// subscription defined by the underlying implementation. Is highly recommended
	// to register handlers in a type-safe way using RegisterHandlers().
	RegisterTopicHandler(topic TopicName, handler TopicHandler, opts ...Option) error

	// PublishMsg publishes a message to the desired topic.
	PublishMsg(ctx context.Context, topicName TopicName, data []byte) error
}

// TopicName is a topic name.
type TopicName string

const (
	// NewBatchCreatedTopic is the topic name for new-batch-created messages.
	NewBatchCreatedTopic TopicName = "new-batch-created"
	// NewBatchPreparedTopic is the topic name for new-batch-prepared messages.
	NewBatchPreparedTopic = "new-batch-prepared"
	// ReadyToBatchTopic is the topic name for ready-to-batch messages.
	ReadyToBatchTopic = "ready-to-batch"
	// ReadyToCreateDealsTopic is the topic name for ready-to-create-deals messages.
	ReadyToCreateDealsTopic = "ready-to-create-deals"
	// FinalizedDealTopic is the topic name for finalized-deal messages.
	FinalizedDealTopic = "finalized-deal"
	// DealProposalAcceptedTopic is the topic name for deal-proposal-accepted messages.
	DealProposalAcceptedTopic = "deal-proposal-accepted"
	// ReadyToAuctionTopic is the topic name for ready-to-auction messages.
	ReadyToAuctionTopic = "ready-to-auction"
)

// NewBatchCreatedListener is a handler for new-batch-created topic.
type NewBatchCreatedListener interface {
	OnNewBatchCreated(context.Context, broker.StorageDealID, cid.Cid, []broker.BrokerRequestID) error
}

// NewBatchPreparedListener is a handler for new-batch-prepared topic.
type NewBatchPreparedListener interface {
	OnNewBatchPrepared(context.Context, broker.StorageDealID, broker.DataPreparationResult) error
}

// ReadyToBatchListener is a handler for ready-to-batch topic.
type ReadyToBatchListener interface {
	OnReadyToBatch(context.Context, []ReadyToBatchData) error
}

// ReadyToBatchData contains broker request data information to be batched.
type ReadyToBatchData struct {
	BrokerRequestID broker.BrokerRequestID
	DataCid         cid.Cid
}

// ReadyToCreateDealsListener is a handler for ready-to-create-deals topic.
type ReadyToCreateDealsListener interface {
	OnReadyToCreateDeals(context.Context, dealer.AuctionDeals) error
}

// FinalizedDealListener is a handler for finalized-deal topic.
type FinalizedDealListener interface {
	OnFinalizedDeal(context.Context, broker.FinalizedDeal) error
}

// DealProposalAcceptedListener is a handler for deal-proposal-accepted topic.
type DealProposalAcceptedListener interface {
	OnDealProposalAccepted(context.Context, auction.AuctionID, auction.BidID, cid.Cid) error
}

// ReadyToAuctionListener is a handler for finalized-deal topic.
type ReadyToAuctionListener interface {
	OnReadyToAuction(
		ctx context.Context,
		id auction.AuctionID,
		storageDealID broker.StorageDealID,
		payloadCid cid.Cid,
		dealSize, dealDuration, dealReplication int,
		dealVerified bool,
		excludedMiners []string,
		filEpochDeadline uint64,
		sources auction.Sources,
	) error
}

// RegisterHandlers automatically calls mb.RegisterTopicHandler in the methods that
// s might satisfy on known XXXListener interfaces. This allows to automatically wire
// s to receive messages from topics of implemented handlers.
func RegisterHandlers(mb MsgBroker, s interface{}, opts ...Option) error {
	var countRegistered int
	if l, ok := s.(NewBatchCreatedListener); ok {
		countRegistered++
		err := mb.RegisterTopicHandler(NewBatchCreatedTopic, func(ctx context.Context, data []byte) error {
			r := &pb.NewBatchCreated{}
			if err := proto.Unmarshal(data, r); err != nil {
				return fmt.Errorf("unmarshal new batch created: %s", err)
			}
			if r.Id == "" {
				return errors.New("storage deal id is empty")
			}
			sdID := broker.StorageDealID(r.Id)

			batchCid, err := cid.Cast(r.BatchCid)
			if err != nil {
				return fmt.Errorf("decoding batch cid: %s", err)
			}
			if !batchCid.Defined() {
				return errors.New("data cid is undefined")
			}
			brids := make([]broker.BrokerRequestID, len(r.BrokerRequestIds))
			for i, id := range r.BrokerRequestIds {
				if id == "" {
					return fmt.Errorf("broker request id can't be empty")
				}
				brids[i] = broker.BrokerRequestID(id)
			}

			if err := l.OnNewBatchCreated(ctx, sdID, batchCid, brids); err != nil {
				return fmt.Errorf("calling on-new-batch-created handler: %s", err)
			}
			return nil
		}, opts...)
		if err != nil {
			return fmt.Errorf("registering handler for new-batch-created topic")
		}
	}

	if l, ok := s.(NewBatchPreparedListener); ok {
		countRegistered++
		err := mb.RegisterTopicHandler(NewBatchPreparedTopic, func(ctx context.Context, data []byte) error {
			r := &pb.NewBatchPrepared{}
			if err := proto.Unmarshal(data, r); err != nil {
				return fmt.Errorf("unmarshal new batch prepared: %s", err)
			}
			pieceCid, err := cid.Cast(r.PieceCid)
			if err != nil {
				return fmt.Errorf("decoding piece cid: %s", err)
			}
			id := broker.StorageDealID(r.Id)
			pr := broker.DataPreparationResult{
				PieceCid:  pieceCid,
				PieceSize: r.PieceSize,
			}
			if err := l.OnNewBatchPrepared(ctx, id, pr); err != nil {
				return fmt.Errorf("calling on-new-batch-prepared handler: %s", err)
			}
			return nil
		}, opts...)
		if err != nil {
			return fmt.Errorf("registering handler for new-batch-prepared topic")
		}
	}

	if l, ok := s.(ReadyToBatchListener); ok {
		countRegistered++
		err := mb.RegisterTopicHandler(ReadyToBatchTopic, func(ctx context.Context, data []byte) error {
			r := &pb.ReadyToBatch{}
			if err := proto.Unmarshal(data, r); err != nil {
				return fmt.Errorf("unmarshal ready to batch: %s", err)
			}
			if len(r.DataCids) == 0 {
				return errors.New("data cids list can't be empty")
			}

			rtb := make([]ReadyToBatchData, len(r.DataCids))
			for i := range r.DataCids {
				if r.DataCids[i].BrokerRequestId == "" {
					return fmt.Errorf("broker request id is empty")
				}
				brID := broker.BrokerRequestID(r.DataCids[i].BrokerRequestId)
				dataCid, err := cid.Cast(r.DataCids[i].DataCid)
				if err != nil {
					return fmt.Errorf("decoding data cid: %s", err)
				}
				rtb[i] = ReadyToBatchData{
					BrokerRequestID: brID,
					DataCid:         dataCid,
				}
			}

			if err := l.OnReadyToBatch(ctx, rtb); err != nil {
				return fmt.Errorf("calling ready-to-batch handler: %s", err)
			}
			return nil
		}, opts...)
		if err != nil {
			return fmt.Errorf("registering handler for ready-to-batch topic")
		}
	}

	if l, ok := s.(ReadyToCreateDealsListener); ok {
		countRegistered++
		err := mb.RegisterTopicHandler(ReadyToCreateDealsTopic, func(ctx context.Context, data []byte) error {
			r := &pb.ReadyToCreateDeals{}
			if err := proto.Unmarshal(data, r); err != nil {
				return fmt.Errorf("unmarshal ready to batch: %s", err)
			}
			if r.StorageDealId == "" {
				return errors.New("storage deal id is empty")
			}

			payloadCid, err := cid.Decode(r.PayloadCid)
			if err != nil {
				return fmt.Errorf("parsing payload cid %s: %s", r.PayloadCid, err)
			}
			pieceCid, err := cid.Decode(r.PieceCid)
			if err != nil {
				return fmt.Errorf("parsing piece cid %s: %s", r.PieceCid, err)
			}

			ads := dealer.AuctionDeals{
				StorageDealID: broker.StorageDealID(r.StorageDealId),
				PayloadCid:    payloadCid,
				PieceCid:      pieceCid,
				PieceSize:     r.PieceSize,
				Duration:      r.Duration,
				Proposals:     make([]dealer.Proposal, len(r.Proposals)),
			}
			for i, t := range r.Proposals {
				if t.Miner == "" {
					return errors.New("miner addr is empty")
				}
				if t.PricePerGibPerEpoch < 0 {
					return errors.New("price per gib per epoch is negative")
				}
				if t.StartEpoch == 0 {
					return errors.New("start epoch should be positive")
				}
				if t.AuctionId == "" {
					return errors.New("auction-id is empty")
				}
				if t.BidId == "" {
					return errors.New("bid-id is empty")
				}
				ads.Proposals[i] = dealer.Proposal{
					Miner:               t.Miner,
					PricePerGiBPerEpoch: t.PricePerGibPerEpoch,
					StartEpoch:          t.StartEpoch,
					Verified:            t.Verified,
					FastRetrieval:       t.FastRetrieval,
					AuctionID:           auction.AuctionID(t.AuctionId),
					BidID:               auction.BidID(t.BidId),
				}
			}
			if err := l.OnReadyToCreateDeals(ctx, ads); err != nil {
				return fmt.Errorf("calling ready-to-create-deals handler: %s", err)
			}
			return nil
		}, opts...)
		if err != nil {
			return fmt.Errorf("registering handler for ready-to-create-deals topic")
		}
	}

	if l, ok := s.(FinalizedDealListener); ok {
		countRegistered++
		err := mb.RegisterTopicHandler(FinalizedDealTopic, func(ctx context.Context, data []byte) error {
			r := &pb.FinalizedDeal{}
			if err := proto.Unmarshal(data, r); err != nil {
				return fmt.Errorf("unmarshal finalized deal msg: %s", err)
			}
			if r.StorageDealId == "" {
				return fmt.Errorf("storage deal id is empty")
			}
			if r.ErrorCause == "" {
				if r.DealId <= 0 {
					return fmt.Errorf(
						"deal id is %d and should be positive",
						r.DealId)
				}
				if r.DealExpiration <= 0 {
					return fmt.Errorf(
						"deal expiration is %d and should be positive",
						r.DealExpiration)
				}
			}
			if r.AuctionId == "" {
				return errors.New("auction-id is empty")
			}
			if r.BidId == "" {
				return errors.New("bid-id is empty")
			}
			fd := broker.FinalizedDeal{
				StorageDealID:  broker.StorageDealID(r.StorageDealId),
				DealID:         r.DealId,
				DealExpiration: r.DealExpiration,
				Miner:          r.MinerId,
				ErrorCause:     r.ErrorCause,
				AuctionID:      auction.AuctionID(r.AuctionId),
				BidID:          auction.BidID(r.BidId),
			}

			if err := l.OnFinalizedDeal(ctx, fd); err != nil {
				return fmt.Errorf("calling finalized-deal handler: %s", err)
			}
			return nil
		}, opts...)
		if err != nil {
			return fmt.Errorf("registering handler for finalized-deal topic")
		}
	}

	if l, ok := s.(DealProposalAcceptedListener); ok {
		countRegistered++
		err := mb.RegisterTopicHandler(DealProposalAcceptedTopic, func(ctx context.Context, data []byte) error {
			r := &pb.DealProposalAccepted{}
			if err := proto.Unmarshal(data, r); err != nil {
				return fmt.Errorf("unmarshal ready to batch: %s", err)
			}
			if r.AuctionId == "" {
				return errors.New("auction id is required")
			}
			if r.BidId == "" {
				return errors.New("bid id is required")
			}
			proposalCid, err := cid.Decode(r.ProposalCid)
			if err != nil || !proposalCid.Defined() {
				return errors.New("invalid proposal cid")
			}

			if err := l.OnDealProposalAccepted(ctx, auction.AuctionID(r.AuctionId), auction.BidID(r.BidId), proposalCid); err != nil {
				return fmt.Errorf("calling deal-proposal-accepted handler: %s", err)
			}
			return nil
		}, opts...)
		if err != nil {
			return fmt.Errorf("registering handler for ready-to-batch topic")
		}
	}

	if l, ok := s.(ReadyToAuctionListener); ok {
		countRegistered++
		err := mb.RegisterTopicHandler(ReadyToAuctionTopic, func(ctx context.Context, data []byte) error {
			req := &pb.ReadyToAuction{}
			if err := proto.Unmarshal(data, req); err != nil {
				return fmt.Errorf("unmarshal finalized deal msg: %s", err)
			}
			if req.Id == "" {
				return errors.New("auction-id is empty")
			}
			if req.StorageDealId == "" {
				return errors.New("storage deal id is empty")
			}
			if req.PayloadCid == "" {
				return errors.New("payload cid is empty")
			}
			payloadCid, err := cid.Parse(req.PayloadCid)
			if err != nil || !payloadCid.Defined() {
				return errors.New("payload cid invalid")
			}
			if req.DealSize == 0 {
				return errors.New("deal size must be greater than zero")
			}
			if req.DealDuration == 0 {
				return errors.New("deal duration must be greater than zero")
			}
			if req.DealReplication == 0 {
				return errors.New("deal replication must be greater than zero")
			}
			sources, err := sourcesFromPb(req.Sources)
			if err != nil {
				return fmt.Errorf("decoding sources: %v", err)
			}

			if err := l.OnReadyToAuction(
				ctx,
				auction.AuctionID(req.Id),
				broker.StorageDealID(req.StorageDealId),
				payloadCid,
				int(req.DealSize),
				int(req.DealDuration),
				int(req.DealReplication),
				req.DealVerified,
				req.ExcludedMiners,
				req.FilEpochDeadline,
				sources,
			); err != nil {
				return fmt.Errorf("calling finalized-deal handler: %s", err)
			}
			return nil
		}, opts...)
		if err != nil {
			return fmt.Errorf("registering handler for finalized-deal topic")
		}
	}

	if countRegistered == 0 {
		return errors.New("no handlers were registered")
	}

	return nil
}

// PublishMsgReadyToBatch publishes a message to the ready-to-batch topic.
func PublishMsgReadyToBatch(ctx context.Context, mb MsgBroker, dataCids []ReadyToBatchData) error {
	if len(dataCids) == 0 {
		return errors.New("data cids is empty")
	}
	msg := &pb.ReadyToBatch{
		DataCids: make([]*pb.ReadyToBatch_ReadyToBatchBR, len(dataCids)),
	}

	for i := range dataCids {
		msg.DataCids[i] = &pb.ReadyToBatch_ReadyToBatchBR{
			BrokerRequestId: string(dataCids[i].BrokerRequestID),
			DataCid:         dataCids[i].DataCid.Bytes(),
		}
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshaling ready-to-batch message: %s", err)
	}
	if err := mb.PublishMsg(ctx, ReadyToBatchTopic, data); err != nil {
		return fmt.Errorf("publishing ready-to-batch message: %s", err)
	}

	return nil
}

// PublishMsgNewBatchCreated publishes a message to the new-batch-created topic.
func PublishMsgNewBatchCreated(
	ctx context.Context,
	mb MsgBroker,
	batchID string,
	batchCid cid.Cid,
	brIDs []broker.BrokerRequestID) error {
	brids := make([]string, len(brIDs))
	for i, bbr := range brIDs {
		brids[i] = string(bbr)
	}
	msg := &pb.NewBatchCreated{
		Id:               batchID,
		BatchCid:         batchCid.Bytes(),
		BrokerRequestIds: brids,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshaling new-batch-created message: %s", err)
	}
	if err := mb.PublishMsg(ctx, NewBatchCreatedTopic, data); err != nil {
		return fmt.Errorf("publishing new-batch-created message: %s", err)
	}

	return nil
}

// PublishMsgNewBatchPrepared publishes a message to the new-batch-prepared topic.
func PublishMsgNewBatchPrepared(
	ctx context.Context,
	mb MsgBroker,
	sdID broker.StorageDealID,
	pieceCid cid.Cid,
	pieceSize uint64) error {
	msg := &pb.NewBatchPrepared{
		Id:        string(sdID),
		PieceCid:  pieceCid.Bytes(),
		PieceSize: pieceSize,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("signaling broker that storage deal is prepared: %s", err)
	}
	if err := mb.PublishMsg(ctx, NewBatchPreparedTopic, data); err != nil {
		return fmt.Errorf("publishing new-prepared-batch message: %s", err)
	}

	return nil
}

// PublishMsgReadyToCreateDeals publishes a message to the ready-to-create-deals topic.
func PublishMsgReadyToCreateDeals(
	ctx context.Context,
	mb MsgBroker,
	ads dealer.AuctionDeals) error {
	msg := &pb.ReadyToCreateDeals{
		StorageDealId: string(ads.StorageDealID),
		PayloadCid:    ads.PayloadCid.String(),
		PieceCid:      ads.PieceCid.String(),
		PieceSize:     ads.PieceSize,
		Duration:      ads.Duration,
		Proposals:     make([]*pb.ReadyToCreateDeals_Proposal, len(ads.Proposals)),
	}
	for i, t := range ads.Proposals {
		if t.Miner == "" {
			return errors.New("miner is empty")
		}
		if t.StartEpoch == 0 {
			return errors.New("start-epoch is zero")
		}
		if t.AuctionID == "" {
			return errors.New("auction-id is empty")
		}
		if t.BidID == "" {
			return errors.New("bid-id is empty")
		}
		msg.Proposals[i] = &pb.ReadyToCreateDeals_Proposal{
			Miner:               t.Miner,
			PricePerGibPerEpoch: t.PricePerGiBPerEpoch,
			StartEpoch:          t.StartEpoch,
			Verified:            t.Verified,
			FastRetrieval:       t.FastRetrieval,
			AuctionId:           string(t.AuctionID),
			BidId:               string(t.BidID),
		}
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("mashaling ready-to-create-deals message: %s", err)
	}
	if err := mb.PublishMsg(ctx, ReadyToCreateDealsTopic, data); err != nil {
		return fmt.Errorf("publishing ready-to-create-deals message: %s", err)
	}

	return nil
}

// PublishMsgFinalizedDeal publishes a message to the finalized-deal topic.
func PublishMsgFinalizedDeal(ctx context.Context, mb MsgBroker, fd broker.FinalizedDeal) error {
	msg := &pb.FinalizedDeal{
		StorageDealId:  string(fd.StorageDealID),
		MinerId:        fd.Miner,
		DealId:         fd.DealID,
		DealExpiration: fd.DealExpiration,
		ErrorCause:     fd.ErrorCause,
		AuctionId:      string(fd.AuctionID),
		BidId:          string(fd.BidID),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("mashaling finalized-deal message: %s", err)
	}
	if err := mb.PublishMsg(ctx, FinalizedDealTopic, data); err != nil {
		return fmt.Errorf("publishing finalized-deal message: %s", err)
	}

	return nil
}

// PublishMsgDealProposalAccepted publishes a message to the deal-proposal-accepted topic.
func PublishMsgDealProposalAccepted(
	ctx context.Context,
	mb MsgBroker,
	sdID broker.StorageDealID,
	auctionID auction.AuctionID,
	bidID auction.BidID,
	miner string,
	propCid cid.Cid) error {
	msg := &pb.DealProposalAccepted{
		StorageDealId: string(sdID),
		Miner:         miner,
		ProposalCid:   propCid.String(),
		AuctionId:     string(auctionID),
		BidId:         string(bidID),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("mashaling deal-proposal-accepted message: %s", err)
	}
	if err := mb.PublishMsg(ctx, DealProposalAcceptedTopic, data); err != nil {
		return fmt.Errorf("publishing deal-proposal-accepted message: %s", err)
	}

	return nil
}

// PublishMsgReadyToAuction publishes a message to the ready-to-auction topic.
func PublishMsgReadyToAuction(
	ctx context.Context,
	mb MsgBroker,
	id auction.AuctionID,
	storageDealID broker.StorageDealID,
	payloadCid cid.Cid,
	dealSize, dealDuration, dealReplication int,
	dealVerified bool,
	excludedMiners []string,
	filEpochDeadline uint64,
	sources auction.Sources) error {
	msg := &pb.ReadyToAuction{
		Id:               string(id),
		StorageDealId:    string(storageDealID),
		PayloadCid:       payloadCid.String(),
		DealSize:         uint64(dealSize),
		DealDuration:     uint64(dealDuration),
		DealReplication:  uint32(dealReplication),
		DealVerified:     dealVerified,
		ExcludedMiners:   excludedMiners,
		FilEpochDeadline: filEpochDeadline,
		Sources:          sourcesToPb(sources),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("mashaling ready-to-auction message: %s", err)
	}
	if err := mb.PublishMsg(ctx, ReadyToAuctionTopic, data); err != nil {
		return fmt.Errorf("publishing ready-to-auction message: %s", err)
	}

	return nil
}

func sourcesToPb(sources auction.Sources) *pb.Sources {
	var carIPFS *pb.Sources_CARIPFS
	if sources.CARIPFS != nil {
		var multiaddrs []string
		for _, addr := range sources.CARIPFS.Multiaddrs {
			multiaddrs = append(multiaddrs, addr.String())
		}
		carIPFS = &pb.Sources_CARIPFS{
			Cid:        sources.CARIPFS.Cid.String(),
			Multiaddrs: multiaddrs,
		}
	}
	var carURL *pb.Sources_CARURL
	if sources.CARURL != nil {
		carURL = &pb.Sources_CARURL{
			URL: sources.CARURL.URL.String(),
		}
	}
	return &pb.Sources{
		CarUrl:  carURL,
		CarIpfs: carIPFS,
	}
}

func sourcesFromPb(pbs *pb.Sources) (sources auction.Sources, err error) {
	if pbs.CarUrl != nil {
		u, err := url.Parse(pbs.CarUrl.URL)
		if err != nil {
			return auction.Sources{}, err
		}
		sources.CARURL = &auction.CARURL{URL: *u}
	}

	if pbs.CarIpfs != nil {
		id, err := cid.Parse(pbs.CarIpfs.Cid)
		if err != nil {
			return auction.Sources{}, err
		}
		var multiaddrs []multiaddr.Multiaddr
		for _, s := range pbs.CarIpfs.Multiaddrs {
			addr, err := multiaddr.NewMultiaddr(s)
			if err != nil {
				return auction.Sources{}, err
			}
			multiaddrs = append(multiaddrs, addr)
		}
		sources.CARIPFS = &auction.CARIPFS{Cid: id, Multiaddrs: multiaddrs}
	}
	return
}
