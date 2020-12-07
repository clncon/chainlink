// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	context "context"

	pipeline "github.com/smartcontractkit/chainlink/core/services/pipeline"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Task is an autogenerated mock type for the Task type
type Task struct {
	mock.Mock
}

// DotID provides a mock function with given fields:
func (_m *Task) DotID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OutputIndex provides a mock function with given fields:
func (_m *Task) OutputIndex() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// OutputTask provides a mock function with given fields:
func (_m *Task) OutputTask() pipeline.Task {
	ret := _m.Called()

	var r0 pipeline.Task
	if rf, ok := ret.Get(0).(func() pipeline.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pipeline.Task)
		}
	}

	return r0
}

// Run provides a mock function with given fields: ctx, taskRun, inputs
func (_m *Task) Run(ctx context.Context, taskRun pipeline.TaskRun, inputs []pipeline.Result) pipeline.Result {
	ret := _m.Called(ctx, taskRun, inputs)

	var r0 pipeline.Result
	if rf, ok := ret.Get(0).(func(context.Context, pipeline.TaskRun, []pipeline.Result) pipeline.Result); ok {
		r0 = rf(ctx, taskRun, inputs)
	} else {
		r0 = ret.Get(0).(pipeline.Result)
	}

	return r0
}

// SetOutputTask provides a mock function with given fields: task
func (_m *Task) SetOutputTask(task pipeline.Task) {
	_m.Called(task)
}

// TaskTimeout provides a mock function with given fields:
func (_m *Task) TaskTimeout() (time.Duration, bool) {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Type provides a mock function with given fields:
func (_m *Task) Type() pipeline.TaskType {
	ret := _m.Called()

	var r0 pipeline.TaskType
	if rf, ok := ret.Get(0).(func() pipeline.TaskType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pipeline.TaskType)
	}

	return r0
}