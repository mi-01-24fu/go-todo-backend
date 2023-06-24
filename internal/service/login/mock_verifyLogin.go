// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/verifyLogin.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	login "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
)

// MockLoginRepository is a mock of LoginRepository interface.
type MockLoginRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLoginRepositoryMockRecorder
}

// MockLoginRepositoryMockRecorder is the mock recorder for MockLoginRepository.
type MockLoginRepositoryMockRecorder struct {
	mock *MockLoginRepository
}

// NewMockLoginRepository creates a new mock instance.
func NewMockLoginRepository(ctrl *gomock.Controller) *MockLoginRepository {
	mock := &MockLoginRepository{ctrl: ctrl}
	mock.recorder = &MockLoginRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginRepository) EXPECT() *MockLoginRepositoryMockRecorder {
	return m.recorder
}

// VerifyLogin mocks base method.
func (m *MockLoginRepository) VerifyLogin(arg0 context.Context, arg1 login.UserInfo) (login.VerifyLoginResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyLogin", arg0, arg1)
	ret0, _ := ret[0].(login.VerifyLoginResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyLogin indicates an expected call of VerifyLogin.
func (mr *MockLoginRepositoryMockRecorder) VerifyLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyLogin", reflect.TypeOf((*MockLoginRepository)(nil).VerifyLogin), arg0, arg1)
}