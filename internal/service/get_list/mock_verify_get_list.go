// Code generated by MockGen. DO NOT EDIT.
// Source: internal\service\get_list\verify_get_list.go

// Package get_list is a generated GoMock package.
package get_list

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	get_list "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/get_list"
)

// MockVerifyGetTODOList is a mock of VerifyGetTODOList interface.
type MockVerifyGetTODOList struct {
	ctrl     *gomock.Controller
	recorder *MockVerifyGetTODOListMockRecorder
}

// MockVerifyGetTODOListMockRecorder is the mock recorder for MockVerifyGetTODOList.
type MockVerifyGetTODOListMockRecorder struct {
	mock *MockVerifyGetTODOList
}

// NewMockVerifyGetTODOList creates a new mock instance.
func NewMockVerifyGetTODOList(ctrl *gomock.Controller) *MockVerifyGetTODOList {
	mock := &MockVerifyGetTODOList{ctrl: ctrl}
	mock.recorder = &MockVerifyGetTODOListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifyGetTODOList) EXPECT() *MockVerifyGetTODOListMockRecorder {
	return m.recorder
}

// GetTODOList mocks base method.
func (m *MockVerifyGetTODOList) GetTODOList(arg0 get_list.GetTODORequest) (ResponseList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTODOList", arg0)
	ret0, _ := ret[0].(ResponseList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTODOList indicates an expected call of GetTODOList.
func (mr *MockVerifyGetTODOListMockRecorder) GetTODOList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTODOList", reflect.TypeOf((*MockVerifyGetTODOList)(nil).GetTODOList), arg0)
}
