// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/samaritan-proxy/samaritan/proc/redis (interfaces: Filter)

// Package redis is a generated GoMock package.
package redis

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFilter is a mock of Filter interface
type MockFilter struct {
	ctrl     *gomock.Controller
	recorder *MockFilterMockRecorder
}

// MockFilterMockRecorder is the mock recorder for MockFilter
type MockFilterMockRecorder struct {
	mock *MockFilter
}

// NewMockFilter creates a new mock instance
func NewMockFilter(ctrl *gomock.Controller) *MockFilter {
	mock := &MockFilter{ctrl: ctrl}
	mock.recorder = &MockFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilter) EXPECT() *MockFilterMockRecorder {
	return m.recorder
}

// Destroy mocks base method
func (m *MockFilter) Destroy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destroy")
}

// Destroy indicates an expected call of Destroy
func (mr *MockFilterMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockFilter)(nil).Destroy))
}

// Do mocks base method
func (m *MockFilter) Do(arg0 *simpleRequest) FilterStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(FilterStatus)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockFilterMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockFilter)(nil).Do), arg0)
}