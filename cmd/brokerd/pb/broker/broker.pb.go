// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cmd/brokerd/pb/broker/broker.proto

package broker

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

// CreateBR()
type CreateBRRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid  string      `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Meta *BRMetadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *CreateBRRequest) Reset() {
	*x = CreateBRRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBRRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBRRequest) ProtoMessage() {}

func (x *CreateBRRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBRRequest.ProtoReflect.Descriptor instead.
func (*CreateBRRequest) Descriptor() ([]byte, []int) {
	return file_cmd_brokerd_pb_broker_broker_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBRRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *CreateBRRequest) GetMeta() *BRMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type BRMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *BRMetadata) Reset() {
	*x = BRMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BRMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BRMetadata) ProtoMessage() {}

func (x *BRMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BRMetadata.ProtoReflect.Descriptor instead.
func (*BRMetadata) Descriptor() ([]byte, []int) {
	return file_cmd_brokerd_pb_broker_broker_proto_rawDescGZIP(), []int{1}
}

func (x *BRMetadata) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type CreateBRResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *BR `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *CreateBRResponse) Reset() {
	*x = CreateBRResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBRResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBRResponse) ProtoMessage() {}

func (x *CreateBRResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBRResponse.ProtoReflect.Descriptor instead.
func (*CreateBRResponse) Descriptor() ([]byte, []int) {
	return file_cmd_brokerd_pb_broker_broker_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBRResponse) GetRequest() *BR {
	if x != nil {
		return x.Request
	}
	return nil
}

// GetBR()
type GetBRRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBRRequest) Reset() {
	*x = GetBRRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBRRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBRRequest) ProtoMessage() {}

