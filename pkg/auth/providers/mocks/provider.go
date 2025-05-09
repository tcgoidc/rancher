// Code generated by MockGen. DO NOT EDIT.
// Source: ../common/provider.go
//
// Generated by this command:
//
//	mockgen -source=../common/provider.go -destination=./provider.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/rancher/norman/types"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	accessor "github.com/rancher/rancher/pkg/auth/accessor"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthProvider is a mock of AuthProvider interface.
type MockAuthProvider struct {
	ctrl     *gomock.Controller
	recorder *MockAuthProviderMockRecorder
	isgomock struct{}
}

// MockAuthProviderMockRecorder is the mock recorder for MockAuthProvider.
type MockAuthProviderMockRecorder struct {
	mock *MockAuthProvider
}

// NewMockAuthProvider creates a new mock instance.
func NewMockAuthProvider(ctrl *gomock.Controller) *MockAuthProvider {
	mock := &MockAuthProvider{ctrl: ctrl}
	mock.recorder = &MockAuthProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthProvider) EXPECT() *MockAuthProviderMockRecorder {
	return m.recorder
}

// AuthenticateUser mocks base method.
func (m *MockAuthProvider) AuthenticateUser(ctx context.Context, input any) (v3.Principal, []v3.Principal, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", ctx, input)
	ret0, _ := ret[0].(v3.Principal)
	ret1, _ := ret[1].([]v3.Principal)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockAuthProviderMockRecorder) AuthenticateUser(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockAuthProvider)(nil).AuthenticateUser), ctx, input)
}

// CanAccessWithGroupProviders mocks base method.
func (m *MockAuthProvider) CanAccessWithGroupProviders(userPrincipalID string, groups []v3.Principal) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanAccessWithGroupProviders", userPrincipalID, groups)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanAccessWithGroupProviders indicates an expected call of CanAccessWithGroupProviders.
func (mr *MockAuthProviderMockRecorder) CanAccessWithGroupProviders(userPrincipalID, groups any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanAccessWithGroupProviders", reflect.TypeOf((*MockAuthProvider)(nil).CanAccessWithGroupProviders), userPrincipalID, groups)
}

// CustomizeSchema mocks base method.
func (m *MockAuthProvider) CustomizeSchema(schema *types.Schema) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CustomizeSchema", schema)
}

// CustomizeSchema indicates an expected call of CustomizeSchema.
func (mr *MockAuthProviderMockRecorder) CustomizeSchema(schema any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomizeSchema", reflect.TypeOf((*MockAuthProvider)(nil).CustomizeSchema), schema)
}

// GetName mocks base method.
func (m *MockAuthProvider) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockAuthProviderMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockAuthProvider)(nil).GetName))
}

// GetPrincipal mocks base method.
func (m *MockAuthProvider) GetPrincipal(principalID string, token accessor.TokenAccessor) (v3.Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrincipal", principalID, token)
	ret0, _ := ret[0].(v3.Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrincipal indicates an expected call of GetPrincipal.
func (mr *MockAuthProviderMockRecorder) GetPrincipal(principalID, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrincipal", reflect.TypeOf((*MockAuthProvider)(nil).GetPrincipal), principalID, token)
}

// GetUserExtraAttributes mocks base method.
func (m *MockAuthProvider) GetUserExtraAttributes(userPrincipal v3.Principal) map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserExtraAttributes", userPrincipal)
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// GetUserExtraAttributes indicates an expected call of GetUserExtraAttributes.
func (mr *MockAuthProviderMockRecorder) GetUserExtraAttributes(userPrincipal any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserExtraAttributes", reflect.TypeOf((*MockAuthProvider)(nil).GetUserExtraAttributes), userPrincipal)
}

// IsDisabledProvider mocks base method.
func (m *MockAuthProvider) IsDisabledProvider() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDisabledProvider")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDisabledProvider indicates an expected call of IsDisabledProvider.
func (mr *MockAuthProviderMockRecorder) IsDisabledProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDisabledProvider", reflect.TypeOf((*MockAuthProvider)(nil).IsDisabledProvider))
}

// Logout mocks base method.
func (m *MockAuthProvider) Logout(apiContext *types.APIContext, token accessor.TokenAccessor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", apiContext, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthProviderMockRecorder) Logout(apiContext, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthProvider)(nil).Logout), apiContext, token)
}

// LogoutAll mocks base method.
func (m *MockAuthProvider) LogoutAll(apiContext *types.APIContext, token accessor.TokenAccessor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogoutAll", apiContext, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogoutAll indicates an expected call of LogoutAll.
func (mr *MockAuthProviderMockRecorder) LogoutAll(apiContext, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogoutAll", reflect.TypeOf((*MockAuthProvider)(nil).LogoutAll), apiContext, token)
}

// RefetchGroupPrincipals mocks base method.
func (m *MockAuthProvider) RefetchGroupPrincipals(principalID, secret string) ([]v3.Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefetchGroupPrincipals", principalID, secret)
	ret0, _ := ret[0].([]v3.Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefetchGroupPrincipals indicates an expected call of RefetchGroupPrincipals.
func (mr *MockAuthProviderMockRecorder) RefetchGroupPrincipals(principalID, secret any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefetchGroupPrincipals", reflect.TypeOf((*MockAuthProvider)(nil).RefetchGroupPrincipals), principalID, secret)
}

// SearchPrincipals mocks base method.
func (m *MockAuthProvider) SearchPrincipals(name, principalType string, myToken accessor.TokenAccessor) ([]v3.Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPrincipals", name, principalType, myToken)
	ret0, _ := ret[0].([]v3.Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPrincipals indicates an expected call of SearchPrincipals.
func (mr *MockAuthProviderMockRecorder) SearchPrincipals(name, principalType, myToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPrincipals", reflect.TypeOf((*MockAuthProvider)(nil).SearchPrincipals), name, principalType, myToken)
}

// TransformToAuthProvider mocks base method.
func (m *MockAuthProvider) TransformToAuthProvider(authConfig map[string]any) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransformToAuthProvider", authConfig)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransformToAuthProvider indicates an expected call of TransformToAuthProvider.
func (mr *MockAuthProviderMockRecorder) TransformToAuthProvider(authConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransformToAuthProvider", reflect.TypeOf((*MockAuthProvider)(nil).TransformToAuthProvider), authConfig)
}
