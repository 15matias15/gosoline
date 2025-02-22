// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	stream "github.com/justtrackio/gosoline/pkg/stream"
	mock "github.com/stretchr/testify/mock"
)

// PartitionedOutput is an autogenerated mock type for the PartitionedOutput type
type PartitionedOutput struct {
	mock.Mock
}

// IsPartitionedOutput provides a mock function with given fields:
func (_m *PartitionedOutput) IsPartitionedOutput() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Write provides a mock function with given fields: ctx, batch
func (_m *PartitionedOutput) Write(ctx context.Context, batch []stream.WritableMessage) error {
	ret := _m.Called(ctx, batch)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []stream.WritableMessage) error); ok {
		r0 = rf(ctx, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteOne provides a mock function with given fields: ctx, msg
func (_m *PartitionedOutput) WriteOne(ctx context.Context, msg stream.WritableMessage) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, stream.WritableMessage) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
