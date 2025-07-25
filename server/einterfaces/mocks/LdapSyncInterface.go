// Code generated by mockery v2.53.4. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	jobs "github.com/mattermost/mattermost/server/v8/einterfaces/jobs"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost/server/public/model"
)

// LdapSyncInterface is an autogenerated mock type for the LdapSyncInterface type
type LdapSyncInterface struct {
	mock.Mock
}

// MakeScheduler provides a mock function with no fields
func (_m *LdapSyncInterface) MakeScheduler() jobs.Scheduler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MakeScheduler")
	}

	var r0 jobs.Scheduler
	if rf, ok := ret.Get(0).(func() jobs.Scheduler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jobs.Scheduler)
		}
	}

	return r0
}

// MakeWorker provides a mock function with no fields
func (_m *LdapSyncInterface) MakeWorker() model.Worker {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MakeWorker")
	}

	var r0 model.Worker
	if rf, ok := ret.Get(0).(func() model.Worker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Worker)
		}
	}

	return r0
}

// NewLdapSyncInterface creates a new instance of LdapSyncInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLdapSyncInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *LdapSyncInterface {
	mock := &LdapSyncInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
