// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/resource (interfaces: ResourceService,ApplicationService)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/service_mock.go github.com/juju/juju/internal/resource ResourceService,ApplicationService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	application "github.com/juju/juju/core/application"
	resource "github.com/juju/juju/core/resource"
	unit "github.com/juju/juju/core/unit"
	charm "github.com/juju/juju/domain/application/charm"
	resource0 "github.com/juju/juju/domain/resource"
	charm0 "github.com/juju/juju/internal/charm"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceService is a mock of ResourceService interface.
type MockResourceService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceServiceMockRecorder
}

// MockResourceServiceMockRecorder is the mock recorder for MockResourceService.
type MockResourceServiceMockRecorder struct {
	mock *MockResourceService
}

// NewMockResourceService creates a new mock instance.
func NewMockResourceService(ctrl *gomock.Controller) *MockResourceService {
	mock := &MockResourceService{ctrl: ctrl}
	mock.recorder = &MockResourceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceService) EXPECT() *MockResourceServiceMockRecorder {
	return m.recorder
}

// GetResource mocks base method.
func (m *MockResourceService) GetResource(arg0 context.Context, arg1 resource.UUID) (resource0.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0, arg1)
	ret0, _ := ret[0].(resource0.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockResourceServiceMockRecorder) GetResource(arg0, arg1 any) *MockResourceServiceGetResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockResourceService)(nil).GetResource), arg0, arg1)
	return &MockResourceServiceGetResourceCall{Call: call}
}

// MockResourceServiceGetResourceCall wrap *gomock.Call
type MockResourceServiceGetResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceServiceGetResourceCall) Return(arg0 resource0.Resource, arg1 error) *MockResourceServiceGetResourceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceServiceGetResourceCall) Do(f func(context.Context, resource.UUID) (resource0.Resource, error)) *MockResourceServiceGetResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceServiceGetResourceCall) DoAndReturn(f func(context.Context, resource.UUID) (resource0.Resource, error)) *MockResourceServiceGetResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetResourceUUID mocks base method.
func (m *MockResourceService) GetResourceUUID(arg0 context.Context, arg1 resource0.GetResourceUUIDArgs) (resource.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceUUID", arg0, arg1)
	ret0, _ := ret[0].(resource.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceUUID indicates an expected call of GetResourceUUID.
func (mr *MockResourceServiceMockRecorder) GetResourceUUID(arg0, arg1 any) *MockResourceServiceGetResourceUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceUUID", reflect.TypeOf((*MockResourceService)(nil).GetResourceUUID), arg0, arg1)
	return &MockResourceServiceGetResourceUUIDCall{Call: call}
}

// MockResourceServiceGetResourceUUIDCall wrap *gomock.Call
type MockResourceServiceGetResourceUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceServiceGetResourceUUIDCall) Return(arg0 resource.UUID, arg1 error) *MockResourceServiceGetResourceUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceServiceGetResourceUUIDCall) Do(f func(context.Context, resource0.GetResourceUUIDArgs) (resource.UUID, error)) *MockResourceServiceGetResourceUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceServiceGetResourceUUIDCall) DoAndReturn(f func(context.Context, resource0.GetResourceUUIDArgs) (resource.UUID, error)) *MockResourceServiceGetResourceUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenResource mocks base method.
func (m *MockResourceService) OpenResource(arg0 context.Context, arg1 resource.UUID) (resource0.Resource, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenResource", arg0, arg1)
	ret0, _ := ret[0].(resource0.Resource)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OpenResource indicates an expected call of OpenResource.
func (mr *MockResourceServiceMockRecorder) OpenResource(arg0, arg1 any) *MockResourceServiceOpenResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenResource", reflect.TypeOf((*MockResourceService)(nil).OpenResource), arg0, arg1)
	return &MockResourceServiceOpenResourceCall{Call: call}
}

