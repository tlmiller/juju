// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/port/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination package_mock_test.go github.com/juju/juju/domain/port/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	set "github.com/juju/collections/set"
	application "github.com/juju/juju/core/application"
	machine "github.com/juju/juju/core/machine"
	network "github.com/juju/juju/core/network"
	unit "github.com/juju/juju/core/unit"
	port "github.com/juju/juju/domain/port"
	gomock "go.uber.org/mock/gomock"
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

// FilterUnitUUIDsForApplication mocks base method.
func (m *MockState) FilterUnitUUIDsForApplication(arg0 context.Context, arg1 []unit.UUID, arg2 application.ID) (set.Strings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterUnitUUIDsForApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].(set.Strings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterUnitUUIDsForApplication indicates an expected call of FilterUnitUUIDsForApplication.
func (mr *MockStateMockRecorder) FilterUnitUUIDsForApplication(arg0, arg1, arg2 any) *MockStateFilterUnitUUIDsForApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterUnitUUIDsForApplication", reflect.TypeOf((*MockState)(nil).FilterUnitUUIDsForApplication), arg0, arg1, arg2)
	return &MockStateFilterUnitUUIDsForApplicationCall{Call: call}
}

// MockStateFilterUnitUUIDsForApplicationCall wrap *gomock.Call
type MockStateFilterUnitUUIDsForApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateFilterUnitUUIDsForApplicationCall) Return(arg0 set.Strings, arg1 error) *MockStateFilterUnitUUIDsForApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateFilterUnitUUIDsForApplicationCall) Do(f func(context.Context, []unit.UUID, application.ID) (set.Strings, error)) *MockStateFilterUnitUUIDsForApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateFilterUnitUUIDsForApplicationCall) DoAndReturn(f func(context.Context, []unit.UUID, application.ID) (set.Strings, error)) *MockStateFilterUnitUUIDsForApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAllOpenedPorts mocks base method.
func (m *MockState) GetAllOpenedPorts(arg0 context.Context) (port.UnitGroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOpenedPorts", arg0)
	ret0, _ := ret[0].(port.UnitGroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOpenedPorts indicates an expected call of GetAllOpenedPorts.
func (mr *MockStateMockRecorder) GetAllOpenedPorts(arg0 any) *MockStateGetAllOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOpenedPorts", reflect.TypeOf((*MockState)(nil).GetAllOpenedPorts), arg0)
	return &MockStateGetAllOpenedPortsCall{Call: call}
}

// MockStateGetAllOpenedPortsCall wrap *gomock.Call
type MockStateGetAllOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetAllOpenedPortsCall) Return(arg0 port.UnitGroupedPortRanges, arg1 error) *MockStateGetAllOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetAllOpenedPortsCall) Do(f func(context.Context) (port.UnitGroupedPortRanges, error)) *MockStateGetAllOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetAllOpenedPortsCall) DoAndReturn(f func(context.Context) (port.UnitGroupedPortRanges, error)) *MockStateGetAllOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetApplicationOpenedPorts mocks base method.
func (m *MockState) GetApplicationOpenedPorts(arg0 context.Context, arg1 application.ID) (port.UnitEndpointPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationOpenedPorts", arg0, arg1)
	ret0, _ := ret[0].(port.UnitEndpointPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationOpenedPorts indicates an expected call of GetApplicationOpenedPorts.
func (mr *MockStateMockRecorder) GetApplicationOpenedPorts(arg0, arg1 any) *MockStateGetApplicationOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationOpenedPorts", reflect.TypeOf((*MockState)(nil).GetApplicationOpenedPorts), arg0, arg1)
	return &MockStateGetApplicationOpenedPortsCall{Call: call}
}

// MockStateGetApplicationOpenedPortsCall wrap *gomock.Call
type MockStateGetApplicationOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetApplicationOpenedPortsCall) Return(arg0 port.UnitEndpointPortRanges, arg1 error) *MockStateGetApplicationOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetApplicationOpenedPortsCall) Do(f func(context.Context, application.ID) (port.UnitEndpointPortRanges, error)) *MockStateGetApplicationOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetApplicationOpenedPortsCall) DoAndReturn(f func(context.Context, application.ID) (port.UnitEndpointPortRanges, error)) *MockStateGetApplicationOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMachineNamesForUnits mocks base method.
func (m *MockState) GetMachineNamesForUnits(arg0 context.Context, arg1 []unit.UUID) ([]machine.Name, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineNamesForUnits", arg0, arg1)
	ret0, _ := ret[0].([]machine.Name)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineNamesForUnits indicates an expected call of GetMachineNamesForUnits.
func (mr *MockStateMockRecorder) GetMachineNamesForUnits(arg0, arg1 any) *MockStateGetMachineNamesForUnitsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineNamesForUnits", reflect.TypeOf((*MockState)(nil).GetMachineNamesForUnits), arg0, arg1)
	return &MockStateGetMachineNamesForUnitsCall{Call: call}
}

