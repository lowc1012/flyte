// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"

	mock "github.com/stretchr/testify/mock"
)

// QualityOfServiceConfiguration is an autogenerated mock type for the QualityOfServiceConfiguration type
type QualityOfServiceConfiguration struct {
	mock.Mock
}

type QualityOfServiceConfiguration_GetDefaultTiers struct {
	*mock.Call
}

func (_m QualityOfServiceConfiguration_GetDefaultTiers) Return(_a0 map[string]core.QualityOfService_Tier) *QualityOfServiceConfiguration_GetDefaultTiers {
	return &QualityOfServiceConfiguration_GetDefaultTiers{Call: _m.Call.Return(_a0)}
}

func (_m *QualityOfServiceConfiguration) OnGetDefaultTiers() *QualityOfServiceConfiguration_GetDefaultTiers {
	c_call := _m.On("GetDefaultTiers")
	return &QualityOfServiceConfiguration_GetDefaultTiers{Call: c_call}
}

func (_m *QualityOfServiceConfiguration) OnGetDefaultTiersMatch(matchers ...interface{}) *QualityOfServiceConfiguration_GetDefaultTiers {
	c_call := _m.On("GetDefaultTiers", matchers...)
	return &QualityOfServiceConfiguration_GetDefaultTiers{Call: c_call}
}

// GetDefaultTiers provides a mock function with given fields:
func (_m *QualityOfServiceConfiguration) GetDefaultTiers() map[string]core.QualityOfService_Tier {
	ret := _m.Called()

	var r0 map[string]core.QualityOfService_Tier
	if rf, ok := ret.Get(0).(func() map[string]core.QualityOfService_Tier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]core.QualityOfService_Tier)
		}
	}

	return r0
}

type QualityOfServiceConfiguration_GetTierExecutionValues struct {
	*mock.Call
}

func (_m QualityOfServiceConfiguration_GetTierExecutionValues) Return(_a0 map[core.QualityOfService_Tier]*core.QualityOfServiceSpec) *QualityOfServiceConfiguration_GetTierExecutionValues {
	return &QualityOfServiceConfiguration_GetTierExecutionValues{Call: _m.Call.Return(_a0)}
}

func (_m *QualityOfServiceConfiguration) OnGetTierExecutionValues() *QualityOfServiceConfiguration_GetTierExecutionValues {
	c_call := _m.On("GetTierExecutionValues")
	return &QualityOfServiceConfiguration_GetTierExecutionValues{Call: c_call}
}

func (_m *QualityOfServiceConfiguration) OnGetTierExecutionValuesMatch(matchers ...interface{}) *QualityOfServiceConfiguration_GetTierExecutionValues {
	c_call := _m.On("GetTierExecutionValues", matchers...)
	return &QualityOfServiceConfiguration_GetTierExecutionValues{Call: c_call}
}

// GetTierExecutionValues provides a mock function with given fields:
func (_m *QualityOfServiceConfiguration) GetTierExecutionValues() map[core.QualityOfService_Tier]*core.QualityOfServiceSpec {
	ret := _m.Called()

	var r0 map[core.QualityOfService_Tier]*core.QualityOfServiceSpec
	if rf, ok := ret.Get(0).(func() map[core.QualityOfService_Tier]*core.QualityOfServiceSpec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[core.QualityOfService_Tier]*core.QualityOfServiceSpec)
		}
	}

	return r0
}
