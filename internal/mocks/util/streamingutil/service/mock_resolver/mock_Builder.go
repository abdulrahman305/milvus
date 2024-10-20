// Code generated by mockery v2.46.0. DO NOT EDIT.

package mock_resolver

import (
	mock "github.com/stretchr/testify/mock"
	resolver "google.golang.org/grpc/resolver"

	serviceresolver "github.com/milvus-io/milvus/internal/util/streamingutil/service/resolver"
)

// MockBuilder is an autogenerated mock type for the Builder type
type MockBuilder struct {
	mock.Mock
}

type MockBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBuilder) EXPECT() *MockBuilder_Expecter {
	return &MockBuilder_Expecter{mock: &_m.Mock}
}

// Build provides a mock function with given fields: target, cc, opts
func (_m *MockBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ret := _m.Called(target, cc, opts)

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 resolver.Resolver
	var r1 error
	if rf, ok := ret.Get(0).(func(resolver.Target, resolver.ClientConn, resolver.BuildOptions) (resolver.Resolver, error)); ok {
		return rf(target, cc, opts)
	}
	if rf, ok := ret.Get(0).(func(resolver.Target, resolver.ClientConn, resolver.BuildOptions) resolver.Resolver); ok {
		r0 = rf(target, cc, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(resolver.Resolver)
		}
	}

	if rf, ok := ret.Get(1).(func(resolver.Target, resolver.ClientConn, resolver.BuildOptions) error); ok {
		r1 = rf(target, cc, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBuilder_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type MockBuilder_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
//   - target resolver.Target
//   - cc resolver.ClientConn
//   - opts resolver.BuildOptions
func (_e *MockBuilder_Expecter) Build(target interface{}, cc interface{}, opts interface{}) *MockBuilder_Build_Call {
	return &MockBuilder_Build_Call{Call: _e.mock.On("Build", target, cc, opts)}
}

func (_c *MockBuilder_Build_Call) Run(run func(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions)) *MockBuilder_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(resolver.Target), args[1].(resolver.ClientConn), args[2].(resolver.BuildOptions))
	})
	return _c
}

func (_c *MockBuilder_Build_Call) Return(_a0 resolver.Resolver, _a1 error) *MockBuilder_Build_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBuilder_Build_Call) RunAndReturn(run func(resolver.Target, resolver.ClientConn, resolver.BuildOptions) (resolver.Resolver, error)) *MockBuilder_Build_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockBuilder) Close() {
	_m.Called()
}

// MockBuilder_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockBuilder_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockBuilder_Expecter) Close() *MockBuilder_Close_Call {
	return &MockBuilder_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockBuilder_Close_Call) Run(run func()) *MockBuilder_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBuilder_Close_Call) Return() *MockBuilder_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBuilder_Close_Call) RunAndReturn(run func()) *MockBuilder_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Resolver provides a mock function with given fields:
func (_m *MockBuilder) Resolver() serviceresolver.Resolver {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Resolver")
	}

	var r0 serviceresolver.Resolver
	if rf, ok := ret.Get(0).(func() serviceresolver.Resolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(serviceresolver.Resolver)
		}
	}

	return r0
}

// MockBuilder_Resolver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resolver'
type MockBuilder_Resolver_Call struct {
	*mock.Call
}

// Resolver is a helper method to define mock.On call
func (_e *MockBuilder_Expecter) Resolver() *MockBuilder_Resolver_Call {
	return &MockBuilder_Resolver_Call{Call: _e.mock.On("Resolver")}
}

func (_c *MockBuilder_Resolver_Call) Run(run func()) *MockBuilder_Resolver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBuilder_Resolver_Call) Return(_a0 serviceresolver.Resolver) *MockBuilder_Resolver_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBuilder_Resolver_Call) RunAndReturn(run func() serviceresolver.Resolver) *MockBuilder_Resolver_Call {
	_c.Call.Return(run)
	return _c
}

// Scheme provides a mock function with given fields:
func (_m *MockBuilder) Scheme() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Scheme")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockBuilder_Scheme_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Scheme'
type MockBuilder_Scheme_Call struct {
	*mock.Call
}

// Scheme is a helper method to define mock.On call
func (_e *MockBuilder_Expecter) Scheme() *MockBuilder_Scheme_Call {
	return &MockBuilder_Scheme_Call{Call: _e.mock.On("Scheme")}
}

func (_c *MockBuilder_Scheme_Call) Run(run func()) *MockBuilder_Scheme_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBuilder_Scheme_Call) Return(_a0 string) *MockBuilder_Scheme_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBuilder_Scheme_Call) RunAndReturn(run func() string) *MockBuilder_Scheme_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBuilder creates a new instance of MockBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBuilder {
	mock := &MockBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
