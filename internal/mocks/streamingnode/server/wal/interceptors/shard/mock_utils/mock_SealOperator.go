// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock_utils

import (
	utils "github.com/milvus-io/milvus/internal/streamingnode/server/wal/interceptors/shard/utils"
	types "github.com/milvus-io/milvus/pkg/v2/streaming/util/types"
	mock "github.com/stretchr/testify/mock"
)

// MockSealOperator is an autogenerated mock type for the SealOperator type
type MockSealOperator struct {
	mock.Mock
}

type MockSealOperator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSealOperator) EXPECT() *MockSealOperator_Expecter {
	return &MockSealOperator_Expecter{mock: &_m.Mock}
}

// AsyncFlushSegment provides a mock function with given fields: infos
func (_m *MockSealOperator) AsyncFlushSegment(infos utils.SealSegmentSignal) {
	_m.Called(infos)
}

// MockSealOperator_AsyncFlushSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AsyncFlushSegment'
type MockSealOperator_AsyncFlushSegment_Call struct {
	*mock.Call
}

// AsyncFlushSegment is a helper method to define mock.On call
//   - infos utils.SealSegmentSignal
func (_e *MockSealOperator_Expecter) AsyncFlushSegment(infos interface{}) *MockSealOperator_AsyncFlushSegment_Call {
	return &MockSealOperator_AsyncFlushSegment_Call{Call: _e.mock.On("AsyncFlushSegment", infos)}
}

func (_c *MockSealOperator_AsyncFlushSegment_Call) Run(run func(infos utils.SealSegmentSignal)) *MockSealOperator_AsyncFlushSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(utils.SealSegmentSignal))
	})
	return _c
}

func (_c *MockSealOperator_AsyncFlushSegment_Call) Return() *MockSealOperator_AsyncFlushSegment_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSealOperator_AsyncFlushSegment_Call) RunAndReturn(run func(utils.SealSegmentSignal)) *MockSealOperator_AsyncFlushSegment_Call {
	_c.Run(run)
	return _c
}

// Channel provides a mock function with no fields
func (_m *MockSealOperator) Channel() types.PChannelInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Channel")
	}

	var r0 types.PChannelInfo
	if rf, ok := ret.Get(0).(func() types.PChannelInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.PChannelInfo)
	}

	return r0
}

// MockSealOperator_Channel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Channel'
type MockSealOperator_Channel_Call struct {
	*mock.Call
}

// Channel is a helper method to define mock.On call
func (_e *MockSealOperator_Expecter) Channel() *MockSealOperator_Channel_Call {
	return &MockSealOperator_Channel_Call{Call: _e.mock.On("Channel")}
}

func (_c *MockSealOperator_Channel_Call) Run(run func()) *MockSealOperator_Channel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSealOperator_Channel_Call) Return(_a0 types.PChannelInfo) *MockSealOperator_Channel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSealOperator_Channel_Call) RunAndReturn(run func() types.PChannelInfo) *MockSealOperator_Channel_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSealOperator creates a new instance of MockSealOperator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSealOperator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSealOperator {
	mock := &MockSealOperator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
