// Code generated by MockGen. DO NOT EDIT.
// Source: thumbnail.go

// Package mock_thumbnail is a generated GoMock package.
package mock_thumbnail

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/romik1505/youtubeThumbnails/internal/app/model"
)

// MockIThumbnailRepository is a mock of IThumbnailRepository interface.
type MockIThumbnailRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIThumbnailRepositoryMockRecorder
}

// MockIThumbnailRepositoryMockRecorder is the mock recorder for MockIThumbnailRepository.
type MockIThumbnailRepositoryMockRecorder struct {
	mock *MockIThumbnailRepository
}

// NewMockIThumbnailRepository creates a new mock instance.
func NewMockIThumbnailRepository(ctrl *gomock.Controller) *MockIThumbnailRepository {
	mock := &MockIThumbnailRepository{ctrl: ctrl}
	mock.recorder = &MockIThumbnailRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIThumbnailRepository) EXPECT() *MockIThumbnailRepositoryMockRecorder {
	return m.recorder
}

// GetThumbnail mocks base method.
func (m *MockIThumbnailRepository) GetThumbnail(ctx context.Context, id string) (model.Thumbnail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThumbnail", ctx, id)
	ret0, _ := ret[0].(model.Thumbnail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThumbnail indicates an expected call of GetThumbnail.
func (mr *MockIThumbnailRepositoryMockRecorder) GetThumbnail(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThumbnail", reflect.TypeOf((*MockIThumbnailRepository)(nil).GetThumbnail), ctx, id)
}

// InsertThumbnail mocks base method.
func (m *MockIThumbnailRepository) InsertThumbnail(ctx context.Context, thumbnail model.Thumbnail) (model.Thumbnail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertThumbnail", ctx, thumbnail)
	ret0, _ := ret[0].(model.Thumbnail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertThumbnail indicates an expected call of InsertThumbnail.
func (mr *MockIThumbnailRepositoryMockRecorder) InsertThumbnail(ctx, thumbnail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertThumbnail", reflect.TypeOf((*MockIThumbnailRepository)(nil).InsertThumbnail), ctx, thumbnail)
}