// MockResourceServiceOpenResourceCall wrap *gomock.Call
type MockResourceServiceOpenResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceServiceOpenResourceCall) Return(arg0 resource0.Resource, arg1 io.ReadCloser, arg2 error) *MockResourceServiceOpenResourceCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceServiceOpenResourceCall) Do(f func(context.Context, resource.UUID) (resource0.Resource, io.ReadCloser, error)) *MockResourceServiceOpenResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceServiceOpenResourceCall) DoAndReturn(f func(context.Context, resource.UUID) (resource0.Resource, io.ReadCloser, error)) *MockResourceServiceOpenResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetApplicationResource mocks base method.
func (m *MockResourceService) SetApplicationResource(arg0 context.Context, arg1 resource.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetApplicationResource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetApplicationResource indicates an expected call of SetApplicationResource.
func (mr *MockResourceServiceMockRecorder) SetApplicationResource(arg0, arg1 any) *MockResourceServiceSetApplicationResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApplicationResource", reflect.TypeOf((*MockResourceService)(nil).SetApplicationResource), arg0, arg1)
	return &MockResourceServiceSetApplicationResourceCall{Call: call}
}

// MockResourceServiceSetApplicationResourceCall wrap *gomock.Call
type MockResourceServiceSetApplicationResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceServiceSetApplicationResourceCall) Return(arg0 error) *MockResourceServiceSetApplicationResourceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceServiceSetApplicationResourceCall) Do(f func(context.Context, resource.UUID) error) *MockResourceServiceSetApplicationResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceServiceSetApplicationResourceCall) DoAndReturn(f func(context.Context, resource.UUID) error) *MockResourceServiceSetApplicationResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetUnitResource mocks base method.
func (m *MockResourceService) SetUnitResource(arg0 context.Context, arg1 resource.UUID, arg2 unit.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnitResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitResource indicates an expected call of SetUnitResource.
func (mr *MockResourceServiceMockRecorder) SetUnitResource(arg0, arg1, arg2 any) *MockResourceServiceSetUnitResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitResource", reflect.TypeOf((*MockResourceService)(nil).SetUnitResource), arg0, arg1, arg2)
	return &MockResourceServiceSetUnitResourceCall{Call: call}
}

// MockResourceServiceSetUnitResourceCall wrap *gomock.Call
type MockResourceServiceSetUnitResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceServiceSetUnitResourceCall) Return(arg0 error) *MockResourceServiceSetUnitResourceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceServiceSetUnitResourceCall) Do(f func(context.Context, resource.UUID, unit.UUID) error) *MockResourceServiceSetUnitResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceServiceSetUnitResourceCall) DoAndReturn(f func(context.Context, resource.UUID, unit.UUID) error) *MockResourceServiceSetUnitResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StoreResource mocks base method.
func (m *MockResourceService) StoreResource(arg0 context.Context, arg1 resource0.StoreResourceArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreResource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreResource indicates an expected call of StoreResource.
func (mr *MockResourceServiceMockRecorder) StoreResource(arg0, arg1 any) *MockResourceServiceStoreResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreResource", reflect.TypeOf((*MockResourceService)(nil).StoreResource), arg0, arg1)
	return &MockResourceServiceStoreResourceCall{Call: call}
}

