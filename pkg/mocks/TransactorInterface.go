// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TransactorInterface is an autogenerated mock type for the TransactorInterface type
type TransactorInterface struct {
	mock.Mock
}

type TransactorInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *TransactorInterface) EXPECT() *TransactorInterface_Expecter {
	return &TransactorInterface_Expecter{mock: &_m.Mock}
}

// MakeTransaction provides a mock function with given fields: _a0, _a1
func (_m *TransactorInterface) MakeTransaction(_a0 context.Context, _a1 func(context.Context) error) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for MakeTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransactorInterface_MakeTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MakeTransaction'
type TransactorInterface_MakeTransaction_Call struct {
	*mock.Call
}

// MakeTransaction is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 func(context.Context) error
func (_e *TransactorInterface_Expecter) MakeTransaction(_a0 interface{}, _a1 interface{}) *TransactorInterface_MakeTransaction_Call {
	return &TransactorInterface_MakeTransaction_Call{Call: _e.mock.On("MakeTransaction", _a0, _a1)}
}

func (_c *TransactorInterface_MakeTransaction_Call) Run(run func(_a0 context.Context, _a1 func(context.Context) error)) *TransactorInterface_MakeTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context) error))
	})
	return _c
}

func (_c *TransactorInterface_MakeTransaction_Call) Return(_a0 error) *TransactorInterface_MakeTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TransactorInterface_MakeTransaction_Call) RunAndReturn(run func(context.Context, func(context.Context) error) error) *TransactorInterface_MakeTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewTransactorInterface creates a new instance of TransactorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactorInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactorInterface {
	mock := &TransactorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
