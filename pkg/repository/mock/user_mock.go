// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/user.go

// Package mock is a generated GoMock package.
package mock

import (
	models "eatry/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(userID int, address models.AddAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", userID, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(userID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), userID, address)
}

// CheckUserAvailability mocks base method.
func (m *MockUserRepository) CheckUserAvailability(email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserAvailability", email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUserAvailability indicates an expected call of CheckUserAvailability.
func (mr *MockUserRepositoryMockRecorder) CheckUserAvailability(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserAvailability", reflect.TypeOf((*MockUserRepository)(nil).CheckUserAvailability), email)
}

// DeleteAddress mocks base method.
func (m *MockUserRepository) DeleteAddress(userID, addressID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", userID, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserRepositoryMockRecorder) DeleteAddress(userID, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserRepository)(nil).DeleteAddress), userID, addressID)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", user)
	ret0, _ := ret[0].(models.UserSignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), user)
}

// GetAddresses mocks base method.
func (m *MockUserRepository) GetAddresses(id int) ([]models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses", id)
	ret0, _ := ret[0].([]models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockUserRepositoryMockRecorder) GetAddresses(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockUserRepository)(nil).GetAddresses), id)
}

// GetAllPaymentOptions mocks base method.
func (m *MockUserRepository) GetAllPaymentOptions() ([]models.PaymentDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPaymentOptions")
	ret0, _ := ret[0].([]models.PaymentDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPaymentOptions indicates an expected call of GetAllPaymentOptions.
func (mr *MockUserRepositoryMockRecorder) GetAllPaymentOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPaymentOptions", reflect.TypeOf((*MockUserRepository)(nil).GetAllPaymentOptions))
}

// LoginHandler mocks base method.
func (m *MockUserRepository) LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginHandler", user)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginHandler indicates an expected call of LoginHandler.
func (mr *MockUserRepositoryMockRecorder) LoginHandler(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginHandler", reflect.TypeOf((*MockUserRepository)(nil).LoginHandler), user)
}

// UserProfile mocks base method.
func (m *MockUserRepository) UserProfile(id int) (models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserProfile", id)
	ret0, _ := ret[0].(models.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserProfile indicates an expected call of UserProfile.
func (mr *MockUserRepositoryMockRecorder) UserProfile(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserProfile", reflect.TypeOf((*MockUserRepository)(nil).UserProfile), id)
}

// UserSignUp mocks base method.
func (m *MockUserRepository) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", user)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserRepositoryMockRecorder) UserSignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserRepository)(nil).UserSignUp), user)
}
