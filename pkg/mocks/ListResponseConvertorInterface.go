// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	"grpc/internal/domain/news"
	"grpc/pkg/proto_gen/grpc"

	mock "github.com/stretchr/testify/mock"
)

// ListResponseConvertorInterface is an autogenerated mock type for the ListResponseConvertorInterface type
type ListResponseConvertorInterface struct {
	mock.Mock
}

type ListResponseConvertorInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ListResponseConvertorInterface) EXPECT() *ListResponseConvertorInterface_Expecter {
	return &ListResponseConvertorInterface_Expecter{mock: &_m.Mock}
}

// Convert provides a mock function with given fields: _a0
func (_m *ListResponseConvertorInterface) Convert(_a0 []news.New) *grpc.NewsList {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Convert")
	}

	var r0 *grpc.NewsList
	if rf, ok := ret.Get(0).(func([]news.New) *grpc.NewsList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.NewsList)
		}
	}

	return r0
}

// ListResponseConvertorInterface_Convert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Convert'
type ListResponseConvertorInterface_Convert_Call struct {
	*mock.Call
}

// Convert is a helper method to define mock.On call
//   - _a0 []news.New
func (_e *ListResponseConvertorInterface_Expecter) Convert(_a0 interface{}) *ListResponseConvertorInterface_Convert_Call {
	return &ListResponseConvertorInterface_Convert_Call{Call: _e.mock.On("Convert", _a0)}
}

func (_c *ListResponseConvertorInterface_Convert_Call) Run(run func(_a0 []news.New)) *ListResponseConvertorInterface_Convert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]news.New))
	})
	return _c
}

func (_c *ListResponseConvertorInterface_Convert_Call) Return(_a0 *grpc.NewsList) *ListResponseConvertorInterface_Convert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ListResponseConvertorInterface_Convert_Call) RunAndReturn(run func([]news.New) *grpc.NewsList) *ListResponseConvertorInterface_Convert_Call {
	_c.Call.Return(run)
	return _c
}

// NewListResponseConvertorInterface creates a new instance of ListResponseConvertorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewListResponseConvertorInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ListResponseConvertorInterface {
	mock := &ListResponseConvertorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
