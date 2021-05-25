// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	cid "github.com/ipfs/go-cid"
	broker "github.com/textileio/broker-core/broker"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Broker is an autogenerated mock type for the Broker type
type Broker struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, c, meta
func (_m *Broker) Create(ctx context.Context, c cid.Cid, meta broker.Metadata) (broker.BrokerRequest, error) {
	ret := _m.Called(ctx, c, meta)

	var r0 broker.BrokerRequest
	if rf, ok := ret.Get(0).(func(context.Context, cid.Cid, broker.Metadata) broker.BrokerRequest); ok {
		r0 = rf(ctx, c, meta)
	} else {
		r0 = ret.Get(0).(broker.BrokerRequest)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, cid.Cid, broker.Metadata) error); ok {
		r1 = rf(ctx, c, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStorageDeal provides a mock function with given fields: ctx, batchCid, srids
func (_m *Broker) CreateStorageDeal(ctx context.Context, batchCid cid.Cid, srids []broker.BrokerRequestID) (broker.StorageDealID, error) {
	ret := _m.Called(ctx, batchCid, srids)

	var r0 broker.StorageDealID
	if rf, ok := ret.Get(0).(func(context.Context, cid.Cid, []broker.BrokerRequestID) broker.StorageDealID); ok {
		r0 = rf(ctx, batchCid, srids)
	} else {
		r0 = ret.Get(0).(broker.StorageDealID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, cid.Cid, []broker.BrokerRequestID) error); ok {
		r1 = rf(ctx, batchCid, srids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, ID
func (_m *Broker) Get(ctx context.Context, ID broker.BrokerRequestID) (broker.BrokerRequest, error) {
	ret := _m.Called(ctx, ID)

	var r0 broker.BrokerRequest
	if rf, ok := ret.Get(0).(func(context.Context, broker.BrokerRequestID) broker.BrokerRequest); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(broker.BrokerRequest)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, broker.BrokerRequestID) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageDealAuctioned provides a mock function with given fields: ctx, auction
func (_m *Broker) StorageDealAuctioned(ctx context.Context, auction broker.Auction) error {
	ret := _m.Called(ctx, auction)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, broker.Auction) error); ok {
		r0 = rf(ctx, auction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageDealFinalizedDeals provides a mock function with given fields: ctx, res
func (_m *Broker) StorageDealFinalizedDeals(ctx context.Context, res []broker.FinalizedAuctionDeal) error {
	ret := _m.Called(ctx, res)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []broker.FinalizedAuctionDeal) error); ok {
		r0 = rf(ctx, res)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageDealPrepared provides a mock function with given fields: ctx, id, pr
func (_m *Broker) StorageDealPrepared(ctx context.Context, id broker.StorageDealID, pr broker.DataPreparationResult) error {
	ret := _m.Called(ctx, id, pr)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, broker.StorageDealID, broker.DataPreparationResult) error); ok {
		r0 = rf(ctx, id, pr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageDealProposalAccepted provides a mock function with given fields: ctx, sdID, miner, proposalCid
func (_m *Broker) StorageDealProposalAccepted(ctx context.Context, sdID broker.StorageDealID, miner string, proposalCid cid.Cid) error {
	ret := _m.Called(ctx, sdID, miner, proposalCid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, broker.StorageDealID, string, cid.Cid) error); ok {
		r0 = rf(ctx, sdID, miner, proposalCid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}