// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	thumbnails "github.com/romik1505/youtubeThumbnails/pkg/api/thumbnails"
)

// MockIThumbnailService is a mock of IThumbnailService interface.
type MockIThumbnailService struct {
	ctrl     *gomock.Controller
	recorder *MockIThumbnailServiceMockRecorder
}

// MockIThumbnailServiceMockRecorder is the mock recorder for MockIThumbnailService.
type MockIThumbnailServiceMockRecorder struct {
	mock *MockIThumbnailService
}

// NewMockIThumbnailService creates a new mock instance.
func NewMockIThumbnailService(ctrl *gomock.Controller) *MockIThumbnailService {
	mock := &MockIThumbnailService{ctrl: ctrl}
	mock.recorder = &MockIThumbnailServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIThumbnailService) EXPECT() *MockIThumbnailServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIThumbnailService) Get(arg0 context.Context, arg1 *thumbnails.GetRequest) (*thumbnails.GetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*thumbnails.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIThumbnailServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIThumbnailService)(nil).Get), arg0, arg1)
}

// MockIFileLoader is a mock of IFileLoader interface.
type MockIFileLoader struct {
	ctrl     *gomock.Controller
	recorder *MockIFileLoaderMockRecorder
}

// MockIFileLoaderMockRecorder is the mock recorder for MockIFileLoader.
type MockIFileLoaderMockRecorder struct {
	mock *MockIFileLoader
}

// NewMockIFileLoader creates a new mock instance.
func NewMockIFileLoader(ctrl *gomock.Controller) *MockIFileLoader {
	mock := &MockIFileLoader{ctrl: ctrl}
	mock.recorder = &MockIFileLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFileLoader) EXPECT() *MockIFileLoaderMockRecorder {
	return m.recorder
}

// LoadImg mocks base method.
func (m *MockIFileLoader) LoadImg(ctx context.Context, url string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadImg", ctx, url)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadImg indicates an expected call of LoadImg.
func (mr *MockIFileLoaderMockRecorder) LoadImg(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImg", reflect.TypeOf((*MockIFileLoader)(nil).LoadImg), ctx, url)
}