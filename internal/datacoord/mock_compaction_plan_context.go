// Code generated by mockery v2.46.0. DO NOT EDIT.

package datacoord

import (
	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"
)

// MockCompactionPlanContext is an autogenerated mock type for the compactionPlanContext type
type MockCompactionPlanContext struct {
	mock.Mock
}

type MockCompactionPlanContext_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCompactionPlanContext) EXPECT() *MockCompactionPlanContext_Expecter {
	return &MockCompactionPlanContext_Expecter{mock: &_m.Mock}
}

// enqueueCompaction provides a mock function with given fields: task
func (_m *MockCompactionPlanContext) enqueueCompaction(task *datapb.CompactionTask) error {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for enqueueCompaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*datapb.CompactionTask) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCompactionPlanContext_enqueueCompaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'enqueueCompaction'
type MockCompactionPlanContext_enqueueCompaction_Call struct {
	*mock.Call
}

// enqueueCompaction is a helper method to define mock.On call
//   - task *datapb.CompactionTask
func (_e *MockCompactionPlanContext_Expecter) enqueueCompaction(task interface{}) *MockCompactionPlanContext_enqueueCompaction_Call {
	return &MockCompactionPlanContext_enqueueCompaction_Call{Call: _e.mock.On("enqueueCompaction", task)}
}

func (_c *MockCompactionPlanContext_enqueueCompaction_Call) Run(run func(task *datapb.CompactionTask)) *MockCompactionPlanContext_enqueueCompaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*datapb.CompactionTask))
	})
	return _c
}

func (_c *MockCompactionPlanContext_enqueueCompaction_Call) Return(_a0 error) *MockCompactionPlanContext_enqueueCompaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_enqueueCompaction_Call) RunAndReturn(run func(*datapb.CompactionTask) error) *MockCompactionPlanContext_enqueueCompaction_Call {
	_c.Call.Return(run)
	return _c
}

// getCompactionInfo provides a mock function with given fields: signalID
func (_m *MockCompactionPlanContext) getCompactionInfo(signalID int64) *compactionInfo {
	ret := _m.Called(signalID)

	if len(ret) == 0 {
		panic("no return value specified for getCompactionInfo")
	}

	var r0 *compactionInfo
	if rf, ok := ret.Get(0).(func(int64) *compactionInfo); ok {
		r0 = rf(signalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compactionInfo)
		}
	}

	return r0
}

// MockCompactionPlanContext_getCompactionInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getCompactionInfo'
type MockCompactionPlanContext_getCompactionInfo_Call struct {
	*mock.Call
}

// getCompactionInfo is a helper method to define mock.On call
//   - signalID int64
func (_e *MockCompactionPlanContext_Expecter) getCompactionInfo(signalID interface{}) *MockCompactionPlanContext_getCompactionInfo_Call {
	return &MockCompactionPlanContext_getCompactionInfo_Call{Call: _e.mock.On("getCompactionInfo", signalID)}
}

func (_c *MockCompactionPlanContext_getCompactionInfo_Call) Run(run func(signalID int64)) *MockCompactionPlanContext_getCompactionInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCompactionPlanContext_getCompactionInfo_Call) Return(_a0 *compactionInfo) *MockCompactionPlanContext_getCompactionInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_getCompactionInfo_Call) RunAndReturn(run func(int64) *compactionInfo) *MockCompactionPlanContext_getCompactionInfo_Call {
	_c.Call.Return(run)
	return _c
}

// getCompactionTasksNumBySignalID provides a mock function with given fields: signalID
func (_m *MockCompactionPlanContext) getCompactionTasksNumBySignalID(signalID int64) int {
	ret := _m.Called(signalID)

	if len(ret) == 0 {
		panic("no return value specified for getCompactionTasksNumBySignalID")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(int64) int); ok {
		r0 = rf(signalID)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getCompactionTasksNumBySignalID'
type MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call struct {
	*mock.Call
}

// getCompactionTasksNumBySignalID is a helper method to define mock.On call
//   - signalID int64
func (_e *MockCompactionPlanContext_Expecter) getCompactionTasksNumBySignalID(signalID interface{}) *MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call {
	return &MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call{Call: _e.mock.On("getCompactionTasksNumBySignalID", signalID)}
}

func (_c *MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call) Run(run func(signalID int64)) *MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call) Return(_a0 int) *MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call) RunAndReturn(run func(int64) int) *MockCompactionPlanContext_getCompactionTasksNumBySignalID_Call {
	_c.Call.Return(run)
	return _c
}

