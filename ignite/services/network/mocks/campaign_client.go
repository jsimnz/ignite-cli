// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tendermint/spn/x/campaign/types"
	"google.golang.org/grpc"
)

// CampaignClient is an autogenerated mock type for the QueryClient type
type CampaignClient struct {
	mock.Mock
}

// AuctionsOfCampaign provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) AuctionsOfCampaign(ctx context.Context, in *types.QueryAuctionsOfCampaignRequest, opts ...grpc.CallOption) (*types.QueryAuctionsOfCampaignResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAuctionsOfCampaignResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAuctionsOfCampaignRequest, ...grpc.CallOption) *types.QueryAuctionsOfCampaignResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAuctionsOfCampaignResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAuctionsOfCampaignRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Campaign provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) Campaign(ctx context.Context, in *types.QueryGetCampaignRequest, opts ...grpc.CallOption) (*types.QueryGetCampaignResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetCampaignResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetCampaignRequest, ...grpc.CallOption) *types.QueryGetCampaignResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetCampaignResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetCampaignRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CampaignAll provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) CampaignAll(ctx context.Context, in *types.QueryAllCampaignRequest, opts ...grpc.CallOption) (*types.QueryAllCampaignResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllCampaignResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllCampaignRequest, ...grpc.CallOption) *types.QueryAllCampaignResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllCampaignResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllCampaignRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CampaignChains provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) CampaignChains(ctx context.Context, in *types.QueryGetCampaignChainsRequest, opts ...grpc.CallOption) (*types.QueryGetCampaignChainsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetCampaignChainsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetCampaignChainsRequest, ...grpc.CallOption) *types.QueryGetCampaignChainsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetCampaignChainsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetCampaignChainsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CampaignSummaries provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) CampaignSummaries(ctx context.Context, in *types.QueryCampaignSummariesRequest, opts ...grpc.CallOption) (*types.QueryCampaignSummariesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryCampaignSummariesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryCampaignSummariesRequest, ...grpc.CallOption) *types.QueryCampaignSummariesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryCampaignSummariesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryCampaignSummariesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CampaignSummary provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) CampaignSummary(ctx context.Context, in *types.QueryCampaignSummaryRequest, opts ...grpc.CallOption) (*types.QueryCampaignSummaryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryCampaignSummaryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryCampaignSummaryRequest, ...grpc.CallOption) *types.QueryCampaignSummaryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryCampaignSummaryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryCampaignSummaryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetAccount provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetAccount(ctx context.Context, in *types.QueryGetMainnetAccountRequest, opts ...grpc.CallOption) (*types.QueryGetMainnetAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetMainnetAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetMainnetAccountRequest, ...grpc.CallOption) *types.QueryGetMainnetAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetMainnetAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetMainnetAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetAccountAll provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetAccountAll(ctx context.Context, in *types.QueryAllMainnetAccountRequest, opts ...grpc.CallOption) (*types.QueryAllMainnetAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllMainnetAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMainnetAccountRequest, ...grpc.CallOption) *types.QueryAllMainnetAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMainnetAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMainnetAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetAccountBalance provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetAccountBalance(ctx context.Context, in *types.QueryGetMainnetAccountBalanceRequest, opts ...grpc.CallOption) (*types.QueryGetMainnetAccountBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetMainnetAccountBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetMainnetAccountBalanceRequest, ...grpc.CallOption) *types.QueryGetMainnetAccountBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetMainnetAccountBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetMainnetAccountBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetAccountBalanceAll provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetAccountBalanceAll(ctx context.Context, in *types.QueryAllMainnetAccountBalanceRequest, opts ...grpc.CallOption) (*types.QueryAllMainnetAccountBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllMainnetAccountBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMainnetAccountBalanceRequest, ...grpc.CallOption) *types.QueryAllMainnetAccountBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMainnetAccountBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMainnetAccountBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetVestingAccount provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetVestingAccount(ctx context.Context, in *types.QueryGetMainnetVestingAccountRequest, opts ...grpc.CallOption) (*types.QueryGetMainnetVestingAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetMainnetVestingAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetMainnetVestingAccountRequest, ...grpc.CallOption) *types.QueryGetMainnetVestingAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetMainnetVestingAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetMainnetVestingAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetVestingAccountAll provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetVestingAccountAll(ctx context.Context, in *types.QueryAllMainnetVestingAccountRequest, opts ...grpc.CallOption) (*types.QueryAllMainnetVestingAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllMainnetVestingAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMainnetVestingAccountRequest, ...grpc.CallOption) *types.QueryAllMainnetVestingAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMainnetVestingAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMainnetVestingAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) *types.QueryParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TotalShares provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) TotalShares(ctx context.Context, in *types.QueryTotalSharesRequest, opts ...grpc.CallOption) (*types.QueryTotalSharesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryTotalSharesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryTotalSharesRequest, ...grpc.CallOption) *types.QueryTotalSharesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryTotalSharesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryTotalSharesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCampaignClient creates a new instance of CampaignClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCampaignClient(t testing.TB) *CampaignClient {
	mock := &CampaignClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
