// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ConfigProviderInterface is an autogenerated mock type for the ConfigProviderInterface type
type ConfigProviderInterface struct {
	mock.Mock
}

type ConfigProviderInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ConfigProviderInterface) EXPECT() *ConfigProviderInterface_Expecter {
	return &ConfigProviderInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: key
func (_m *ConfigProviderInterface) Get(key string) interface{} {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// ConfigProviderInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ConfigProviderInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *ConfigProviderInterface_Expecter) Get(key interface{}) *ConfigProviderInterface_Get_Call {
	return &ConfigProviderInterface_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *ConfigProviderInterface_Get_Call) Run(run func(key string)) *ConfigProviderInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ConfigProviderInterface_Get_Call) Return(_a0 interface{}) *ConfigProviderInterface_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigProviderInterface_Get_Call) RunAndReturn(run func(string) interface{}) *ConfigProviderInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetBool provides a mock function with given fields: key
func (_m *ConfigProviderInterface) GetBool(key string) bool {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetBool")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ConfigProviderInterface_GetBool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBool'
type ConfigProviderInterface_GetBool_Call struct {
	*mock.Call
}

// GetBool is a helper method to define mock.On call
//   - key string
func (_e *ConfigProviderInterface_Expecter) GetBool(key interface{}) *ConfigProviderInterface_GetBool_Call {
	return &ConfigProviderInterface_GetBool_Call{Call: _e.mock.On("GetBool", key)}
}

func (_c *ConfigProviderInterface_GetBool_Call) Run(run func(key string)) *ConfigProviderInterface_GetBool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ConfigProviderInterface_GetBool_Call) Return(_a0 bool) *ConfigProviderInterface_GetBool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigProviderInterface_GetBool_Call) RunAndReturn(run func(string) bool) *ConfigProviderInterface_GetBool_Call {
	_c.Call.Return(run)
	return _c
}

// GetInt provides a mock function with given fields: key
func (_m *ConfigProviderInterface) GetInt(key string) int {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetInt")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ConfigProviderInterface_GetInt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInt'
type ConfigProviderInterface_GetInt_Call struct {
	*mock.Call
}

// GetInt is a helper method to define mock.On call
//   - key string
func (_e *ConfigProviderInterface_Expecter) GetInt(key interface{}) *ConfigProviderInterface_GetInt_Call {
	return &ConfigProviderInterface_GetInt_Call{Call: _e.mock.On("GetInt", key)}
}

func (_c *ConfigProviderInterface_GetInt_Call) Run(run func(key string)) *ConfigProviderInterface_GetInt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ConfigProviderInterface_GetInt_Call) Return(_a0 int) *ConfigProviderInterface_GetInt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigProviderInterface_GetInt_Call) RunAndReturn(run func(string) int) *ConfigProviderInterface_GetInt_Call {
	_c.Call.Return(run)
	return _c
}

// GetString provides a mock function with given fields: key
func (_m *ConfigProviderInterface) GetString(key string) string {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetString")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ConfigProviderInterface_GetString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetString'
type ConfigProviderInterface_GetString_Call struct {
	*mock.Call
}

// GetString is a helper method to define mock.On call
//   - key string
func (_e *ConfigProviderInterface_Expecter) GetString(key interface{}) *ConfigProviderInterface_GetString_Call {
	return &ConfigProviderInterface_GetString_Call{Call: _e.mock.On("GetString", key)}
}

func (_c *ConfigProviderInterface_GetString_Call) Run(run func(key string)) *ConfigProviderInterface_GetString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ConfigProviderInterface_GetString_Call) Return(_a0 string) *ConfigProviderInterface_GetString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigProviderInterface_GetString_Call) RunAndReturn(run func(string) string) *ConfigProviderInterface_GetString_Call {
	_c.Call.Return(run)
	return _c
}

// GetStringMap provides a mock function with given fields: key
func (_m *ConfigProviderInterface) GetStringMap(key string) map[string]interface{} {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetStringMap")
	}

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// ConfigProviderInterface_GetStringMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStringMap'
type ConfigProviderInterface_GetStringMap_Call struct {
	*mock.Call
}

// GetStringMap is a helper method to define mock.On call
//   - key string
func (_e *ConfigProviderInterface_Expecter) GetStringMap(key interface{}) *ConfigProviderInterface_GetStringMap_Call {
	return &ConfigProviderInterface_GetStringMap_Call{Call: _e.mock.On("GetStringMap", key)}
}

func (_c *ConfigProviderInterface_GetStringMap_Call) Run(run func(key string)) *ConfigProviderInterface_GetStringMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ConfigProviderInterface_GetStringMap_Call) Return(_a0 map[string]interface{}) *ConfigProviderInterface_GetStringMap_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigProviderInterface_GetStringMap_Call) RunAndReturn(run func(string) map[string]interface{}) *ConfigProviderInterface_GetStringMap_Call {
	_c.Call.Return(run)
	return _c
}

// NewConfigProviderInterface creates a new instance of ConfigProviderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigProviderInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigProviderInterface {
	mock := &ConfigProviderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