// isFull provides a mock function with given fields:
func (_m *MockCompactionPlanContext) isFull() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for isFull")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCompactionPlanContext_isFull_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isFull'
type MockCompactionPlanContext_isFull_Call struct {
	*mock.Call
}

// isFull is a helper method to define mock.On call
func (_e *MockCompactionPlanContext_Expecter) isFull() *MockCompactionPlanContext_isFull_Call {
	return &MockCompactionPlanContext_isFull_Call{Call: _e.mock.On("isFull")}
}

func (_c *MockCompactionPlanContext_isFull_Call) Run(run func()) *MockCompactionPlanContext_isFull_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionPlanContext_isFull_Call) Return(_a0 bool) *MockCompactionPlanContext_isFull_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_isFull_Call) RunAndReturn(run func() bool) *MockCompactionPlanContext_isFull_Call {
	_c.Call.Return(run)
	return _c
}

// removeTasksByChannel provides a mock function with given fields: channel
func (_m *MockCompactionPlanContext) removeTasksByChannel(channel string) {
	_m.Called(channel)
}

// MockCompactionPlanContext_removeTasksByChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'removeTasksByChannel'
type MockCompactionPlanContext_removeTasksByChannel_Call struct {
	*mock.Call
}

// removeTasksByChannel is a helper method to define mock.On call
//   - channel string
func (_e *MockCompactionPlanContext_Expecter) removeTasksByChannel(channel interface{}) *MockCompactionPlanContext_removeTasksByChannel_Call {
	return &MockCompactionPlanContext_removeTasksByChannel_Call{Call: _e.mock.On("removeTasksByChannel", channel)}
}

func (_c *MockCompactionPlanContext_removeTasksByChannel_Call) Run(run func(channel string)) *MockCompactionPlanContext_removeTasksByChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCompactionPlanContext_removeTasksByChannel_Call) Return() *MockCompactionPlanContext_removeTasksByChannel_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCompactionPlanContext_removeTasksByChannel_Call) RunAndReturn(run func(string)) *MockCompactionPlanContext_removeTasksByChannel_Call {
	_c.Call.Return(run)
	return _c
}

// start provides a mock function with given fields:
func (_m *MockCompactionPlanContext) start() {
	_m.Called()
}

// MockCompactionPlanContext_start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'start'
type MockCompactionPlanContext_start_Call struct {
	*mock.Call
}

// start is a helper method to define mock.On call
func (_e *MockCompactionPlanContext_Expecter) start() *MockCompactionPlanContext_start_Call {
	return &MockCompactionPlanContext_start_Call{Call: _e.mock.On("start")}
}

func (_c *MockCompactionPlanContext_start_Call) Run(run func()) *MockCompactionPlanContext_start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionPlanContext_start_Call) Return() *MockCompactionPlanContext_start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCompactionPlanContext_start_Call) RunAndReturn(run func()) *MockCompactionPlanContext_start_Call {
	_c.Call.Return(run)
	return _c
}

// stop provides a mock function with given fields:
func (_m *MockCompactionPlanContext) stop() {
	_m.Called()
}

// MockCompactionPlanContext_stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'stop'
type MockCompactionPlanContext_stop_Call struct {
	*mock.Call
}

// stop is a helper method to define mock.On call
func (_e *MockCompactionPlanContext_Expecter) stop() *MockCompactionPlanContext_stop_Call {
	return &MockCompactionPlanContext_stop_Call{Call: _e.mock.On("stop")}
}

func (_c *MockCompactionPlanContext_stop_Call) Run(run func()) *MockCompactionPlanContext_stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionPlanContext_stop_Call) Return() *MockCompactionPlanContext_stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCompactionPlanContext_stop_Call) RunAndReturn(run func()) *MockCompactionPlanContext_stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCompactionPlanContext creates a new instance of MockCompactionPlanContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCompactionPlanContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCompactionPlanContext {
	mock := &MockCompactionPlanContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