func (x *GetBRRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBRRequest.ProtoReflect.Descriptor instead.
func (*GetBRRequest) Descriptor() ([]byte, []int) {
	return file_cmd_brokerd_pb_broker_broker_proto_rawDescGZIP(), []int{3}
}

func (x *GetBRRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBRResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrokerRequest *BR `protobuf:"bytes,1,opt,name=broker_request,json=brokerRequest,proto3" json:"broker_request,omitempty"`
}

func (x *GetBRResponse) Reset() {
	*x = GetBRResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBRResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBRResponse) ProtoMessage() {}

func (x *GetBRResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBRResponse.ProtoReflect.Descriptor instead.
func (*GetBRResponse) Descriptor() ([]byte, []int) {
	return file_cmd_brokerd_pb_broker_broker_proto_rawDescGZIP(), []int{4}
}

func (x *GetBRResponse) GetBrokerRequest() *BR {
	if x != nil {
		return x.BrokerRequest
	}
	return nil
}

// General
type BR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cid           string               `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Meta          *BRMetadata          `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	StorageDealId string               `protobuf:"bytes,4,opt,name=storage_deal_id,json=storageDealId,proto3" json:"storage_deal_id,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *BR) Reset() {
	*x = BR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BR) ProtoMessage() {}

func (x *BR) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_brokerd_pb_broker_broker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BR.ProtoReflect.Descriptor instead.
func (*BR) Descriptor() ([]byte, []int) {
	return file_cmd_brokerd_pb_broker_broker_proto_rawDescGZIP(), []int{5}
}

func (x *BR) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BR) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *BR) GetMeta() *BRMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *BR) GetStorageDealId() string {
	if x != nil {
		return x.StorageDealId
	}
	return ""
}

func (x *BR) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BR) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_cmd_brokerd_pb_broker_broker_proto protoreflect.FileDescriptor

var file_cmd_brokerd_pb_broker_broker_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6d, 0x64, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x64, 0x2f, 0x70, 0x62,
	0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6d, 0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x64, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x52, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69,
	0x64, 0x12, 0x35, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x62,
	0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x52, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x0a, 0x42, 0x52, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x47,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x52, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x64, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x52, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x52,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x52,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70,
	0x62, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x52, 0x52, 0x0d, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfb, 0x01, 0x0a, 0x02, 0x42,
	0x52, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x64, 0x2e,
	0x70, 0x62, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x52, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x65, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xc1, 0x01, 0x0a, 0x0a, 0x41, 0x50, 0x49,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x52, 0x12, 0x26, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x64, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x52, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6d,
	0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x52, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x42, 0x52, 0x12,
	0x23, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x62,
	0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x52, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x64, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x52, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x69,
	0x6c, 0x65, 0x69, 0x6f, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x64, 0x2f, 0x70, 0x62, 0x2f,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_brokerd_pb_broker_broker_proto_rawDescOnce sync.Once
	file_cmd_brokerd_pb_broker_broker_proto_rawDescData = file_cmd_brokerd_pb_broker_broker_proto_rawDesc
)

func file_cmd_brokerd_pb_broker_broker_proto_rawDescGZIP() []byte {
	file_cmd_brokerd_pb_broker_broker_proto_rawDescOnce.Do(func() {
		file_cmd_brokerd_pb_broker_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_brokerd_pb_broker_broker_proto_rawDescData)
	})
	return file_cmd_brokerd_pb_broker_broker_proto_rawDescData
}

var file_cmd_brokerd_pb_broker_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cmd_brokerd_pb_broker_broker_proto_goTypes = []interface{}{
	(*CreateBRRequest)(nil),     // 0: cmd.brokerd.pb.broker.CreateBRRequest
	(*BRMetadata)(nil),          // 1: cmd.brokerd.pb.broker.BRMetadata
	(*CreateBRResponse)(nil),    // 2: cmd.brokerd.pb.broker.CreateBRResponse
	(*GetBRRequest)(nil),        // 3: cmd.brokerd.pb.broker.GetBRRequest
	(*GetBRResponse)(nil),       // 4: cmd.brokerd.pb.broker.GetBRResponse
	(*BR)(nil),                  // 5: cmd.brokerd.pb.broker.BR
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_cmd_brokerd_pb_broker_broker_proto_depIdxs = []int32{
	1, // 0: cmd.brokerd.pb.broker.CreateBRRequest.meta:type_name -> cmd.brokerd.pb.broker.BRMetadata
	5, // 1: cmd.brokerd.pb.broker.CreateBRResponse.request:type_name -> cmd.brokerd.pb.broker.BR
	5, // 2: cmd.brokerd.pb.broker.GetBRResponse.broker_request:type_name -> cmd.brokerd.pb.broker.BR
	1, // 3: cmd.brokerd.pb.broker.BR.meta:type_name -> cmd.brokerd.pb.broker.BRMetadata
	6, // 4: cmd.brokerd.pb.broker.BR.created_at:type_name -> google.protobuf.Timestamp
	6, // 5: cmd.brokerd.pb.broker.BR.updated_at:type_name -> google.protobuf.Timestamp
	0, // 6: cmd.brokerd.pb.broker.APIService.CreateBR:input_type -> cmd.brokerd.pb.broker.CreateBRRequest
	3, // 7: cmd.brokerd.pb.broker.APIService.GetBR:input_type -> cmd.brokerd.pb.broker.GetBRRequest
	2, // 8: cmd.brokerd.pb.broker.APIService.CreateBR:output_type -> cmd.brokerd.pb.broker.CreateBRResponse
	4, // 9: cmd.brokerd.pb.broker.APIService.GetBR:output_type -> cmd.brokerd.pb.broker.GetBRResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cmd_brokerd_pb_broker_broker_proto_init() }
func file_cmd_brokerd_pb_broker_broker_proto_init() {
	if File_cmd_brokerd_pb_broker_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_brokerd_pb_broker_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBRRequest); i {
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
		file_cmd_brokerd_pb_broker_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BRMetadata); i {
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
		file_cmd_brokerd_pb_broker_broker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBRResponse); i {
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
		file_cmd_brokerd_pb_broker_broker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBRRequest); i {
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
		file_cmd_brokerd_pb_broker_broker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBRResponse); i {
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
		file_cmd_brokerd_pb_broker_broker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BR); i {
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
			RawDescriptor: file_cmd_brokerd_pb_broker_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_brokerd_pb_broker_broker_proto_goTypes,
		DependencyIndexes: file_cmd_brokerd_pb_broker_broker_proto_depIdxs,
		MessageInfos:      file_cmd_brokerd_pb_broker_broker_proto_msgTypes,
	}.Build()
	File_cmd_brokerd_pb_broker_broker_proto = out.File
	file_cmd_brokerd_pb_broker_broker_proto_rawDesc = nil
	file_cmd_brokerd_pb_broker_broker_proto_goTypes = nil
	file_cmd_brokerd_pb_broker_broker_proto_depIdxs = nil
}
