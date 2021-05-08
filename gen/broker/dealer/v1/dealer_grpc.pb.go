// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dealer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIServiceClient interface {
	ReadyToCreateDeals(ctx context.Context, in *ReadyToCreateDealsRequest, opts ...grpc.CallOption) (*ReadyToCreateDealsResponse, error)
}

type aPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIServiceClient(cc grpc.ClientConnInterface) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) ReadyToCreateDeals(ctx context.Context, in *ReadyToCreateDealsRequest, opts ...grpc.CallOption) (*ReadyToCreateDealsResponse, error) {
	out := new(ReadyToCreateDealsResponse)
	err := c.cc.Invoke(ctx, "/broker.dealer.v1.APIService/ReadyToCreateDeals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
// All implementations must embed UnimplementedAPIServiceServer
// for forward compatibility
type APIServiceServer interface {
	ReadyToCreateDeals(context.Context, *ReadyToCreateDealsRequest) (*ReadyToCreateDealsResponse, error)
	mustEmbedUnimplementedAPIServiceServer()
}

// UnimplementedAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (UnimplementedAPIServiceServer) ReadyToCreateDeals(context.Context, *ReadyToCreateDealsRequest) (*ReadyToCreateDealsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadyToCreateDeals not implemented")
}
func (UnimplementedAPIServiceServer) mustEmbedUnimplementedAPIServiceServer() {}

// UnsafeAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServiceServer will
// result in compilation errors.
type UnsafeAPIServiceServer interface {
	mustEmbedUnimplementedAPIServiceServer()
}

func RegisterAPIServiceServer(s grpc.ServiceRegistrar, srv APIServiceServer) {
	s.RegisterService(&APIService_ServiceDesc, srv)
}

func _APIService_ReadyToCreateDeals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadyToCreateDealsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).ReadyToCreateDeals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broker.dealer.v1.APIService/ReadyToCreateDeals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).ReadyToCreateDeals(ctx, req.(*ReadyToCreateDealsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// APIService_ServiceDesc is the grpc.ServiceDesc for APIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "broker.dealer.v1.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadyToCreateDeals",
			Handler:    _APIService_ReadyToCreateDeals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "broker/dealer/v1/dealer.proto",
}