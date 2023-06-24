// Code generated by MockGen. DO NOT EDIT.
// Source: internal/infrastructure/login/login_info.go

// Package mock_access_login is a generated GoMock package.
package login

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccessLoginInfo is a mock of AccessLoginInfo interface.
type MockAccessLoginInfo struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLoginInfoMockRecorder
}

// MockAccessLoginInfoMockRecorder is the mock recorder for MockAccessLoginInfo.
type MockAccessLoginInfoMockRecorder struct {
	mock *MockAccessLoginInfo
}

// NewMockAccessLoginInfo creates a new mock instance.
func NewMockAccessLoginInfo(ctrl *gomock.Controller) *MockAccessLoginInfo {
	mock := &MockAccessLoginInfo{ctrl: ctrl}
	mock.recorder = &MockAccessLoginInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLoginInfo) EXPECT() *MockAccessLoginInfoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAccessLoginInfo) Get(arg0 context.Context) (UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAccessLoginInfoMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccessLoginInfo)(nil).Get), arg0)
}
