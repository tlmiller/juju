// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/uniter (interfaces: LXDProfileBackend,LXDProfileMachine,LXDProfileUnit,LXDProfileBackendV2,LXDProfileMachineV2,LXDProfileUnitV2,LXDProfileCharmV2)
//
// Generated by this command:
//
//	mockgen -typed -package uniter_test -destination package_mocks_test.go github.com/juju/juju/apiserver/facades/agent/uniter LXDProfileBackend,LXDProfileMachine,LXDProfileUnit,LXDProfileBackendV2,LXDProfileMachineV2,LXDProfileUnitV2,LXDProfileCharmV2
//

// Package uniter_test is a generated GoMock package.
package uniter_test

import (
	reflect "reflect"

	uniter "github.com/juju/juju/apiserver/facades/agent/uniter"
	instance "github.com/juju/juju/core/instance"
	lxdprofile "github.com/juju/juju/core/lxdprofile"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v6"
	gomock "go.uber.org/mock/gomock"
)

// MockLXDProfileBackend is a mock of LXDProfileBackend interface.
type MockLXDProfileBackend struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileBackendMockRecorder
}

// MockLXDProfileBackendMockRecorder is the mock recorder for MockLXDProfileBackend.
type MockLXDProfileBackendMockRecorder struct {
	mock *MockLXDProfileBackend
}

// NewMockLXDProfileBackend creates a new mock instance.
func NewMockLXDProfileBackend(ctrl *gomock.Controller) *MockLXDProfileBackend {
	mock := &MockLXDProfileBackend{ctrl: ctrl}
	mock.recorder = &MockLXDProfileBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileBackend) EXPECT() *MockLXDProfileBackendMockRecorder {
	return m.recorder
}

// Machine mocks base method.
func (m *MockLXDProfileBackend) Machine(arg0 string) (uniter.LXDProfileMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockLXDProfileBackendMockRecorder) Machine(arg0 any) *MockLXDProfileBackendMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockLXDProfileBackend)(nil).Machine), arg0)
	return &MockLXDProfileBackendMachineCall{Call: call}
}

// MockLXDProfileBackendMachineCall wrap *gomock.Call
type MockLXDProfileBackendMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileBackendMachineCall) Return(arg0 uniter.LXDProfileMachine, arg1 error) *MockLXDProfileBackendMachineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileBackendMachineCall) Do(f func(string) (uniter.LXDProfileMachine, error)) *MockLXDProfileBackendMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileBackendMachineCall) DoAndReturn(f func(string) (uniter.LXDProfileMachine, error)) *MockLXDProfileBackendMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Unit mocks base method.
func (m *MockLXDProfileBackend) Unit(arg0 string) (uniter.LXDProfileUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockLXDProfileBackendMockRecorder) Unit(arg0 any) *MockLXDProfileBackendUnitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockLXDProfileBackend)(nil).Unit), arg0)
	return &MockLXDProfileBackendUnitCall{Call: call}
}

// MockLXDProfileBackendUnitCall wrap *gomock.Call
type MockLXDProfileBackendUnitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileBackendUnitCall) Return(arg0 uniter.LXDProfileUnit, arg1 error) *MockLXDProfileBackendUnitCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileBackendUnitCall) Do(f func(string) (uniter.LXDProfileUnit, error)) *MockLXDProfileBackendUnitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileBackendUnitCall) DoAndReturn(f func(string) (uniter.LXDProfileUnit, error)) *MockLXDProfileBackendUnitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLXDProfileMachine is a mock of LXDProfileMachine interface.
type MockLXDProfileMachine struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileMachineMockRecorder
}

// MockLXDProfileMachineMockRecorder is the mock recorder for MockLXDProfileMachine.
type MockLXDProfileMachineMockRecorder struct {
	mock *MockLXDProfileMachine
}

// NewMockLXDProfileMachine creates a new mock instance.
func NewMockLXDProfileMachine(ctrl *gomock.Controller) *MockLXDProfileMachine {
	mock := &MockLXDProfileMachine{ctrl: ctrl}
	mock.recorder = &MockLXDProfileMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileMachine) EXPECT() *MockLXDProfileMachineMockRecorder {
	return m.recorder
}

