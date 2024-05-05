// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/firewaller (interfaces: State,ControllerConfigAPI)
//
// Generated by this command:
//
//	mockgen -package firewaller_test -destination package_mock_test.go github.com/juju/juju/apiserver/facades/controller/firewaller State,ControllerConfigAPI
//

// Package firewaller_test is a generated GoMock package.
package firewaller_test

import (
	context "context"
	reflect "reflect"
	time "time"

	firewall "github.com/juju/juju/apiserver/common/firewall"
	params "github.com/juju/juju/rpc/params"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
	macaroon "gopkg.in/macaroon.v2"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// AllEndpointBindings mocks base method.
func (m *MockState) AllEndpointBindings() (map[string]map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllEndpointBindings")
	ret0, _ := ret[0].(map[string]map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllEndpointBindings indicates an expected call of AllEndpointBindings.
func (mr *MockStateMockRecorder) AllEndpointBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllEndpointBindings", reflect.TypeOf((*MockState)(nil).AllEndpointBindings))
}

// Application mocks base method.
func (m *MockState) Application(arg0 string) (firewall.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(firewall.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockStateMockRecorder) Application(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockState)(nil).Application), arg0)
}

// FindEntity mocks base method.
func (m *MockState) FindEntity(arg0 names.Tag) (state.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntity", arg0)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntity indicates an expected call of FindEntity.
func (mr *MockStateMockRecorder) FindEntity(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntity", reflect.TypeOf((*MockState)(nil).FindEntity), arg0)
}

// GetMacaroon mocks base method.
func (m *MockState) GetMacaroon(arg0 names.Tag) (*macaroon.Macaroon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMacaroon", arg0)
	ret0, _ := ret[0].(*macaroon.Macaroon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMacaroon indicates an expected call of GetMacaroon.
func (mr *MockStateMockRecorder) GetMacaroon(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMacaroon", reflect.TypeOf((*MockState)(nil).GetMacaroon), arg0)
}

// IsController mocks base method.
func (m *MockState) IsController() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsController")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsController indicates an expected call of IsController.
func (mr *MockStateMockRecorder) IsController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsController", reflect.TypeOf((*MockState)(nil).IsController))
}

// KeyRelation mocks base method.
func (m *MockState) KeyRelation(arg0 string) (firewall.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyRelation", arg0)
	ret0, _ := ret[0].(firewall.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeyRelation indicates an expected call of KeyRelation.
func (mr *MockStateMockRecorder) KeyRelation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyRelation", reflect.TypeOf((*MockState)(nil).KeyRelation), arg0)
}

// Machine mocks base method.
func (m *MockState) Machine(arg0 string) (firewall.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(firewall.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockStateMockRecorder) Machine(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockState)(nil).Machine), arg0)
}

// ModelUUID mocks base method.
func (m *MockState) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockStateMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockState)(nil).ModelUUID))
}

// Unit mocks base method.
func (m *MockState) Unit(arg0 string) (firewall.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(firewall.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockStateMockRecorder) Unit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockState)(nil).Unit), arg0)
}

// WatchModelMachineStartTimes mocks base method.
func (m *MockState) WatchModelMachineStartTimes(arg0 time.Duration) state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchModelMachineStartTimes", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchModelMachineStartTimes indicates an expected call of WatchModelMachineStartTimes.
func (mr *MockStateMockRecorder) WatchModelMachineStartTimes(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchModelMachineStartTimes", reflect.TypeOf((*MockState)(nil).WatchModelMachineStartTimes), arg0)
}

// WatchModelMachines mocks base method.
func (m *MockState) WatchModelMachines() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchModelMachines")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchModelMachines indicates an expected call of WatchModelMachines.
func (mr *MockStateMockRecorder) WatchModelMachines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchModelMachines", reflect.TypeOf((*MockState)(nil).WatchModelMachines))
}

// WatchOpenedPorts mocks base method.
func (m *MockState) WatchOpenedPorts() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOpenedPorts")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchOpenedPorts indicates an expected call of WatchOpenedPorts.
func (mr *MockStateMockRecorder) WatchOpenedPorts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOpenedPorts", reflect.TypeOf((*MockState)(nil).WatchOpenedPorts))
}

// MockControllerConfigAPI is a mock of ControllerConfigAPI interface.
type MockControllerConfigAPI struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigAPIMockRecorder
}

// MockControllerConfigAPIMockRecorder is the mock recorder for MockControllerConfigAPI.
type MockControllerConfigAPIMockRecorder struct {
	mock *MockControllerConfigAPI
}

// NewMockControllerConfigAPI creates a new mock instance.
func NewMockControllerConfigAPI(ctrl *gomock.Controller) *MockControllerConfigAPI {
	mock := &MockControllerConfigAPI{ctrl: ctrl}
	mock.recorder = &MockControllerConfigAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigAPI) EXPECT() *MockControllerConfigAPIMockRecorder {
	return m.recorder
}

// ControllerAPIInfoForModels mocks base method.
func (m *MockControllerConfigAPI) ControllerAPIInfoForModels(arg0 context.Context, arg1 params.Entities) (params.ControllerAPIInfoResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerAPIInfoForModels", arg0, arg1)
	ret0, _ := ret[0].(params.ControllerAPIInfoResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerAPIInfoForModels indicates an expected call of ControllerAPIInfoForModels.
func (mr *MockControllerConfigAPIMockRecorder) ControllerAPIInfoForModels(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerAPIInfoForModels", reflect.TypeOf((*MockControllerConfigAPI)(nil).ControllerAPIInfoForModels), arg0, arg1)
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigAPI) ControllerConfig(arg0 context.Context) (params.ControllerConfigResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(params.ControllerConfigResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigAPIMockRecorder) ControllerConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigAPI)(nil).ControllerConfig), arg0)
}
