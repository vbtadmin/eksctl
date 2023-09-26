// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	aws "github.com/aws/aws-sdk-go/aws"
	client "github.com/aws/aws-sdk-go/aws/client"

	mock "github.com/stretchr/testify/mock"
)

// ConfigProvider is an autogenerated mock type for the ConfigProvider type
type ConfigProvider struct {
	mock.Mock
}

// ClientConfig provides a mock function with given fields: serviceName, cfgs
func (_m *ConfigProvider) ClientConfig(serviceName string, cfgs ...*aws.Config) client.Config {
	_va := make([]interface{}, len(cfgs))
	for _i := range cfgs {
		_va[_i] = cfgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, serviceName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 client.Config
	if rf, ok := ret.Get(0).(func(string, ...*aws.Config) client.Config); ok {
		r0 = rf(serviceName, cfgs...)
	} else {
		r0 = ret.Get(0).(client.Config)
	}

	return r0
}

// NewConfigProvider creates a new instance of ConfigProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigProvider {
	mock := &ConfigProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