// MockResourceServiceStoreResourceCall wrap *gomock.Call
type MockResourceServiceStoreResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceServiceStoreResourceCall) Return(arg0 error) *MockResourceServiceStoreResourceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceServiceStoreResourceCall) Do(f func(context.Context, resource0.StoreResourceArgs) error) *MockResourceServiceStoreResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceServiceStoreResourceCall) DoAndReturn(f func(context.Context, resource0.StoreResourceArgs) error) *MockResourceServiceStoreResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// GetApplicationIDByName mocks base method.
func (m *MockApplicationService) GetApplicationIDByName(arg0 context.Context, arg1 string) (application.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationIDByName", arg0, arg1)
	ret0, _ := ret[0].(application.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationIDByName indicates an expected call of GetApplicationIDByName.
func (mr *MockApplicationServiceMockRecorder) GetApplicationIDByName(arg0, arg1 any) *MockApplicationServiceGetApplicationIDByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationIDByName", reflect.TypeOf((*MockApplicationService)(nil).GetApplicationIDByName), arg0, arg1)
	return &MockApplicationServiceGetApplicationIDByNameCall{Call: call}
}

// MockApplicationServiceGetApplicationIDByNameCall wrap *gomock.Call
type MockApplicationServiceGetApplicationIDByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetApplicationIDByNameCall) Return(arg0 application.ID, arg1 error) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetApplicationIDByNameCall) Do(f func(context.Context, string) (application.ID, error)) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetApplicationIDByNameCall) DoAndReturn(f func(context.Context, string) (application.ID, error)) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetApplicationIDByUnitName mocks base method.
func (m *MockApplicationService) GetApplicationIDByUnitName(arg0 context.Context, arg1 unit.Name) (application.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationIDByUnitName", arg0, arg1)
	ret0, _ := ret[0].(application.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationIDByUnitName indicates an expected call of GetApplicationIDByUnitName.
func (mr *MockApplicationServiceMockRecorder) GetApplicationIDByUnitName(arg0, arg1 any) *MockApplicationServiceGetApplicationIDByUnitNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationIDByUnitName", reflect.TypeOf((*MockApplicationService)(nil).GetApplicationIDByUnitName), arg0, arg1)
	return &MockApplicationServiceGetApplicationIDByUnitNameCall{Call: call}
}

// MockApplicationServiceGetApplicationIDByUnitNameCall wrap *gomock.Call
type MockApplicationServiceGetApplicationIDByUnitNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetApplicationIDByUnitNameCall) Return(arg0 application.ID, arg1 error) *MockApplicationServiceGetApplicationIDByUnitNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetApplicationIDByUnitNameCall) Do(f func(context.Context, unit.Name) (application.ID, error)) *MockApplicationServiceGetApplicationIDByUnitNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetApplicationIDByUnitNameCall) DoAndReturn(f func(context.Context, unit.Name) (application.ID, error)) *MockApplicationServiceGetApplicationIDByUnitNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmByApplicationID mocks base method.
func (m *MockApplicationService) GetCharmByApplicationID(arg0 context.Context, arg1 application.ID) (charm0.Charm, charm.CharmLocator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmByApplicationID", arg0, arg1)
	ret0, _ := ret[0].(charm0.Charm)
	ret1, _ := ret[1].(charm.CharmLocator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCharmByApplicationID indicates an expected call of GetCharmByApplicationID.
func (mr *MockApplicationServiceMockRecorder) GetCharmByApplicationID(arg0, arg1 any) *MockApplicationServiceGetCharmByApplicationIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmByApplicationID", reflect.TypeOf((*MockApplicationService)(nil).GetCharmByApplicationID), arg0, arg1)
	return &MockApplicationServiceGetCharmByApplicationIDCall{Call: call}
}

// MockApplicationServiceGetCharmByApplicationIDCall wrap *gomock.Call
type MockApplicationServiceGetCharmByApplicationIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetCharmByApplicationIDCall) Return(arg0 charm0.Charm, arg1 charm.CharmLocator, arg2 error) *MockApplicationServiceGetCharmByApplicationIDCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetCharmByApplicationIDCall) Do(f func(context.Context, application.ID) (charm0.Charm, charm.CharmLocator, error)) *MockApplicationServiceGetCharmByApplicationIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetCharmByApplicationIDCall) DoAndReturn(f func(context.Context, application.ID) (charm0.Charm, charm.CharmLocator, error)) *MockApplicationServiceGetCharmByApplicationIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitUUID mocks base method.
func (m *MockApplicationService) GetUnitUUID(arg0 context.Context, arg1 unit.Name) (unit.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitUUID", arg0, arg1)
	ret0, _ := ret[0].(unit.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitUUID indicates an expected call of GetUnitUUID.
func (mr *MockApplicationServiceMockRecorder) GetUnitUUID(arg0, arg1 any) *MockApplicationServiceGetUnitUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitUUID", reflect.TypeOf((*MockApplicationService)(nil).GetUnitUUID), arg0, arg1)
	return &MockApplicationServiceGetUnitUUIDCall{Call: call}
}

// MockApplicationServiceGetUnitUUIDCall wrap *gomock.Call
type MockApplicationServiceGetUnitUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetUnitUUIDCall) Return(arg0 unit.UUID, arg1 error) *MockApplicationServiceGetUnitUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetUnitUUIDCall) Do(f func(context.Context, unit.Name) (unit.UUID, error)) *MockApplicationServiceGetUnitUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetUnitUUIDCall) DoAndReturn(f func(context.Context, unit.Name) (unit.UUID, error)) *MockApplicationServiceGetUnitUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
