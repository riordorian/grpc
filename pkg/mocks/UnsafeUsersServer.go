// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeUsersServer is an autogenerated mock type for the UnsafeUsersServer type
type UnsafeUsersServer struct {
	mock.Mock
}

type UnsafeUsersServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeUsersServer) EXPECT() *UnsafeUsersServer_Expecter {
	return &UnsafeUsersServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedUsersServer provides a mock function with given fields:
func (_m *UnsafeUsersServer) mustEmbedUnimplementedUsersServer() {
	_m.Called()
}

// UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedUsersServer'
type UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedUsersServer is a helper method to define mock.On call
func (_e *UnsafeUsersServer_Expecter) mustEmbedUnimplementedUsersServer() *UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call {
	return &UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call{Call: _e.mock.On("mustEmbedUnimplementedUsersServer")}
}

func (_c *UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call) Run(run func()) *UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call) Return() *UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call) RunAndReturn(run func()) *UnsafeUsersServer_mustEmbedUnimplementedUsersServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewUnsafeUsersServer creates a new instance of UnsafeUsersServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeUsersServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeUsersServer {
	mock := &UnsafeUsersServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
