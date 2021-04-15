// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cmd/auctioneerd/pb/auctioneer.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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
	return file_cmd_auctioneerd_pb_auctioneer_proto_enumTypes[0].Descriptor()
}

func (Auction_Status) Type() protoreflect.EnumType {
	return &file_cmd_auctioneerd_pb_auctioneer_proto_enumTypes[0]
}

func (x Auction_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Auction_Status.Descriptor instead.
func (Auction_Status) EnumDescriptor() ([]byte, []int) {
	return file_cmd_auctioneerd_pb_auctioneer_proto_rawDescGZIP(), []int{0, 0}
}

type Auction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    Auction_Status          `protobuf:"varint,2,opt,name=status,proto3,enum=cmd.auctioneerd.pb.Auction_Status" json:"status,omitempty"`
	Bids      map[string]*Auction_Bid `protobuf:"bytes,3,rep,name=bids,proto3" json:"bids,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Winner    string                  `protobuf:"bytes,4,opt,name=winner,proto3" json:"winner,omitempty"`
	StartedAt *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	Duration  int64                   `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Error     string                  `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Auction) Reset() {
	*x = Auction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction) ProtoMessage() {}

func (x *Auction) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[0]
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
	return file_cmd_auctioneerd_pb_auctioneer_proto_rawDescGZIP(), []int{0}
}

func (x *Auction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
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

func (x *Auction) GetWinner() string {
	if x != nil {
		return x.Winner
	}
	return ""
}

func (x *Auction) GetStartedAt() *timestamp.Timestamp {
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

func (x *Auction) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CreateAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAuctionRequest) Reset() {
	*x = CreateAuctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuctionRequest) ProtoMessage() {}

func (x *CreateAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuctionRequest.ProtoReflect.Descriptor instead.
func (*CreateAuctionRequest) Descriptor() ([]byte, []int) {
	return file_cmd_auctioneerd_pb_auctioneer_proto_rawDescGZIP(), []int{1}
}

type CreateAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateAuctionResponse) Reset() {
	*x = CreateAuctionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuctionResponse) ProtoMessage() {}

func (x *CreateAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuctionResponse.ProtoReflect.Descriptor instead.
func (*CreateAuctionResponse) Descriptor() ([]byte, []int) {
	return file_cmd_auctioneerd_pb_auctioneer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuctionResponse) GetId() string {
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
		mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuctionRequest) ProtoMessage() {}

func (x *GetAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[3]
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
	return file_cmd_auctioneerd_pb_auctioneer_proto_rawDescGZIP(), []int{3}
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
		mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuctionResponse) ProtoMessage() {}

func (x *GetAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[4]
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
	return file_cmd_auctioneerd_pb_auctioneer_proto_rawDescGZIP(), []int{4}
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

	From       string               `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Amount     int64                `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ReceivedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
}

func (x *Auction_Bid) Reset() {
	*x = Auction_Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auction_Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction_Bid) ProtoMessage() {}

func (x *Auction_Bid) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[6]
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
	return file_cmd_auctioneerd_pb_auctioneer_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Auction_Bid) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Auction_Bid) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Auction_Bid) GetReceivedAt() *timestamp.Timestamp {
	if x != nil {
		return x.ReceivedAt
	}
	return nil
}

var File_cmd_auctioneerd_pb_auctioneer_proto protoreflect.FileDescriptor

var file_cmd_auctioneerd_pb_auctioneer_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6d, 0x64, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72,
	0x64, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6d, 0x64, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x65, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x04, 0x0a, 0x07, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65,
	0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x1a, 0x58, 0x0a, 0x09, 0x42, 0x69, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65,
	0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69,
	0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6e, 0x0a, 0x03,
	0x42, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6b, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xd3, 0x01, 0x0a,
	0x0a, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x63,
	0x6d, 0x64, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x64, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x65, 0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65, 0x69, 0x6f, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x65, 0x65, 0x72, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cmd_auctioneerd_pb_auctioneer_proto_rawDescOnce sync.Once
	file_cmd_auctioneerd_pb_auctioneer_proto_rawDescData = file_cmd_auctioneerd_pb_auctioneer_proto_rawDesc
)

func file_cmd_auctioneerd_pb_auctioneer_proto_rawDescGZIP() []byte {
	file_cmd_auctioneerd_pb_auctioneer_proto_rawDescOnce.Do(func() {
		file_cmd_auctioneerd_pb_auctioneer_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_auctioneerd_pb_auctioneer_proto_rawDescData)
	})
	return file_cmd_auctioneerd_pb_auctioneer_proto_rawDescData
}

var file_cmd_auctioneerd_pb_auctioneer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cmd_auctioneerd_pb_auctioneer_proto_goTypes = []interface{}{
	(Auction_Status)(0),           // 0: cmd.auctioneerd.pb.Auction.Status
	(*Auction)(nil),               // 1: cmd.auctioneerd.pb.Auction
	(*CreateAuctionRequest)(nil),  // 2: cmd.auctioneerd.pb.CreateAuctionRequest
	(*CreateAuctionResponse)(nil), // 3: cmd.auctioneerd.pb.CreateAuctionResponse
	(*GetAuctionRequest)(nil),     // 4: cmd.auctioneerd.pb.GetAuctionRequest
	(*GetAuctionResponse)(nil),    // 5: cmd.auctioneerd.pb.GetAuctionResponse
	nil,                           // 6: cmd.auctioneerd.pb.Auction.BidsEntry
	(*Auction_Bid)(nil),           // 7: cmd.auctioneerd.pb.Auction.Bid
	(*timestamp.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_cmd_auctioneerd_pb_auctioneer_proto_depIdxs = []int32{
	0, // 0: cmd.auctioneerd.pb.Auction.status:type_name -> cmd.auctioneerd.pb.Auction.Status
	6, // 1: cmd.auctioneerd.pb.Auction.bids:type_name -> cmd.auctioneerd.pb.Auction.BidsEntry
	8, // 2: cmd.auctioneerd.pb.Auction.started_at:type_name -> google.protobuf.Timestamp
	1, // 3: cmd.auctioneerd.pb.GetAuctionResponse.auction:type_name -> cmd.auctioneerd.pb.Auction
	7, // 4: cmd.auctioneerd.pb.Auction.BidsEntry.value:type_name -> cmd.auctioneerd.pb.Auction.Bid
	8, // 5: cmd.auctioneerd.pb.Auction.Bid.received_at:type_name -> google.protobuf.Timestamp
	2, // 6: cmd.auctioneerd.pb.APIService.CreateAuction:input_type -> cmd.auctioneerd.pb.CreateAuctionRequest
	4, // 7: cmd.auctioneerd.pb.APIService.GetAuction:input_type -> cmd.auctioneerd.pb.GetAuctionRequest
	3, // 8: cmd.auctioneerd.pb.APIService.CreateAuction:output_type -> cmd.auctioneerd.pb.CreateAuctionResponse
	5, // 9: cmd.auctioneerd.pb.APIService.GetAuction:output_type -> cmd.auctioneerd.pb.GetAuctionResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cmd_auctioneerd_pb_auctioneer_proto_init() }
func file_cmd_auctioneerd_pb_auctioneer_proto_init() {
	if File_cmd_auctioneerd_pb_auctioneer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuctionRequest); i {
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
		file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuctionResponse); i {
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
		file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_cmd_auctioneerd_pb_auctioneer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_auctioneerd_pb_auctioneer_proto_goTypes,
		DependencyIndexes: file_cmd_auctioneerd_pb_auctioneer_proto_depIdxs,
		EnumInfos:         file_cmd_auctioneerd_pb_auctioneer_proto_enumTypes,
		MessageInfos:      file_cmd_auctioneerd_pb_auctioneer_proto_msgTypes,
	}.Build()
	File_cmd_auctioneerd_pb_auctioneer_proto = out.File
	file_cmd_auctioneerd_pb_auctioneer_proto_rawDesc = nil
	file_cmd_auctioneerd_pb_auctioneer_proto_goTypes = nil
	file_cmd_auctioneerd_pb_auctioneer_proto_depIdxs = nil
}