// MockStateGetMachineNamesForUnitsCall wrap *gomock.Call
type MockStateGetMachineNamesForUnitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetMachineNamesForUnitsCall) Return(arg0 []machine.Name, arg1 error) *MockStateGetMachineNamesForUnitsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetMachineNamesForUnitsCall) Do(f func(context.Context, []unit.UUID) ([]machine.Name, error)) *MockStateGetMachineNamesForUnitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetMachineNamesForUnitsCall) DoAndReturn(f func(context.Context, []unit.UUID) ([]machine.Name, error)) *MockStateGetMachineNamesForUnitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMachineOpenedPorts mocks base method.
func (m *MockState) GetMachineOpenedPorts(arg0 context.Context, arg1 string) (map[unit.Name]network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineOpenedPorts", arg0, arg1)
	ret0, _ := ret[0].(map[unit.Name]network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineOpenedPorts indicates an expected call of GetMachineOpenedPorts.
func (mr *MockStateMockRecorder) GetMachineOpenedPorts(arg0, arg1 any) *MockStateGetMachineOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineOpenedPorts", reflect.TypeOf((*MockState)(nil).GetMachineOpenedPorts), arg0, arg1)
	return &MockStateGetMachineOpenedPortsCall{Call: call}
}

// MockStateGetMachineOpenedPortsCall wrap *gomock.Call
type MockStateGetMachineOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetMachineOpenedPortsCall) Return(arg0 map[unit.Name]network.GroupedPortRanges, arg1 error) *MockStateGetMachineOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetMachineOpenedPortsCall) Do(f func(context.Context, string) (map[unit.Name]network.GroupedPortRanges, error)) *MockStateGetMachineOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetMachineOpenedPortsCall) DoAndReturn(f func(context.Context, string) (map[unit.Name]network.GroupedPortRanges, error)) *MockStateGetMachineOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitOpenedPorts mocks base method.
func (m *MockState) GetUnitOpenedPorts(arg0 context.Context, arg1 unit.UUID) (network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitOpenedPorts", arg0, arg1)
	ret0, _ := ret[0].(network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitOpenedPorts indicates an expected call of GetUnitOpenedPorts.
func (mr *MockStateMockRecorder) GetUnitOpenedPorts(arg0, arg1 any) *MockStateGetUnitOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitOpenedPorts", reflect.TypeOf((*MockState)(nil).GetUnitOpenedPorts), arg0, arg1)
	return &MockStateGetUnitOpenedPortsCall{Call: call}
}

// MockStateGetUnitOpenedPortsCall wrap *gomock.Call
type MockStateGetUnitOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetUnitOpenedPortsCall) Return(arg0 network.GroupedPortRanges, arg1 error) *MockStateGetUnitOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetUnitOpenedPortsCall) Do(f func(context.Context, unit.UUID) (network.GroupedPortRanges, error)) *MockStateGetUnitOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetUnitOpenedPortsCall) DoAndReturn(f func(context.Context, unit.UUID) (network.GroupedPortRanges, error)) *MockStateGetUnitOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitUUID mocks base method.
func (m *MockState) GetUnitUUID(arg0 context.Context, arg1 unit.Name) (unit.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitUUID", arg0, arg1)
	ret0, _ := ret[0].(unit.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitUUID indicates an expected call of GetUnitUUID.
func (mr *MockStateMockRecorder) GetUnitUUID(arg0, arg1 any) *MockStateGetUnitUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitUUID", reflect.TypeOf((*MockState)(nil).GetUnitUUID), arg0, arg1)
	return &MockStateGetUnitUUIDCall{Call: call}
}

// MockStateGetUnitUUIDCall wrap *gomock.Call
type MockStateGetUnitUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetUnitUUIDCall) Return(arg0 unit.UUID, arg1 error) *MockStateGetUnitUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetUnitUUIDCall) Do(f func(context.Context, unit.Name) (unit.UUID, error)) *MockStateGetUnitUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetUnitUUIDCall) DoAndReturn(f func(context.Context, unit.Name) (unit.UUID, error)) *MockStateGetUnitUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InitialWatchMachineOpenedPortsStatement mocks base method.
func (m *MockState) InitialWatchMachineOpenedPortsStatement() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitialWatchMachineOpenedPortsStatement")
	ret0, _ := ret[0].(string)
	return ret0
}

