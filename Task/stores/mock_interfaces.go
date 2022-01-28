// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package store is a generated GoMock package.
package store

import (
	reflect "reflect"
	models "zopsmart/Task/models"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStore) Create(Id int, Name, Email, Phone string, Age int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", Id, Name, Email, Phone, Age)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStoreMockRecorder) Create(Id, Name, Email, Phone, Age interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), Id, Name, Email, Phone, Age)
}

// Delete mocks base method.
func (m *MockStore) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), id)
}

// GetAllUsersStore mocks base method.
func (m *MockStore) GetAllUsersStore() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsersStore")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsersStore indicates an expected call of GetAllUsersStore.
func (mr *MockStoreMockRecorder) GetAllUsersStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsersStore", reflect.TypeOf((*MockStore)(nil).GetAllUsersStore))
}

// GetMail mocks base method.
func (m *MockStore) GetMail(mail string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMail", mail)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMail indicates an expected call of GetMail.
func (mr *MockStoreMockRecorder) GetMail(mail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMail", reflect.TypeOf((*MockStore)(nil).GetMail), mail)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), id)
}

// Update mocks base method.
func (m *MockStore) Update(Id int, Phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", Id, Phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStoreMockRecorder) Update(Id, Phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), Id, Phone)
}