// WatchLXDProfileUpgradeNotifications mocks base method.
func (m *MockLXDProfileMachine) WatchLXDProfileUpgradeNotifications(arg0 string) (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchLXDProfileUpgradeNotifications", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchLXDProfileUpgradeNotifications indicates an expected call of WatchLXDProfileUpgradeNotifications.
func (mr *MockLXDProfileMachineMockRecorder) WatchLXDProfileUpgradeNotifications(arg0 any) *MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLXDProfileUpgradeNotifications", reflect.TypeOf((*MockLXDProfileMachine)(nil).WatchLXDProfileUpgradeNotifications), arg0)
	return &MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall{Call: call}
}

// MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall wrap *gomock.Call
type MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall) Return(arg0 state.StringsWatcher, arg1 error) *MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall) Do(f func(string) (state.StringsWatcher, error)) *MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall) DoAndReturn(f func(string) (state.StringsWatcher, error)) *MockLXDProfileMachineWatchLXDProfileUpgradeNotificationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLXDProfileUnit is a mock of LXDProfileUnit interface.
type MockLXDProfileUnit struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileUnitMockRecorder
}

// MockLXDProfileUnitMockRecorder is the mock recorder for MockLXDProfileUnit.
type MockLXDProfileUnitMockRecorder struct {
	mock *MockLXDProfileUnit
}

// NewMockLXDProfileUnit creates a new mock instance.
func NewMockLXDProfileUnit(ctrl *gomock.Controller) *MockLXDProfileUnit {
	mock := &MockLXDProfileUnit{ctrl: ctrl}
	mock.recorder = &MockLXDProfileUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileUnit) EXPECT() *MockLXDProfileUnitMockRecorder {
	return m.recorder
}

// AssignedMachineId mocks base method.
func (m *MockLXDProfileUnit) AssignedMachineId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId.
func (mr *MockLXDProfileUnitMockRecorder) AssignedMachineId() *MockLXDProfileUnitAssignedMachineIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockLXDProfileUnit)(nil).AssignedMachineId))
	return &MockLXDProfileUnitAssignedMachineIdCall{Call: call}
}

// MockLXDProfileUnitAssignedMachineIdCall wrap *gomock.Call
type MockLXDProfileUnitAssignedMachineIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitAssignedMachineIdCall) Return(arg0 string, arg1 error) *MockLXDProfileUnitAssignedMachineIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitAssignedMachineIdCall) Do(f func() (string, error)) *MockLXDProfileUnitAssignedMachineIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitAssignedMachineIdCall) DoAndReturn(f func() (string, error)) *MockLXDProfileUnitAssignedMachineIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockLXDProfileUnit) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockLXDProfileUnitMockRecorder) Name() *MockLXDProfileUnitNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLXDProfileUnit)(nil).Name))
	return &MockLXDProfileUnitNameCall{Call: call}
}

// MockLXDProfileUnitNameCall wrap *gomock.Call
type MockLXDProfileUnitNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitNameCall) Return(arg0 string) *MockLXDProfileUnitNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitNameCall) Do(f func() string) *MockLXDProfileUnitNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitNameCall) DoAndReturn(f func() string) *MockLXDProfileUnitNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Tag mocks base method.
func (m *MockLXDProfileUnit) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockLXDProfileUnitMockRecorder) Tag() *MockLXDProfileUnitTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockLXDProfileUnit)(nil).Tag))
	return &MockLXDProfileUnitTagCall{Call: call}
}

// MockLXDProfileUnitTagCall wrap *gomock.Call
type MockLXDProfileUnitTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitTagCall) Return(arg0 names.Tag) *MockLXDProfileUnitTagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitTagCall) Do(f func() names.Tag) *MockLXDProfileUnitTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitTagCall) DoAndReturn(f func() names.Tag) *MockLXDProfileUnitTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchLXDProfileUpgradeNotifications mocks base method.
func (m *MockLXDProfileUnit) WatchLXDProfileUpgradeNotifications() (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchLXDProfileUpgradeNotifications")
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchLXDProfileUpgradeNotifications indicates an expected call of WatchLXDProfileUpgradeNotifications.
func (mr *MockLXDProfileUnitMockRecorder) WatchLXDProfileUpgradeNotifications() *MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLXDProfileUpgradeNotifications", reflect.TypeOf((*MockLXDProfileUnit)(nil).WatchLXDProfileUpgradeNotifications))
	return &MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall{Call: call}
}

// MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall wrap *gomock.Call
type MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall) Return(arg0 state.StringsWatcher, arg1 error) *MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall) Do(f func() (state.StringsWatcher, error)) *MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall) DoAndReturn(f func() (state.StringsWatcher, error)) *MockLXDProfileUnitWatchLXDProfileUpgradeNotificationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLXDProfileBackendV2 is a mock of LXDProfileBackendV2 interface.
type MockLXDProfileBackendV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileBackendV2MockRecorder
}

// MockLXDProfileBackendV2MockRecorder is the mock recorder for MockLXDProfileBackendV2.
type MockLXDProfileBackendV2MockRecorder struct {
	mock *MockLXDProfileBackendV2
}

// NewMockLXDProfileBackendV2 creates a new mock instance.
func NewMockLXDProfileBackendV2(ctrl *gomock.Controller) *MockLXDProfileBackendV2 {
	mock := &MockLXDProfileBackendV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileBackendV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileBackendV2) EXPECT() *MockLXDProfileBackendV2MockRecorder {
	return m.recorder
}

// Machine mocks base method.
func (m *MockLXDProfileBackendV2) Machine(arg0 string) (uniter.LXDProfileMachineV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileMachineV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockLXDProfileBackendV2MockRecorder) Machine(arg0 any) *MockLXDProfileBackendV2MachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockLXDProfileBackendV2)(nil).Machine), arg0)
	return &MockLXDProfileBackendV2MachineCall{Call: call}
}

// MockLXDProfileBackendV2MachineCall wrap *gomock.Call
type MockLXDProfileBackendV2MachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileBackendV2MachineCall) Return(arg0 uniter.LXDProfileMachineV2, arg1 error) *MockLXDProfileBackendV2MachineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileBackendV2MachineCall) Do(f func(string) (uniter.LXDProfileMachineV2, error)) *MockLXDProfileBackendV2MachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileBackendV2MachineCall) DoAndReturn(f func(string) (uniter.LXDProfileMachineV2, error)) *MockLXDProfileBackendV2MachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Unit mocks base method.
func (m *MockLXDProfileBackendV2) Unit(arg0 string) (uniter.LXDProfileUnitV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileUnitV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockLXDProfileBackendV2MockRecorder) Unit(arg0 any) *MockLXDProfileBackendV2UnitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockLXDProfileBackendV2)(nil).Unit), arg0)
	return &MockLXDProfileBackendV2UnitCall{Call: call}
}

// MockLXDProfileBackendV2UnitCall wrap *gomock.Call
type MockLXDProfileBackendV2UnitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileBackendV2UnitCall) Return(arg0 uniter.LXDProfileUnitV2, arg1 error) *MockLXDProfileBackendV2UnitCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileBackendV2UnitCall) Do(f func(string) (uniter.LXDProfileUnitV2, error)) *MockLXDProfileBackendV2UnitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileBackendV2UnitCall) DoAndReturn(f func(string) (uniter.LXDProfileUnitV2, error)) *MockLXDProfileBackendV2UnitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLXDProfileMachineV2 is a mock of LXDProfileMachineV2 interface.
type MockLXDProfileMachineV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileMachineV2MockRecorder
}

// MockLXDProfileMachineV2MockRecorder is the mock recorder for MockLXDProfileMachineV2.
type MockLXDProfileMachineV2MockRecorder struct {
	mock *MockLXDProfileMachineV2
}

// NewMockLXDProfileMachineV2 creates a new mock instance.
func NewMockLXDProfileMachineV2(ctrl *gomock.Controller) *MockLXDProfileMachineV2 {
	mock := &MockLXDProfileMachineV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileMachineV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileMachineV2) EXPECT() *MockLXDProfileMachineV2MockRecorder {
	return m.recorder
}

