// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	tokenizerapi "github.com/milvus-io/milvus/internal/util/tokenizerapi"
	mock "github.com/stretchr/testify/mock"
)

// Tokenizer is an autogenerated mock type for the Tokenizer type
type Tokenizer struct {
	mock.Mock
}

type Tokenizer_Expecter struct {
	mock *mock.Mock
}

func (_m *Tokenizer) EXPECT() *Tokenizer_Expecter {
	return &Tokenizer_Expecter{mock: &_m.Mock}
}

// Destroy provides a mock function with given fields:
func (_m *Tokenizer) Destroy() {
	_m.Called()
}

// Tokenizer_Destroy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Destroy'
type Tokenizer_Destroy_Call struct {
	*mock.Call
}

// Destroy is a helper method to define mock.On call
func (_e *Tokenizer_Expecter) Destroy() *Tokenizer_Destroy_Call {
	return &Tokenizer_Destroy_Call{Call: _e.mock.On("Destroy")}
}

func (_c *Tokenizer_Destroy_Call) Run(run func()) *Tokenizer_Destroy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tokenizer_Destroy_Call) Return() *Tokenizer_Destroy_Call {
	_c.Call.Return()
	return _c
}

func (_c *Tokenizer_Destroy_Call) RunAndReturn(run func()) *Tokenizer_Destroy_Call {
	_c.Call.Return(run)
	return _c
}

// NewTokenStream provides a mock function with given fields: text
func (_m *Tokenizer) NewTokenStream(text string) tokenizerapi.TokenStream {
	ret := _m.Called(text)

	var r0 tokenizerapi.TokenStream
	if rf, ok := ret.Get(0).(func(string) tokenizerapi.TokenStream); ok {
		r0 = rf(text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tokenizerapi.TokenStream)
		}
	}

	return r0
}

// Tokenizer_NewTokenStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewTokenStream'
type Tokenizer_NewTokenStream_Call struct {
	*mock.Call
}

// NewTokenStream is a helper method to define mock.On call
//   - text string
func (_e *Tokenizer_Expecter) NewTokenStream(text interface{}) *Tokenizer_NewTokenStream_Call {
	return &Tokenizer_NewTokenStream_Call{Call: _e.mock.On("NewTokenStream", text)}
}

func (_c *Tokenizer_NewTokenStream_Call) Run(run func(text string)) *Tokenizer_NewTokenStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Tokenizer_NewTokenStream_Call) Return(_a0 tokenizerapi.TokenStream) *Tokenizer_NewTokenStream_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tokenizer_NewTokenStream_Call) RunAndReturn(run func(string) tokenizerapi.TokenStream) *Tokenizer_NewTokenStream_Call {
	_c.Call.Return(run)
	return _c
}

// NewTokenizer creates a new instance of Tokenizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenizer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tokenizer {
	mock := &Tokenizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
