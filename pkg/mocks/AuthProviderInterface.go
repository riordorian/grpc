// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	jwt "github.com/golang-jwt/jwt"

	mock "github.com/stretchr/testify/mock"
)

// AuthProviderInterface is an autogenerated mock type for the AuthProviderInterface type
type AuthProviderInterface struct {
	mock.Mock
}

type AuthProviderInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthProviderInterface) EXPECT() *AuthProviderInterface_Expecter {
	return &AuthProviderInterface_Expecter{mock: &_m.Mock}
}

// Can provides a mock function with given fields: action, accessToken
func (_m *AuthProviderInterface) Can(action string, accessToken string) (bool, error) {
	ret := _m.Called(action, accessToken)

	if len(ret) == 0 {
		panic("no return value specified for Can")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(action, accessToken)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(action, accessToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(action, accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthProviderInterface_Can_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Can'
type AuthProviderInterface_Can_Call struct {
	*mock.Call
}

// Can is a helper method to define mock.On call
//   - action string
//   - accessToken string
func (_e *AuthProviderInterface_Expecter) Can(action interface{}, accessToken interface{}) *AuthProviderInterface_Can_Call {
	return &AuthProviderInterface_Can_Call{Call: _e.mock.On("Can", action, accessToken)}
}

func (_c *AuthProviderInterface_Can_Call) Run(run func(action string, accessToken string)) *AuthProviderInterface_Can_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *AuthProviderInterface_Can_Call) Return(_a0 bool, _a1 error) *AuthProviderInterface_Can_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthProviderInterface_Can_Call) RunAndReturn(run func(string, string) (bool, error)) *AuthProviderInterface_Can_Call {
	_c.Call.Return(run)
	return _c
}

// CheckLogin provides a mock function with given fields: accessToken
func (_m *AuthProviderInterface) CheckLogin(accessToken string) (bool, error) {
	ret := _m.Called(accessToken)

	if len(ret) == 0 {
		panic("no return value specified for CheckLogin")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(accessToken)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(accessToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthProviderInterface_CheckLogin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckLogin'
type AuthProviderInterface_CheckLogin_Call struct {
	*mock.Call
}

// CheckLogin is a helper method to define mock.On call
//   - accessToken string
func (_e *AuthProviderInterface_Expecter) CheckLogin(accessToken interface{}) *AuthProviderInterface_CheckLogin_Call {
	return &AuthProviderInterface_CheckLogin_Call{Call: _e.mock.On("CheckLogin", accessToken)}
}

func (_c *AuthProviderInterface_CheckLogin_Call) Run(run func(accessToken string)) *AuthProviderInterface_CheckLogin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AuthProviderInterface_CheckLogin_Call) Return(_a0 bool, _a1 error) *AuthProviderInterface_CheckLogin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthProviderInterface_CheckLogin_Call) RunAndReturn(run func(string) (bool, error)) *AuthProviderInterface_CheckLogin_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: ctx, login, password
func (_m *AuthProviderInterface) Login(ctx context.Context, login string, password string) (jwt.Token, error) {
	ret := _m.Called(ctx, login, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (jwt.Token, error)); ok {
		return rf(ctx, login, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) jwt.Token); ok {
		r0 = rf(ctx, login, password)
	} else {
		r0 = ret.Get(0).(jwt.Token)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, login, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthProviderInterface_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type AuthProviderInterface_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - login string
//   - password string
func (_e *AuthProviderInterface_Expecter) Login(ctx interface{}, login interface{}, password interface{}) *AuthProviderInterface_Login_Call {
	return &AuthProviderInterface_Login_Call{Call: _e.mock.On("Login", ctx, login, password)}
}

func (_c *AuthProviderInterface_Login_Call) Run(run func(ctx context.Context, login string, password string)) *AuthProviderInterface_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AuthProviderInterface_Login_Call) Return(_a0 jwt.Token, _a1 error) *AuthProviderInterface_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthProviderInterface_Login_Call) RunAndReturn(run func(context.Context, string, string) (jwt.Token, error)) *AuthProviderInterface_Login_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthProviderInterface creates a new instance of AuthProviderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthProviderInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthProviderInterface {
	mock := &AuthProviderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
