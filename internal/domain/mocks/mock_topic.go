// Code generated by MockGen. DO NOT EDIT.
// Source: topic.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

// MockTopicRepo is a mock of TopicRepo interface.
type MockTopicRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTopicRepoMockRecorder
}

// MockTopicRepoMockRecorder is the mock recorder for MockTopicRepo.
type MockTopicRepoMockRecorder struct {
	mock *MockTopicRepo
}

// NewMockTopicRepo creates a new mock instance.
func NewMockTopicRepo(ctrl *gomock.Controller) *MockTopicRepo {
	mock := &MockTopicRepo{ctrl: ctrl}
	mock.recorder = &MockTopicRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopicRepo) EXPECT() *MockTopicRepoMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockTopicRepo) DeleteByID(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockTopicRepoMockRecorder) DeleteByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockTopicRepo)(nil).DeleteByID), ctx, id)
}

// GetAll mocks base method.
func (m *MockTopicRepo) GetAll(ctx context.Context) ([]domain.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]domain.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTopicRepoMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTopicRepo)(nil).GetAll), ctx)
}

// GetByID mocks base method.
func (m *MockTopicRepo) GetByID(id int32) (domain.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(domain.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTopicRepoMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTopicRepo)(nil).GetByID), id)
}

// Store mocks base method.
func (m *MockTopicRepo) Store(ctx context.Context, topic domain.Topic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTopicRepoMockRecorder) Store(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTopicRepo)(nil).Store), ctx, topic)
}
