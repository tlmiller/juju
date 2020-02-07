// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/storage/v1 (interfaces: StorageV1Interface,StorageClassInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/storage/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/kubernetes/typed/storage/v1"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockStorageV1Interface is a mock of StorageV1Interface interface
type MockStorageV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageV1InterfaceMockRecorder
}

// MockStorageV1InterfaceMockRecorder is the mock recorder for MockStorageV1Interface
type MockStorageV1InterfaceMockRecorder struct {
	mock *MockStorageV1Interface
}

// NewMockStorageV1Interface creates a new mock instance
func NewMockStorageV1Interface(ctrl *gomock.Controller) *MockStorageV1Interface {
	mock := &MockStorageV1Interface{ctrl: ctrl}
	mock.recorder = &MockStorageV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageV1Interface) EXPECT() *MockStorageV1InterfaceMockRecorder {
	return m.recorder
}

// RESTClient mocks base method
func (m *MockStorageV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockStorageV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockStorageV1Interface)(nil).RESTClient))
}

// StorageClasses mocks base method
func (m *MockStorageV1Interface) StorageClasses() v11.StorageClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageClasses")
	ret0, _ := ret[0].(v11.StorageClassInterface)
	return ret0
}

// StorageClasses indicates an expected call of StorageClasses
func (mr *MockStorageV1InterfaceMockRecorder) StorageClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageClasses", reflect.TypeOf((*MockStorageV1Interface)(nil).StorageClasses))
}

// MockStorageClassInterface is a mock of StorageClassInterface interface
type MockStorageClassInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClassInterfaceMockRecorder
}

// MockStorageClassInterfaceMockRecorder is the mock recorder for MockStorageClassInterface
type MockStorageClassInterfaceMockRecorder struct {
	mock *MockStorageClassInterface
}

// NewMockStorageClassInterface creates a new mock instance
func NewMockStorageClassInterface(ctrl *gomock.Controller) *MockStorageClassInterface {
	mock := &MockStorageClassInterface{ctrl: ctrl}
	mock.recorder = &MockStorageClassInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageClassInterface) EXPECT() *MockStorageClassInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockStorageClassInterface) Create(arg0 *v1.StorageClass) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockStorageClassInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStorageClassInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockStorageClassInterface) Delete(arg0 string, arg1 *v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStorageClassInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorageClassInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockStorageClassInterface) DeleteCollection(arg0 *v10.DeleteOptions, arg1 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockStorageClassInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockStorageClassInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockStorageClassInterface) Get(arg0 string, arg1 v10.GetOptions) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStorageClassInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorageClassInterface)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockStorageClassInterface) List(arg0 v10.ListOptions) (*v1.StorageClassList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.StorageClassList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStorageClassInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageClassInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockStorageClassInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockStorageClassInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockStorageClassInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockStorageClassInterface) Update(arg0 *v1.StorageClass) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockStorageClassInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStorageClassInterface)(nil).Update), arg0)
}

// Watch mocks base method
func (m *MockStorageClassInterface) Watch(arg0 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockStorageClassInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockStorageClassInterface)(nil).Watch), arg0)
}