// InitialWatchMachineOpenedPortsStatement indicates an expected call of InitialWatchMachineOpenedPortsStatement.
func (mr *MockStateMockRecorder) InitialWatchMachineOpenedPortsStatement() *MockStateInitialWatchMachineOpenedPortsStatementCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitialWatchMachineOpenedPortsStatement", reflect.TypeOf((*MockState)(nil).InitialWatchMachineOpenedPortsStatement))
	return &MockStateInitialWatchMachineOpenedPortsStatementCall{Call: call}
}

// MockStateInitialWatchMachineOpenedPortsStatementCall wrap *gomock.Call
type MockStateInitialWatchMachineOpenedPortsStatementCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateInitialWatchMachineOpenedPortsStatementCall) Return(arg0 string) *MockStateInitialWatchMachineOpenedPortsStatementCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateInitialWatchMachineOpenedPortsStatementCall) Do(f func() string) *MockStateInitialWatchMachineOpenedPortsStatementCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateInitialWatchMachineOpenedPortsStatementCall) DoAndReturn(f func() string) *MockStateInitialWatchMachineOpenedPortsStatementCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUnitPorts mocks base method.
func (m *MockState) UpdateUnitPorts(arg0 context.Context, arg1 unit.UUID, arg2, arg3 network.GroupedPortRanges) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnitPorts", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUnitPorts indicates an expected call of UpdateUnitPorts.
func (mr *MockStateMockRecorder) UpdateUnitPorts(arg0, arg1, arg2, arg3 any) *MockStateUpdateUnitPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnitPorts", reflect.TypeOf((*MockState)(nil).UpdateUnitPorts), arg0, arg1, arg2, arg3)
	return &MockStateUpdateUnitPortsCall{Call: call}
}

// MockStateUpdateUnitPortsCall wrap *gomock.Call
type MockStateUpdateUnitPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateUpdateUnitPortsCall) Return(arg0 error) *MockStateUpdateUnitPortsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateUpdateUnitPortsCall) Do(f func(context.Context, unit.UUID, network.GroupedPortRanges, network.GroupedPortRanges) error) *MockStateUpdateUnitPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateUpdateUnitPortsCall) DoAndReturn(f func(context.Context, unit.UUID, network.GroupedPortRanges, network.GroupedPortRanges) error) *MockStateUpdateUnitPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchOpenedPortsTable mocks base method.
func (m *MockState) WatchOpenedPortsTable() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOpenedPortsTable")
	ret0, _ := ret[0].(string)
	return ret0
}

// WatchOpenedPortsTable indicates an expected call of WatchOpenedPortsTable.
func (mr *MockStateMockRecorder) WatchOpenedPortsTable() *MockStateWatchOpenedPortsTableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOpenedPortsTable", reflect.TypeOf((*MockState)(nil).WatchOpenedPortsTable))
	return &MockStateWatchOpenedPortsTableCall{Call: call}
}

// MockStateWatchOpenedPortsTableCall wrap *gomock.Call
type MockStateWatchOpenedPortsTableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateWatchOpenedPortsTableCall) Return(arg0 string) *MockStateWatchOpenedPortsTableCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateWatchOpenedPortsTableCall) Do(f func() string) *MockStateWatchOpenedPortsTableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateWatchOpenedPortsTableCall) DoAndReturn(f func() string) *MockStateWatchOpenedPortsTableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
