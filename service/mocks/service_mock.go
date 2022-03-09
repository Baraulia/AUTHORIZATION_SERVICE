// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	model "stlab.itechart-group.com/go/food_delivery/authorization_service/model"
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

// CheckRights mocks base method.
func (m *MockAuthorization) CheckRights(token, requiredRole string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRights", token, requiredRole)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRights indicates an expected call of CheckRights.
func (mr *MockAuthorizationMockRecorder) CheckRights(token, requiredRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRights", reflect.TypeOf((*MockAuthorization)(nil).CheckRights), token, requiredRole)
}

// GenerateTokensByAuthUser mocks base method.
func (m *MockAuthorization) GenerateTokensByAuthUser(user *authProto.User) (*authProto.GeneratedTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokensByAuthUser", user)
	ret0, _ := ret[0].(*authProto.GeneratedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTokensByAuthUser indicates an expected call of GenerateTokensByAuthUser.
func (mr *MockAuthorizationMockRecorder) GenerateTokensByAuthUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokensByAuthUser", reflect.TypeOf((*MockAuthorization)(nil).GenerateTokensByAuthUser), user)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (*authProto.UserRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(*authProto.UserRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// RefreshTokens mocks base method.
func (m *MockAuthorization) RefreshTokens(refreshToken string) (*authProto.GeneratedTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", refreshToken)
	ret0, _ := ret[0].(*authProto.GeneratedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens.
func (mr *MockAuthorizationMockRecorder) RefreshTokens(refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockAuthorization)(nil).RefreshTokens), refreshToken)
}

// MockRolePerm is a mock of RolePerm interface.
type MockRolePerm struct {
	ctrl     *gomock.Controller
	recorder *MockRolePermMockRecorder
}

// MockRolePermMockRecorder is the mock recorder for MockRolePerm.
type MockRolePermMockRecorder struct {
	mock *MockRolePerm
}

// NewMockRolePerm creates a new mock instance.
func NewMockRolePerm(ctrl *gomock.Controller) *MockRolePerm {
	mock := &MockRolePerm{ctrl: ctrl}
	mock.recorder = &MockRolePermMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRolePerm) EXPECT() *MockRolePermMockRecorder {
	return m.recorder
}

// AddRoleToUser mocks base method.
func (m *MockRolePerm) AddRoleToUser(user *authProto.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoleToUser", user)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRoleToUser indicates an expected call of AddRoleToUser.
func (mr *MockRolePermMockRecorder) AddRoleToUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoleToUser", reflect.TypeOf((*MockRolePerm)(nil).AddRoleToUser), user)
}

// BindRoleWithPerms mocks base method.
func (m *MockRolePerm) BindRoleWithPerms(rp *model.BindRoleWithPermission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindRoleWithPerms", rp)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindRoleWithPerms indicates an expected call of BindRoleWithPerms.
func (mr *MockRolePermMockRecorder) BindRoleWithPerms(rp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindRoleWithPerms", reflect.TypeOf((*MockRolePerm)(nil).BindRoleWithPerms), rp)
}

// CreatePermission mocks base method.
func (m *MockRolePerm) CreatePermission(permission string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermission", permission)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermission indicates an expected call of CreatePermission.
func (mr *MockRolePermMockRecorder) CreatePermission(permission interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermission", reflect.TypeOf((*MockRolePerm)(nil).CreatePermission), permission)
}

// CreateRole mocks base method.
func (m *MockRolePerm) CreateRole(role string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", role)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRolePermMockRecorder) CreateRole(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRolePerm)(nil).CreateRole), role)
}

// GetAllPerms mocks base method.
func (m *MockRolePerm) GetAllPerms() ([]model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPerms")
	ret0, _ := ret[0].([]model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPerms indicates an expected call of GetAllPerms.
func (mr *MockRolePermMockRecorder) GetAllPerms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPerms", reflect.TypeOf((*MockRolePerm)(nil).GetAllPerms))
}

// GetAllRoles mocks base method.
func (m *MockRolePerm) GetAllRoles() ([]model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRoles")
	ret0, _ := ret[0].([]model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRoles indicates an expected call of GetAllRoles.
func (mr *MockRolePermMockRecorder) GetAllRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoles", reflect.TypeOf((*MockRolePerm)(nil).GetAllRoles))
}

// GetPermsByRoleId mocks base method.
func (m *MockRolePerm) GetPermsByRoleId(id int) ([]model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermsByRoleId", id)
	ret0, _ := ret[0].([]model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermsByRoleId indicates an expected call of GetPermsByRoleId.
func (mr *MockRolePermMockRecorder) GetPermsByRoleId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermsByRoleId", reflect.TypeOf((*MockRolePerm)(nil).GetPermsByRoleId), id)
}

// GetRoleById mocks base method.
func (m *MockRolePerm) GetRoleById(id int) (*model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleById", id)
	ret0, _ := ret[0].(*model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleById indicates an expected call of GetRoleById.
func (mr *MockRolePermMockRecorder) GetRoleById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleById", reflect.TypeOf((*MockRolePerm)(nil).GetRoleById), id)
}

// GetRoleByUserId mocks base method.
func (m *MockRolePerm) GetRoleByUserId(userId int) (*model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleByUserId", userId)
	ret0, _ := ret[0].(*model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleByUserId indicates an expected call of GetRoleByUserId.
func (mr *MockRolePermMockRecorder) GetRoleByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleByUserId", reflect.TypeOf((*MockRolePerm)(nil).GetRoleByUserId), userId)
}
