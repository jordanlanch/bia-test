// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/jordanlanch/bia-test/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) *domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, user
func (_m *UserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) *domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
