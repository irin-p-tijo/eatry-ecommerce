// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interfaces/user.go

// Package mock is a generated GoMock package.
package mock

import (
	models "eatry/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserUseCase) AddAddress(userID int, address models.AddAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", userID, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserUseCaseMockRecorder) AddAddress(userID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserUseCase)(nil).AddAddress), userID, address)
}

// CheckOut mocks base method.
func (m *MockUserUseCase) CheckOut(userID int) (models.CheckoutDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOut", userID)
	ret0, _ := ret[0].(models.CheckoutDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOut indicates an expected call of CheckOut.
func (mr *MockUserUseCaseMockRecorder) CheckOut(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOut", reflect.TypeOf((*MockUserUseCase)(nil).CheckOut), userID)
}

// DeleteAddress mocks base method.
func (m *MockUserUseCase) DeleteAddress(userID, addressID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", userID, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserUseCaseMockRecorder) DeleteAddress(userID, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserUseCase)(nil).DeleteAddress), userID, addressID)
}

// GetAllAddress mocks base method.
func (m *MockUserUseCase) GetAllAddress(id int) ([]models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAddress", id)
	ret0, _ := ret[0].([]models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAddress indicates an expected call of GetAllAddress.
func (mr *MockUserUseCaseMockRecorder) GetAllAddress(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAddress", reflect.TypeOf((*MockUserUseCase)(nil).GetAllAddress), id)
}

// LoginHandler mocks base method.
func (m *MockUserUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginHandler", user)
	ret0, _ := ret[0].(models.TokenUsers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginHandler indicates an expected call of LoginHandler.
func (mr *MockUserUseCaseMockRecorder) LoginHandler(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginHandler", reflect.TypeOf((*MockUserUseCase)(nil).LoginHandler), user)
}

// UserProfile mocks base method.
func (m *MockUserUseCase) UserProfile(id int) (models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserProfile", id)
	ret0, _ := ret[0].(models.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserProfile indicates an expected call of UserProfile.
func (mr *MockUserUseCaseMockRecorder) UserProfile(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserProfile", reflect.TypeOf((*MockUserUseCase)(nil).UserProfile), id)
}

// UserSignUp mocks base method.
func (m *MockUserUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", user)
	ret0, _ := ret[0].(models.TokenUsers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserUseCaseMockRecorder) UserSignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserUseCase)(nil).UserSignUp), user)
}
