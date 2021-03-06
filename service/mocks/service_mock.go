// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	model "github.com/Baraulia/AUTHORIZATION_SERVICE/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), email, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockRoleList is a mock of RoleList interface.
type MockRoleList struct {
	ctrl     *gomock.Controller
	recorder *MockRoleListMockRecorder
}

// MockRoleListMockRecorder is the mock recorder for MockRoleList.
type MockRoleListMockRecorder struct {
	mock *MockRoleList
}

// NewMockRoleList creates a new mock instance.
func NewMockRoleList(ctrl *gomock.Controller) *MockRoleList {
	mock := &MockRoleList{ctrl: ctrl}
	mock.recorder = &MockRoleListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleList) EXPECT() *MockRoleListMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockRoleList) GetById(id int) (*model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRoleListMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRoleList)(nil).GetById), id)
}

// SelectPermission mocks base method.
func (m *MockRoleList) SelectPermission(id int) []model.Permission {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectPermission", id)
	ret0, _ := ret[0].([]model.Permission)
	return ret0
}

// SelectPermission indicates an expected call of SelectPermission.
func (mr *MockRoleListMockRecorder) SelectPermission(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectPermission", reflect.TypeOf((*MockRoleList)(nil).SelectPermission), id)
}
