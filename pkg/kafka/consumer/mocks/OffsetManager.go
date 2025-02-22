// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	kafka "github.com/segmentio/kafka-go"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// OffsetManager is an autogenerated mock type for the OffsetManager type
type OffsetManager struct {
	mock.Mock
}

// Batch provides a mock function with given fields: ctx
func (_m *OffsetManager) Batch(ctx context.Context) []kafka.Message {
	ret := _m.Called(ctx)

	var r0 []kafka.Message
	if rf, ok := ret.Get(0).(func(context.Context) []kafka.Message); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kafka.Message)
		}
	}

	return r0
}

// Commit provides a mock function with given fields: ctx, msgs
func (_m *OffsetManager) Commit(ctx context.Context, msgs ...kafka.Message) error {
	_va := make([]interface{}, len(msgs))
	for _i := range msgs {
		_va[_i] = msgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...kafka.Message) error); ok {
		r0 = rf(ctx, msgs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Flush provides a mock function with given fields:
func (_m *OffsetManager) Flush() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx
func (_m *OffsetManager) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOffsetManager creates a new instance of OffsetManager. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewOffsetManager(t testing.TB) *OffsetManager {
	mock := &OffsetManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
