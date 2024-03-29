// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	"grpc/pkg/proto_gen/grpc"

	mock "github.com/stretchr/testify/mock"
)

// NewsServer is an autogenerated mock type for the NewsServer type
type NewsServer struct {
	mock.Mock
}

type NewsServer_Expecter struct {
	mock *mock.Mock
}

func (_m *NewsServer) EXPECT() *NewsServer_Expecter {
	return &NewsServer_Expecter{mock: &_m.Mock}
}

// List provides a mock function with given fields: _a0, _a1
func (_m *NewsServer) List(_a0 context.Context, _a1 *grpc.ListRequest) (*grpc.NewsList, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *grpc.NewsList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *grpc.ListRequest) (*grpc.NewsList, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *grpc.ListRequest) *grpc.NewsList); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.NewsList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *grpc.ListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewsServer_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type NewsServer_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *grpc.ListRequest
func (_e *NewsServer_Expecter) List(_a0 interface{}, _a1 interface{}) *NewsServer_List_Call {
	return &NewsServer_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *NewsServer_List_Call) Run(run func(_a0 context.Context, _a1 *grpc.ListRequest)) *NewsServer_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*grpc.ListRequest))
	})
	return _c
}

func (_c *NewsServer_List_Call) Return(_a0 *grpc.NewsList, _a1 error) *NewsServer_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NewsServer_List_Call) RunAndReturn(run func(context.Context, *grpc.ListRequest) (*grpc.NewsList, error)) *NewsServer_List_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedNewsServer provides a mock function with given fields:
func (_m *NewsServer) mustEmbedUnimplementedNewsServer() {
	_m.Called()
}

// NewsServer_mustEmbedUnimplementedNewsServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedNewsServer'
type NewsServer_mustEmbedUnimplementedNewsServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedNewsServer is a helper method to define mock.On call
func (_e *NewsServer_Expecter) mustEmbedUnimplementedNewsServer() *NewsServer_mustEmbedUnimplementedNewsServer_Call {
	return &NewsServer_mustEmbedUnimplementedNewsServer_Call{Call: _e.mock.On("mustEmbedUnimplementedNewsServer")}
}

func (_c *NewsServer_mustEmbedUnimplementedNewsServer_Call) Run(run func()) *NewsServer_mustEmbedUnimplementedNewsServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NewsServer_mustEmbedUnimplementedNewsServer_Call) Return() *NewsServer_mustEmbedUnimplementedNewsServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *NewsServer_mustEmbedUnimplementedNewsServer_Call) RunAndReturn(run func()) *NewsServer_mustEmbedUnimplementedNewsServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewNewsServer creates a new instance of NewsServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNewsServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *NewsServer {
	mock := &NewsServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
