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

// MockTagRepository is a mock of TagRepository interface.
type MockTagRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTagRepositoryMockRecorder
}

// MockTagRepositoryMockRecorder is the mock recorder for MockTagRepository.
type MockTagRepositoryMockRecorder struct {
	mock *MockTagRepository
}

// NewMockTagRepository creates a new mock instance.
func NewMockTagRepository(ctrl *gomock.Controller) *MockTagRepository {
	mock := &MockTagRepository{ctrl: ctrl}
	mock.recorder = &MockTagRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagRepository) EXPECT() *MockTagRepositoryMockRecorder {
	return m.recorder
}

// GetOrCreate mocks base method.
func (m *MockTagRepository) GetOrCreate(tag *model.Tag, tx *gorm.DB) (*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreate", tag, tx)
	ret0, _ := ret[0].(*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreate indicates an expected call of GetOrCreate.
func (mr *MockTagRepositoryMockRecorder) GetOrCreate(tag, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreate", reflect.TypeOf((*MockTagRepository)(nil).GetOrCreate), tag, tx)
}