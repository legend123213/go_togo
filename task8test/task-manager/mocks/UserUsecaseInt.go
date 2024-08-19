// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecaseInt is an autogenerated mock type for the UserUsecaseInt type
type UserUsecaseInt struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UserUsecaseInt) Delete(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: id, user
func (_m *UserUsecaseInt) Edit(id string, user *domain.User) (*domain.User, error) {
	ret := _m.Called(id, user)

	if len(ret) == 0 {
		panic("no return value specified for Edit")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *domain.User) (*domain.User, error)); ok {
		return rf(id, user)
	}
	if rf, ok := ret.Get(0).(func(string, *domain.User) *domain.User); ok {
		r0 = rf(id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *domain.User) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: id
func (_m *UserUsecaseInt) Fetch(id string) (*domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchAllUser provides a mock function with given fields:
func (_m *UserUsecaseInt) FetchAllUser() *[]domain.User {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FetchAllUser")
	}

	var r0 *[]domain.User
	if rf, ok := ret.Get(0).(func() *[]domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.User)
		}
	}

	return r0
}

// FetchUserByUname provides a mock function with given fields: username
func (_m *UserUsecaseInt) FetchUserByUname(username string) (*domain.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for FetchUserByUname")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsUsernameUnique provides a mock function with given fields: username
func (_m *UserUsecaseInt) IsUsernameUnique(username string) error {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for IsUsernameUnique")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: u
func (_m *UserUsecaseInt) Login(u *domain.User) (string, error) {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User) (string, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(*domain.User) string); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: u
func (_m *UserUsecaseInt) Register(u *domain.User) (string, error) {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User) (string, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(*domain.User) string); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleChanger provides a mock function with given fields: id
func (_m *UserUsecaseInt) RoleChanger(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for RoleChanger")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUsecaseInt creates a new instance of UserUsecaseInt. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecaseInt(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecaseInt {
	mock := &UserUsecaseInt{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}