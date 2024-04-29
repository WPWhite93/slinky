// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	consumertypes "github.com/cosmos/interchain-security/v5/x/ccv/consumer/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CCValidatorStore is an autogenerated mock type for the CCValidatorStore type
type CCValidatorStore struct {
	mock.Mock
}

// GetAllCCValidator provides a mock function with given fields: ctx
func (_m *CCValidatorStore) GetAllCCValidator(ctx types.Context) []consumertypes.CrossChainValidator {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCCValidator")
	}

	var r0 []consumertypes.CrossChainValidator
	if rf, ok := ret.Get(0).(func(types.Context) []consumertypes.CrossChainValidator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consumertypes.CrossChainValidator)
		}
	}

	return r0
}

// GetCCValidator provides a mock function with given fields: ctx, addr
func (_m *CCValidatorStore) GetCCValidator(ctx types.Context, addr []byte) (consumertypes.CrossChainValidator, bool) {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for GetCCValidator")
	}

	var r0 consumertypes.CrossChainValidator
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, []byte) (consumertypes.CrossChainValidator, bool)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []byte) consumertypes.CrossChainValidator); ok {
		r0 = rf(ctx, addr)
	} else {
		r0 = ret.Get(0).(consumertypes.CrossChainValidator)
	}

	if rf, ok := ret.Get(1).(func(types.Context, []byte) bool); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// NewCCValidatorStore creates a new instance of CCValidatorStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCCValidatorStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *CCValidatorStore {
	mock := &CCValidatorStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}