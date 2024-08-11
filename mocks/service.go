// Package mock_services is a generated GoMock package.
package mock_services

import (
	entity "bank-service/internal/entity"
	services "bank-service/internal/services"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockReposI is a mock of ReposI interface.
type MockReposI struct {
	ctrl     *gomock.Controller
	recorder *MockReposIMockRecorder
}

// MockReposIMockRecorder is the mock recorder for MockReposI.
type MockReposIMockRecorder struct {
	mock *MockReposI
}

// NewMockReposI creates a new mock instance.
func NewMockReposI(ctrl *gomock.Controller) *MockReposI {
	mock := &MockReposI{ctrl: ctrl}
	mock.recorder = &MockReposIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReposI) EXPECT() *MockReposIMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockReposI) CreateAccount(user *services.CreateAccount) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockReposIMockRecorder) CreateAccount(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockReposI)(nil).CreateAccount), user)
}

// GetBalance mocks base method.
func (m *MockReposI) GetBalance(user *services.GetBalance) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockReposIMockRecorder) GetBalance(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockReposI)(nil).GetBalance), user)
}

// UpdateBalance mocks base method.
func (m *MockReposI) UpdateBalance(user *services.UpdateBalance) (*entity.User, error) {
	ret := m.ctrl.Call(m, "UpdateBalance", user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockReposIMockRecorder) UpdateBalance(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockReposI)(nil).UpdateBalance), user)
}
