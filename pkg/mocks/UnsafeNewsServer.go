// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeNewsServer is an autogenerated mock type for the UnsafeNewsServer type
type UnsafeNewsServer struct {
	mock.Mock
}

type UnsafeNewsServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeNewsServer) EXPECT() *UnsafeNewsServer_Expecter {
	return &UnsafeNewsServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedNewsServer provides a mock function with given fields:
func (_m *UnsafeNewsServer) mustEmbedUnimplementedNewsServer() {
	_m.Called()
}

// UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedNewsServer'
type UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedNewsServer is a helper method to define mock.On call
func (_e *UnsafeNewsServer_Expecter) mustEmbedUnimplementedNewsServer() *UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call {
	return &UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call{Call: _e.mock.On("mustEmbedUnimplementedNewsServer")}
}

func (_c *UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call) Run(run func()) *UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call) Return() *UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call) RunAndReturn(run func()) *UnsafeNewsServer_mustEmbedUnimplementedNewsServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewUnsafeNewsServer creates a new instance of UnsafeNewsServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeNewsServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeNewsServer {
	mock := &UnsafeNewsServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