// ContainerType mocks base method.
func (m *MockLXDProfileMachineV2) ContainerType() instance.ContainerType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType.
func (mr *MockLXDProfileMachineV2MockRecorder) ContainerType() *MockLXDProfileMachineV2ContainerTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockLXDProfileMachineV2)(nil).ContainerType))
	return &MockLXDProfileMachineV2ContainerTypeCall{Call: call}
}

// MockLXDProfileMachineV2ContainerTypeCall wrap *gomock.Call
type MockLXDProfileMachineV2ContainerTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileMachineV2ContainerTypeCall) Return(arg0 instance.ContainerType) *MockLXDProfileMachineV2ContainerTypeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileMachineV2ContainerTypeCall) Do(f func() instance.ContainerType) *MockLXDProfileMachineV2ContainerTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileMachineV2ContainerTypeCall) DoAndReturn(f func() instance.ContainerType) *MockLXDProfileMachineV2ContainerTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsManual mocks base method.
func (m *MockLXDProfileMachineV2) IsManual() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsManual")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsManual indicates an expected call of IsManual.
func (mr *MockLXDProfileMachineV2MockRecorder) IsManual() *MockLXDProfileMachineV2IsManualCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManual", reflect.TypeOf((*MockLXDProfileMachineV2)(nil).IsManual))
	return &MockLXDProfileMachineV2IsManualCall{Call: call}
}

// MockLXDProfileMachineV2IsManualCall wrap *gomock.Call
type MockLXDProfileMachineV2IsManualCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileMachineV2IsManualCall) Return(arg0 bool, arg1 error) *MockLXDProfileMachineV2IsManualCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileMachineV2IsManualCall) Do(f func() (bool, error)) *MockLXDProfileMachineV2IsManualCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileMachineV2IsManualCall) DoAndReturn(f func() (bool, error)) *MockLXDProfileMachineV2IsManualCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLXDProfileUnitV2 is a mock of LXDProfileUnitV2 interface.
type MockLXDProfileUnitV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileUnitV2MockRecorder
}

// MockLXDProfileUnitV2MockRecorder is the mock recorder for MockLXDProfileUnitV2.
type MockLXDProfileUnitV2MockRecorder struct {
	mock *MockLXDProfileUnitV2
}

// NewMockLXDProfileUnitV2 creates a new mock instance.
func NewMockLXDProfileUnitV2(ctrl *gomock.Controller) *MockLXDProfileUnitV2 {
	mock := &MockLXDProfileUnitV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileUnitV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileUnitV2) EXPECT() *MockLXDProfileUnitV2MockRecorder {
	return m.recorder
}

// ApplicationName mocks base method.
func (m *MockLXDProfileUnitV2) ApplicationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ApplicationName indicates an expected call of ApplicationName.
func (mr *MockLXDProfileUnitV2MockRecorder) ApplicationName() *MockLXDProfileUnitV2ApplicationNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationName", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).ApplicationName))
	return &MockLXDProfileUnitV2ApplicationNameCall{Call: call}
}

// MockLXDProfileUnitV2ApplicationNameCall wrap *gomock.Call
type MockLXDProfileUnitV2ApplicationNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitV2ApplicationNameCall) Return(arg0 string) *MockLXDProfileUnitV2ApplicationNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitV2ApplicationNameCall) Do(f func() string) *MockLXDProfileUnitV2ApplicationNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitV2ApplicationNameCall) DoAndReturn(f func() string) *MockLXDProfileUnitV2ApplicationNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AssignedMachineId mocks base method.
func (m *MockLXDProfileUnitV2) AssignedMachineId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId.
func (mr *MockLXDProfileUnitV2MockRecorder) AssignedMachineId() *MockLXDProfileUnitV2AssignedMachineIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).AssignedMachineId))
	return &MockLXDProfileUnitV2AssignedMachineIdCall{Call: call}
}

// MockLXDProfileUnitV2AssignedMachineIdCall wrap *gomock.Call
type MockLXDProfileUnitV2AssignedMachineIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitV2AssignedMachineIdCall) Return(arg0 string, arg1 error) *MockLXDProfileUnitV2AssignedMachineIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitV2AssignedMachineIdCall) Do(f func() (string, error)) *MockLXDProfileUnitV2AssignedMachineIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitV2AssignedMachineIdCall) DoAndReturn(f func() (string, error)) *MockLXDProfileUnitV2AssignedMachineIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CharmURL mocks base method.
func (m *MockLXDProfileUnitV2) CharmURL() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmURL")
	ret0, _ := ret[0].(*string)
	return ret0
}

