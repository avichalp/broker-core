// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	chainapi "github.com/textileio/broker-core/gen/broker/chainapi/v1"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// ChainApiServiceClient is an autogenerated mock type for the ChainApiServiceClient type
type ChainApiServiceClient struct {
	mock.Mock
}

// HasDeposit provides a mock function with given fields: ctx, in, opts
func (_m *ChainApiServiceClient) HasDeposit(ctx context.Context, in *chainapi.HasDepositRequest, opts ...grpc.CallOption) (*chainapi.HasDepositResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chainapi.HasDepositResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chainapi.HasDepositRequest, ...grpc.CallOption) *chainapi.HasDepositResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chainapi.HasDepositResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chainapi.HasDepositRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePayload provides a mock function with given fields: ctx, in, opts
func (_m *ChainApiServiceClient) UpdatePayload(ctx context.Context, in *chainapi.UpdatePayloadRequest, opts ...grpc.CallOption) (*chainapi.UpdatePayloadResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chainapi.UpdatePayloadResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chainapi.UpdatePayloadRequest, ...grpc.CallOption) *chainapi.UpdatePayloadResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chainapi.UpdatePayloadResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chainapi.UpdatePayloadRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
