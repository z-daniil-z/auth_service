// Code generated by MockGen. DO NOT EDIT.
// Source: auth/internal/entities (interfaces: ISessions)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockISessions is a mock of ISessions interface.
type MockISessions struct {
	ctrl     *gomock.Controller
	recorder *MockISessionsMockRecorder
}

// MockISessionsMockRecorder is the mock recorder for MockISessions.
type MockISessionsMockRecorder struct {
	mock *MockISessions
}

// NewMockISessions creates a new mock instance.
func NewMockISessions(ctrl *gomock.Controller) *MockISessions {
	mock := &MockISessions{ctrl: ctrl}
	mock.recorder = &MockISessionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISessions) EXPECT() *MockISessionsMockRecorder {
	return m.recorder
}

// GetUserIDByToken mocks base method.
func (m *MockISessions) GetUserIDByToken(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDByToken", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDByToken indicates an expected call of GetUserIDByToken.
func (mr *MockISessionsMockRecorder) GetUserIDByToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDByToken", reflect.TypeOf((*MockISessions)(nil).GetUserIDByToken), arg0)
}

// Insert mocks base method.
func (m *MockISessions) Insert(arg0 string, arg1 int64, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockISessionsMockRecorder) Insert(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockISessions)(nil).Insert), arg0, arg1, arg2)
}
