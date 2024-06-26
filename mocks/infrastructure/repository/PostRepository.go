// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/zakiyalmaya/assetfindr-assignment/model"
	gorm "gorm.io/gorm"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// Assosiate mocks base method.
func (m *MockPostRepository) Assosiate(post *model.Post, tags []*model.Tag, tx ...*gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{post, tags}
	for _, a := range tx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Assosiate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Assosiate indicates an expected call of Assosiate.
func (mr *MockPostRepositoryMockRecorder) Assosiate(post, tags interface{}, tx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{post, tags}, tx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assosiate", reflect.TypeOf((*MockPostRepository)(nil).Assosiate), varargs...)
}

// Create mocks base method.
func (m *MockPostRepository) Create(post *model.Post, tx ...*gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{post}
	for _, a := range tx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPostRepositoryMockRecorder) Create(post interface{}, tx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{post}, tx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostRepository)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockPostRepository) Delete(id int, tx ...*gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range tx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostRepositoryMockRecorder) Delete(id interface{}, tx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, tx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostRepository)(nil).Delete), varargs...)
}

// GetAll mocks base method.
func (m *MockPostRepository) GetAll() ([]*model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPostRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPostRepository)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockPostRepository) GetByID(id int, tx ...*gorm.DB) (*model.Post, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range tx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByID", varargs...)
	ret0, _ := ret[0].(*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockPostRepositoryMockRecorder) GetByID(id interface{}, tx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, tx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPostRepository)(nil).GetByID), varargs...)
}

// Update mocks base method.
func (m *MockPostRepository) Update(post *model.Post, tx ...*gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{post}
	for _, a := range tx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostRepositoryMockRecorder) Update(post interface{}, tx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{post}, tx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostRepository)(nil).Update), varargs...)
}
