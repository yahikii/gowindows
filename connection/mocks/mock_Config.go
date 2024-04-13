// Code generated by mockery v2.38.0. DO NOT EDIT.

package connection

import (
	connection "github.com/d-strobel/gowindows/connection"
	mock "github.com/stretchr/testify/mock"
)

// MockConfig is an autogenerated mock type for the Config type
type MockConfig struct {
	mock.Mock
}

type MockConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConfig) EXPECT() *MockConfig_Expecter {
	return &MockConfig_Expecter{mock: &_m.Mock}
}

// NewConnection provides a mock function with given fields:
func (_m *MockConfig) NewConnection() (connection.Connection, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NewConnection")
	}

	var r0 connection.Connection
	var r1 error
	if rf, ok := ret.Get(0).(func() (connection.Connection, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() connection.Connection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(connection.Connection)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConfig_NewConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewConnection'
type MockConfig_NewConnection_Call struct {
	*mock.Call
}

// NewConnection is a helper method to define mock.On call
func (_e *MockConfig_Expecter) NewConnection() *MockConfig_NewConnection_Call {
	return &MockConfig_NewConnection_Call{Call: _e.mock.On("NewConnection")}
}

func (_c *MockConfig_NewConnection_Call) Run(run func()) *MockConfig_NewConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_NewConnection_Call) Return(_a0 connection.Connection, _a1 error) *MockConfig_NewConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConfig_NewConnection_Call) RunAndReturn(run func() (connection.Connection, error)) *MockConfig_NewConnection_Call {
	_c.Call.Return(run)
	return _c
}

// defaults provides a mock function with given fields:
func (_m *MockConfig) defaults() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for defaults")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConfig_defaults_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'defaults'
type MockConfig_defaults_Call struct {
	*mock.Call
}

// defaults is a helper method to define mock.On call
func (_e *MockConfig_Expecter) defaults() *MockConfig_defaults_Call {
	return &MockConfig_defaults_Call{Call: _e.mock.On("defaults")}
}

func (_c *MockConfig_defaults_Call) Run(run func()) *MockConfig_defaults_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_defaults_Call) Return(_a0 error) *MockConfig_defaults_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_defaults_Call) RunAndReturn(run func() error) *MockConfig_defaults_Call {
	_c.Call.Return(run)
	return _c
}

// validate provides a mock function with given fields:
func (_m *MockConfig) validate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConfig_validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'validate'
type MockConfig_validate_Call struct {
	*mock.Call
}

// validate is a helper method to define mock.On call
func (_e *MockConfig_Expecter) validate() *MockConfig_validate_Call {
	return &MockConfig_validate_Call{Call: _e.mock.On("validate")}
}

func (_c *MockConfig_validate_Call) Run(run func()) *MockConfig_validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_validate_Call) Return(_a0 error) *MockConfig_validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_validate_Call) RunAndReturn(run func() error) *MockConfig_validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConfig creates a new instance of MockConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConfig {
	mock := &MockConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
