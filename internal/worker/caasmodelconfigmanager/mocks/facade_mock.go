// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/caasmodelconfigmanager (interfaces: Facade)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/facade_mock.go github.com/juju/juju/internal/worker/caasmodelconfigmanager Facade
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	controller "github.com/juju/juju/controller"
	watcher "github.com/juju/juju/core/watcher"
	gomock "go.uber.org/mock/gomock"
)

// MockFacade is a mock of Facade interface.
type MockFacade struct {
	ctrl     *gomock.Controller
	recorder *MockFacadeMockRecorder
}

// MockFacadeMockRecorder is the mock recorder for MockFacade.
type MockFacadeMockRecorder struct {
	mock *MockFacade
}

// NewMockFacade creates a new mock instance.
func NewMockFacade(ctrl *gomock.Controller) *MockFacade {
	mock := &MockFacade{ctrl: ctrl}
	mock.recorder = &MockFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacade) EXPECT() *MockFacadeMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockFacade) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockFacadeMockRecorder) ControllerConfig(arg0 any) *MockFacadeControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockFacade)(nil).ControllerConfig), arg0)
	return &MockFacadeControllerConfigCall{Call: call}
}

// MockFacadeControllerConfigCall wrap *gomock.Call
type MockFacadeControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeControllerConfigCall) Return(arg0 controller.Config, arg1 error) *MockFacadeControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeControllerConfigCall) Do(f func(context.Context) (controller.Config, error)) *MockFacadeControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeControllerConfigCall) DoAndReturn(f func(context.Context) (controller.Config, error)) *MockFacadeControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchControllerConfig mocks base method.
func (m *MockFacade) WatchControllerConfig() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchControllerConfig")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchControllerConfig indicates an expected call of WatchControllerConfig.
func (mr *MockFacadeMockRecorder) WatchControllerConfig() *MockFacadeWatchControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchControllerConfig", reflect.TypeOf((*MockFacade)(nil).WatchControllerConfig))
	return &MockFacadeWatchControllerConfigCall{Call: call}
}

// MockFacadeWatchControllerConfigCall wrap *gomock.Call
type MockFacadeWatchControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeWatchControllerConfigCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockFacadeWatchControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeWatchControllerConfigCall) Do(f func() (watcher.Watcher[[]string], error)) *MockFacadeWatchControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeWatchControllerConfigCall) DoAndReturn(f func() (watcher.Watcher[[]string], error)) *MockFacadeWatchControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