// CharmURL indicates an expected call of CharmURL.
func (mr *MockLXDProfileUnitV2MockRecorder) CharmURL() *MockLXDProfileUnitV2CharmURLCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmURL", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).CharmURL))
	return &MockLXDProfileUnitV2CharmURLCall{Call: call}
}

// MockLXDProfileUnitV2CharmURLCall wrap *gomock.Call
type MockLXDProfileUnitV2CharmURLCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitV2CharmURLCall) Return(arg0 *string) *MockLXDProfileUnitV2CharmURLCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitV2CharmURLCall) Do(f func() *string) *MockLXDProfileUnitV2CharmURLCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitV2CharmURLCall) DoAndReturn(f func() *string) *MockLXDProfileUnitV2CharmURLCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockLXDProfileUnitV2) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockLXDProfileUnitV2MockRecorder) Name() *MockLXDProfileUnitV2NameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).Name))
	return &MockLXDProfileUnitV2NameCall{Call: call}
}

// MockLXDProfileUnitV2NameCall wrap *gomock.Call
type MockLXDProfileUnitV2NameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitV2NameCall) Return(arg0 string) *MockLXDProfileUnitV2NameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitV2NameCall) Do(f func() string) *MockLXDProfileUnitV2NameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitV2NameCall) DoAndReturn(f func() string) *MockLXDProfileUnitV2NameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Tag mocks base method.
func (m *MockLXDProfileUnitV2) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockLXDProfileUnitV2MockRecorder) Tag() *MockLXDProfileUnitV2TagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).Tag))
	return &MockLXDProfileUnitV2TagCall{Call: call}
}

// MockLXDProfileUnitV2TagCall wrap *gomock.Call
type MockLXDProfileUnitV2TagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileUnitV2TagCall) Return(arg0 names.Tag) *MockLXDProfileUnitV2TagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileUnitV2TagCall) Do(f func() names.Tag) *MockLXDProfileUnitV2TagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileUnitV2TagCall) DoAndReturn(f func() names.Tag) *MockLXDProfileUnitV2TagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLXDProfileCharmV2 is a mock of LXDProfileCharmV2 interface.
type MockLXDProfileCharmV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileCharmV2MockRecorder
}

// MockLXDProfileCharmV2MockRecorder is the mock recorder for MockLXDProfileCharmV2.
type MockLXDProfileCharmV2MockRecorder struct {
	mock *MockLXDProfileCharmV2
}

// NewMockLXDProfileCharmV2 creates a new mock instance.
func NewMockLXDProfileCharmV2(ctrl *gomock.Controller) *MockLXDProfileCharmV2 {
	mock := &MockLXDProfileCharmV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileCharmV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileCharmV2) EXPECT() *MockLXDProfileCharmV2MockRecorder {
	return m.recorder
}

// LXDProfile mocks base method.
func (m *MockLXDProfileCharmV2) LXDProfile() lxdprofile.Profile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(lxdprofile.Profile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile.
func (mr *MockLXDProfileCharmV2MockRecorder) LXDProfile() *MockLXDProfileCharmV2LXDProfileCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockLXDProfileCharmV2)(nil).LXDProfile))
	return &MockLXDProfileCharmV2LXDProfileCall{Call: call}
}

// MockLXDProfileCharmV2LXDProfileCall wrap *gomock.Call
type MockLXDProfileCharmV2LXDProfileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileCharmV2LXDProfileCall) Return(arg0 lxdprofile.Profile) *MockLXDProfileCharmV2LXDProfileCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileCharmV2LXDProfileCall) Do(f func() lxdprofile.Profile) *MockLXDProfileCharmV2LXDProfileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileCharmV2LXDProfileCall) DoAndReturn(f func() lxdprofile.Profile) *MockLXDProfileCharmV2LXDProfileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
