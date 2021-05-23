// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: broker/auctioneer/v1/auctioneer.proto

package auctioneer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Auction_Status int32

const (
	Auction_STATUS_UNSPECIFIED Auction_Status = 0
	Auction_STATUS_QUEUED      Auction_Status = 1
	Auction_STATUS_STARTED     Auction_Status = 2
	Auction_STATUS_ENDED       Auction_Status = 3
	Auction_STATUS_ERROR       Auction_Status = 4
)

// Enum value maps for Auction_Status.
var (
	Auction_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_QUEUED",
		2: "STATUS_STARTED",
		3: "STATUS_ENDED",
		4: "STATUS_ERROR",
	}
	Auction_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_QUEUED":      1,
		"STATUS_STARTED":     2,
		"STATUS_ENDED":       3,
		"STATUS_ERROR":       4,
	}
)

func (x Auction_Status) Enum() *Auction_Status {
	p := new(Auction_Status)
	*p = x
	return p
}

func (x Auction_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Auction_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_broker_auctioneer_v1_auctioneer_proto_enumTypes[0].Descriptor()
}

func (Auction_Status) Type() protoreflect.EnumType {
	return &file_broker_auctioneer_v1_auctioneer_proto_enumTypes[0]
}

func (x Auction_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Auction_Status.Descriptor instead.
func (Auction_Status) EnumDescriptor() ([]byte, []int) {
	return file_broker_auctioneer_v1_auctioneer_proto_rawDescGZIP(), []int{0, 0}
}

type Auction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StorageDealId   string                  `protobuf:"bytes,2,opt,name=storage_deal_id,json=storageDealId,proto3" json:"storage_deal_id,omitempty"`
	DealSize        uint64                  `protobuf:"varint,3,opt,name=deal_size,json=dealSize,proto3" json:"deal_size,omitempty"`
	DealDuration    uint64                  `protobuf:"varint,4,opt,name=deal_duration,json=dealDuration,proto3" json:"deal_duration,omitempty"`
	DealReplication uint32                  `protobuf:"varint,5,opt,name=deal_replication,json=dealReplication,proto3" json:"deal_replication,omitempty"`
	DealVerified    bool                    `protobuf:"varint,6,opt,name=deal_verified,json=dealVerified,proto3" json:"deal_verified,omitempty"`
	Status          Auction_Status          `protobuf:"varint,7,opt,name=status,proto3,enum=broker.auctioneer.v1.Auction_Status" json:"status,omitempty"`
	Bids            map[string]*Auction_Bid `protobuf:"bytes,8,rep,name=bids,proto3" json:"bids,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WinningBids     []string                `protobuf:"bytes,9,rep,name=winning_bids,json=winningBids,proto3" json:"winning_bids,omitempty"`
	StartedAt       *timestamppb.Timestamp  `protobuf:"bytes,10,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	Duration        int64                   `protobuf:"varint,11,opt,name=duration,proto3" json:"duration,omitempty"`
	Attempts        uint32                  `protobuf:"varint,12,opt,name=attempts,proto3" json:"attempts,omitempty"`
	Error           string                  `protobuf:"bytes,13,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Auction) Reset() {
	*x = Auction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction) ProtoMessage() {}

func (x *Auction) ProtoReflect() protoreflect.Message {
	mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auction.ProtoReflect.Descriptor instead.
func (*Auction) Descriptor() ([]byte, []int) {
	return file_broker_auctioneer_v1_auctioneer_proto_rawDescGZIP(), []int{0}
}

func (x *Auction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Auction) GetStorageDealId() string {
	if x != nil {
		return x.StorageDealId
	}
	return ""
}

func (x *Auction) GetDealSize() uint64 {
	if x != nil {
		return x.DealSize
	}
	return 0
}

func (x *Auction) GetDealDuration() uint64 {
	if x != nil {
		return x.DealDuration
	}
	return 0
}

func (x *Auction) GetDealReplication() uint32 {
	if x != nil {
		return x.DealReplication
	}
	return 0
}

func (x *Auction) GetDealVerified() bool {
	if x != nil {
		return x.DealVerified
	}
	return false
}

func (x *Auction) GetStatus() Auction_Status {
	if x != nil {
		return x.Status
	}
	return Auction_STATUS_UNSPECIFIED
}

func (x *Auction) GetBids() map[string]*Auction_Bid {
	if x != nil {
		return x.Bids
	}
	return nil
}

func (x *Auction) GetWinningBids() []string {
	if x != nil {
		return x.WinningBids
	}
	return nil
}

func (x *Auction) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Auction) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Auction) GetAttempts() uint32 {
	if x != nil {
		return x.Attempts
	}
	return 0
}

func (x *Auction) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ReadyToAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageDealId   string `protobuf:"bytes,1,opt,name=storage_deal_id,json=storageDealId,proto3" json:"storage_deal_id,omitempty"`
	DealSize        uint64 `protobuf:"varint,2,opt,name=deal_size,json=dealSize,proto3" json:"deal_size,omitempty"`
	DealDuration    uint64 `protobuf:"varint,3,opt,name=deal_duration,json=dealDuration,proto3" json:"deal_duration,omitempty"`
	DealReplication uint32 `protobuf:"varint,4,opt,name=deal_replication,json=dealReplication,proto3" json:"deal_replication,omitempty"`
	DealVerified    bool   `protobuf:"varint,5,opt,name=deal_verified,json=dealVerified,proto3" json:"deal_verified,omitempty"`
}

func (x *ReadyToAuctionRequest) Reset() {
	*x = ReadyToAuctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyToAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyToAuctionRequest) ProtoMessage() {}

func (x *ReadyToAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyToAuctionRequest.ProtoReflect.Descriptor instead.
func (*ReadyToAuctionRequest) Descriptor() ([]byte, []int) {
	return file_broker_auctioneer_v1_auctioneer_proto_rawDescGZIP(), []int{1}
}

func (x *ReadyToAuctionRequest) GetStorageDealId() string {
	if x != nil {
		return x.StorageDealId
	}
	return ""
}

func (x *ReadyToAuctionRequest) GetDealSize() uint64 {
	if x != nil {
		return x.DealSize
	}
	return 0
}

func (x *ReadyToAuctionRequest) GetDealDuration() uint64 {
	if x != nil {
		return x.DealDuration
	}
	return 0
}

func (x *ReadyToAuctionRequest) GetDealReplication() uint32 {
	if x != nil {
		return x.DealReplication
	}
	return 0
}

func (x *ReadyToAuctionRequest) GetDealVerified() bool {
	if x != nil {
		return x.DealVerified
	}
	return false
}

type ReadyToAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadyToAuctionResponse) Reset() {
	*x = ReadyToAuctionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyToAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyToAuctionResponse) ProtoMessage() {}

func (x *ReadyToAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyToAuctionResponse.ProtoReflect.Descriptor instead.
func (*ReadyToAuctionResponse) Descriptor() ([]byte, []int) {
	return file_broker_auctioneer_v1_auctioneer_proto_rawDescGZIP(), []int{2}
}

func (x *ReadyToAuctionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAuctionRequest) Reset() {
	*x = GetAuctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuctionRequest) ProtoMessage() {}

func (x *GetAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuctionRequest.ProtoReflect.Descriptor instead.
func (*GetAuctionRequest) Descriptor() ([]byte, []int) {
	return file_broker_auctioneer_v1_auctioneer_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuctionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auction *Auction `protobuf:"bytes,1,opt,name=auction,proto3" json:"auction,omitempty"`
}

func (x *GetAuctionResponse) Reset() {
	*x = GetAuctionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuctionResponse) ProtoMessage() {}

func (x *GetAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuctionResponse.ProtoReflect.Descriptor instead.
func (*GetAuctionResponse) Descriptor() ([]byte, []int) {
	return file_broker_auctioneer_v1_auctioneer_proto_rawDescGZIP(), []int{4}
}

func (x *GetAuctionResponse) GetAuction() *Auction {
	if x != nil {
		return x.Auction
	}
	return nil
}

type Auction_Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinerAddr        string                 `protobuf:"bytes,1,opt,name=miner_addr,json=minerAddr,proto3" json:"miner_addr,omitempty"`
	WalletAddrSig    []byte                 `protobuf:"bytes,2,opt,name=wallet_addr_sig,json=walletAddrSig,proto3" json:"wallet_addr_sig,omitempty"`
	BidderId         string                 `protobuf:"bytes,3,opt,name=bidder_id,json=bidderId,proto3" json:"bidder_id,omitempty"`
	AskPrice         int64                  `protobuf:"varint,4,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	VerifiedAskPrice int64                  `protobuf:"varint,5,opt,name=verified_ask_price,json=verifiedAskPrice,proto3" json:"verified_ask_price,omitempty"`
	StartEpoch       uint64                 `protobuf:"varint,6,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	FastRetrieval    bool                   `protobuf:"varint,7,opt,name=fast_retrieval,json=fastRetrieval,proto3" json:"fast_retrieval,omitempty"`
	ReceivedAt       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
}

func (x *Auction_Bid) Reset() {
	*x = Auction_Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auction_Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction_Bid) ProtoMessage() {}

func (x *Auction_Bid) ProtoReflect() protoreflect.Message {
	mi := &file_broker_auctioneer_v1_auctioneer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auction_Bid.ProtoReflect.Descriptor instead.
func (*Auction_Bid) Descriptor() ([]byte, []int) {
	return file_broker_auctioneer_v1_auctioneer_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Auction_Bid) GetMinerAddr() string {
	if x != nil {
		return x.MinerAddr
	}
	return ""
}

func (x *Auction_Bid) GetWalletAddrSig() []byte {
	if x != nil {
		return x.WalletAddrSig
	}
	return nil
}

func (x *Auction_Bid) GetBidderId() string {
	if x != nil {
		return x.BidderId
	}
	return ""
}

func (x *Auction_Bid) GetAskPrice() int64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *Auction_Bid) GetVerifiedAskPrice() int64 {
	if x != nil {
		return x.VerifiedAskPrice
	}
	return 0
}

func (x *Auction_Bid) GetStartEpoch() uint64 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *Auction_Bid) GetFastRetrieval() bool {
	if x != nil {
		return x.FastRetrieval
	}
	return false
}

func (x *Auction_Bid) GetReceivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedAt
	}
	return nil
}

var File_broker_auctioneer_v1_auctioneer_proto protoreflect.FileDescriptor

var file_broker_auctioneer_v1_auctioneer_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x65, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff,
	0x07, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x65, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x64, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x42, 0x69, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x69, 0x64, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x69,
	0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x5a, 0x0a, 0x09, 0x42,
	0x69, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xb9, 0x02, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x26,
	0x0a, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x73, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x53, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x6b,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x25, 0x0a, 0x0e, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x61, 0x73, 0x74, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x6b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04,
	0x22, 0xd1, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x65, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x64, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0xde, 0x01, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x54, 0x6f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x61, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65, 0x69, 0x6f, 0x2f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_broker_auctioneer_v1_auctioneer_proto_rawDescOnce sync.Once
	file_broker_auctioneer_v1_auctioneer_proto_rawDescData = file_broker_auctioneer_v1_auctioneer_proto_rawDesc
)

func file_broker_auctioneer_v1_auctioneer_proto_rawDescGZIP() []byte {
	file_broker_auctioneer_v1_auctioneer_proto_rawDescOnce.Do(func() {
		file_broker_auctioneer_v1_auctioneer_proto_rawDescData = protoimpl.X.CompressGZIP(file_broker_auctioneer_v1_auctioneer_proto_rawDescData)
	})
	return file_broker_auctioneer_v1_auctioneer_proto_rawDescData
}

var file_broker_auctioneer_v1_auctioneer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_broker_auctioneer_v1_auctioneer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_broker_auctioneer_v1_auctioneer_proto_goTypes = []interface{}{
	(Auction_Status)(0),            // 0: broker.auctioneer.v1.Auction.Status
	(*Auction)(nil),                // 1: broker.auctioneer.v1.Auction
	(*ReadyToAuctionRequest)(nil),  // 2: broker.auctioneer.v1.ReadyToAuctionRequest
	(*ReadyToAuctionResponse)(nil), // 3: broker.auctioneer.v1.ReadyToAuctionResponse
	(*GetAuctionRequest)(nil),      // 4: broker.auctioneer.v1.GetAuctionRequest
	(*GetAuctionResponse)(nil),     // 5: broker.auctioneer.v1.GetAuctionResponse
	nil,                            // 6: broker.auctioneer.v1.Auction.BidsEntry
	(*Auction_Bid)(nil),            // 7: broker.auctioneer.v1.Auction.Bid
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
}
var file_broker_auctioneer_v1_auctioneer_proto_depIdxs = []int32{
	0, // 0: broker.auctioneer.v1.Auction.status:type_name -> broker.auctioneer.v1.Auction.Status
	6, // 1: broker.auctioneer.v1.Auction.bids:type_name -> broker.auctioneer.v1.Auction.BidsEntry
	8, // 2: broker.auctioneer.v1.Auction.started_at:type_name -> google.protobuf.Timestamp
	1, // 3: broker.auctioneer.v1.GetAuctionResponse.auction:type_name -> broker.auctioneer.v1.Auction
	7, // 4: broker.auctioneer.v1.Auction.BidsEntry.value:type_name -> broker.auctioneer.v1.Auction.Bid
	8, // 5: broker.auctioneer.v1.Auction.Bid.received_at:type_name -> google.protobuf.Timestamp
	2, // 6: broker.auctioneer.v1.APIService.ReadyToAuction:input_type -> broker.auctioneer.v1.ReadyToAuctionRequest
	4, // 7: broker.auctioneer.v1.APIService.GetAuction:input_type -> broker.auctioneer.v1.GetAuctionRequest
	3, // 8: broker.auctioneer.v1.APIService.ReadyToAuction:output_type -> broker.auctioneer.v1.ReadyToAuctionResponse
	5, // 9: broker.auctioneer.v1.APIService.GetAuction:output_type -> broker.auctioneer.v1.GetAuctionResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_broker_auctioneer_v1_auctioneer_proto_init() }
func file_broker_auctioneer_v1_auctioneer_proto_init() {
	if File_broker_auctioneer_v1_auctioneer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broker_auctioneer_v1_auctioneer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_broker_auctioneer_v1_auctioneer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyToAuctionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_broker_auctioneer_v1_auctioneer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyToAuctionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_broker_auctioneer_v1_auctioneer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuctionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_broker_auctioneer_v1_auctioneer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuctionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_broker_auctioneer_v1_auctioneer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auction_Bid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_broker_auctioneer_v1_auctioneer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_broker_auctioneer_v1_auctioneer_proto_goTypes,
		DependencyIndexes: file_broker_auctioneer_v1_auctioneer_proto_depIdxs,
		EnumInfos:         file_broker_auctioneer_v1_auctioneer_proto_enumTypes,
		MessageInfos:      file_broker_auctioneer_v1_auctioneer_proto_msgTypes,
	}.Build()
	File_broker_auctioneer_v1_auctioneer_proto = out.File
	file_broker_auctioneer_v1_auctioneer_proto_rawDesc = nil
	file_broker_auctioneer_v1_auctioneer_proto_goTypes = nil
	file_broker_auctioneer_v1_auctioneer_proto_depIdxs = nil
}
