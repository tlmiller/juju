// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/charm/downloader (interfaces: RepositoryGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/juju/core/charm"
	downloader "github.com/juju/juju/core/charm/downloader"
)

// MockRepositoryGetter is a mock of RepositoryGetter interface.
type MockRepositoryGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryGetterMockRecorder
}

// MockRepositoryGetterMockRecorder is the mock recorder for MockRepositoryGetter.
type MockRepositoryGetterMockRecorder struct {
	mock *MockRepositoryGetter
}

// NewMockRepositoryGetter creates a new mock instance.
func NewMockRepositoryGetter(ctrl *gomock.Controller) *MockRepositoryGetter {
	mock := &MockRepositoryGetter{ctrl: ctrl}
	mock.recorder = &MockRepositoryGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryGetter) EXPECT() *MockRepositoryGetterMockRecorder {
	return m.recorder
}

// GetCharmRepository mocks base method.
func (m *MockRepositoryGetter) GetCharmRepository(arg0 charm.Source) (downloader.CharmRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmRepository", arg0)
	ret0, _ := ret[0].(downloader.CharmRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmRepository indicates an expected call of GetCharmRepository.
func (mr *MockRepositoryGetterMockRecorder) GetCharmRepository(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmRepository", reflect.TypeOf((*MockRepositoryGetter)(nil).GetCharmRepository), arg0)
}
