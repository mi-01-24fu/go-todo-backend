// Code generated by MockGen. DO NOT EDIT.
// Source: internal\infrastructure\get_list\access_list.go

// Package get_list is a generated GoMock package.
package getlist

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccessTODO is a mock of AccessTODO interface.
type MockAccessTODO struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTODOMockRecorder
}

// MockAccessTODOMockRecorder is the mock recorder for MockAccessTODO.
type MockAccessTODOMockRecorder struct {
	mock *MockAccessTODO
}

// NewMockAccessTODO creates a new mock instance.
func NewMockAccessTODO(ctrl *gomock.Controller) *MockAccessTODO {
	mock := &MockAccessTODO{ctrl: ctrl}
	mock.recorder = &MockAccessTODOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessTODO) EXPECT() *MockAccessTODOMockRecorder {
	return m.recorder
}

// CheckID mocks base method.
func (m *MockAccessTODO) CheckID(arg0 GetTODORequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckID", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckID indicates an expected call of CheckID.
func (mr *MockAccessTODOMockRecorder) CheckID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckID", reflect.TypeOf((*MockAccessTODO)(nil).CheckID), arg0)
}

// GetTODOList mocks base method.
func (m *MockAccessTODO) GetTODOList(arg0 GetTODORequest) (GetLists, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTODOList", arg0)
	ret0, _ := ret[0].(GetLists)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTODOList indicates an expected call of GetTODOList.
func (mr *MockAccessTODOMockRecorder) GetTODOList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTODOList", reflect.TypeOf((*MockAccessTODO)(nil).GetTODOList), arg0)
}
