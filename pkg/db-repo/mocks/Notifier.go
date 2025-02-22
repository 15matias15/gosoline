// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"
	mock "github.com/stretchr/testify/mock"
)

// Notifier is an autogenerated mock type for the Notifier type
type Notifier struct {
	mock.Mock
}

// Send provides a mock function with given fields: ctx, notificationType, value
func (_m *Notifier) Send(ctx context.Context, notificationType string, value db_repo.ModelBased) error {
	ret := _m.Called(ctx, notificationType, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, notificationType, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
