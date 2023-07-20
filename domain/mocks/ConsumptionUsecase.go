// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/jordanlanch/bia-test/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ConsumptionUsecase is an autogenerated mock type for the ConsumptionUsecase type
type ConsumptionUsecase struct {
	mock.Mock
}

// GetConsumptionById provides a mock function with given fields: id
func (_m *ConsumptionUsecase) GetConsumptionById(id uuid.UUID) (*domain.Consumption, error) {
	ret := _m.Called(id)

	var r0 *domain.Consumption
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.Consumption); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Consumption)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsumptionsByPeriod provides a mock function with given fields: period_type, start, end, meterIDs, pagination
func (_m *ConsumptionUsecase) GetConsumptionsByPeriod(period_type string, start string, end string, meterIDs []int, pagination *domain.Pagination) (*domain.Response, error) {
	ret := _m.Called(period_type, start, end, meterIDs, pagination)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(string, string, string, []int, *domain.Pagination) *domain.Response); ok {
		r0 = rf(period_type, start, end, meterIDs, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, []int, *domain.Pagination) error); ok {
		r1 = rf(period_type, start, end, meterIDs, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewConsumptionUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumptionUsecase creates a new instance of ConsumptionUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumptionUsecase(t mockConstructorTestingTNewConsumptionUsecase) *ConsumptionUsecase {
	mock := &ConsumptionUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
