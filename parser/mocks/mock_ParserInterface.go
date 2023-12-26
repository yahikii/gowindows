// Code generated by mockery v2.38.0. DO NOT EDIT.

package parser

import mock "github.com/stretchr/testify/mock"

// MockParserInterface is an autogenerated mock type for the ParserInterface type
type MockParserInterface struct {
	mock.Mock
}

type MockParserInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockParserInterface) EXPECT() *MockParserInterface_Expecter {
	return &MockParserInterface_Expecter{mock: &_m.Mock}
}

// DecodeCLIXML provides a mock function with given fields: clixml
func (_m *MockParserInterface) DecodeCLIXML(clixml string) (string, error) {
	ret := _m.Called(clixml)

	if len(ret) == 0 {
		panic("no return value specified for DecodeCLIXML")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(clixml)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(clixml)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clixml)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockParserInterface_DecodeCLIXML_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecodeCLIXML'
type MockParserInterface_DecodeCLIXML_Call struct {
	*mock.Call
}

// DecodeCLIXML is a helper method to define mock.On call
//   - clixml string
func (_e *MockParserInterface_Expecter) DecodeCLIXML(clixml interface{}) *MockParserInterface_DecodeCLIXML_Call {
	return &MockParserInterface_DecodeCLIXML_Call{Call: _e.mock.On("DecodeCLIXML", clixml)}
}

func (_c *MockParserInterface_DecodeCLIXML_Call) Run(run func(clixml string)) *MockParserInterface_DecodeCLIXML_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockParserInterface_DecodeCLIXML_Call) Return(_a0 string, _a1 error) *MockParserInterface_DecodeCLIXML_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockParserInterface_DecodeCLIXML_Call) RunAndReturn(run func(string) (string, error)) *MockParserInterface_DecodeCLIXML_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockParserInterface creates a new instance of MockParserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockParserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockParserInterface {
	mock := &MockParserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}