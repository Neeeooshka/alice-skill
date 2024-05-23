// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/links.go

// Package storage is a generated GoMock package.
package storage

import (
	http "net/http"
	reflect "reflect"

	storage "github.com/Neeeooshka/alice-skill.git/internal/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockLinkStorage is a mock of LinkStorage interface.
type MockLinkStorage struct {
	ctrl     *gomock.Controller
	recorder *MockLinkStorageMockRecorder
}

// MockLinkStorageMockRecorder is the mock recorder for MockLinkStorage.
type MockLinkStorageMockRecorder struct {
	mock *MockLinkStorage
}

// NewMockLinkStorage creates a new mock instance.
func NewMockLinkStorage(ctrl *gomock.Controller) *MockLinkStorage {
	mock := &MockLinkStorage{ctrl: ctrl}
	mock.recorder = &MockLinkStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinkStorage) EXPECT() *MockLinkStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockLinkStorage) Add(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockLinkStorageMockRecorder) Add(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockLinkStorage)(nil).Add), arg0, arg1, arg2)
}

// AddBatch mocks base method.
func (m *MockLinkStorage) AddBatch(arg0 []storage.Batch, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBatch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBatch indicates an expected call of AddBatch.
func (mr *MockLinkStorageMockRecorder) AddBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBatch", reflect.TypeOf((*MockLinkStorage)(nil).AddBatch), arg0, arg1)
}

// Close mocks base method.
func (m *MockLinkStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockLinkStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLinkStorage)(nil).Close))
}

// Get mocks base method.
func (m *MockLinkStorage) Get(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLinkStorageMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLinkStorage)(nil).Get), arg0)
}

// GetUserURLs mocks base method.
func (m *MockLinkStorage) GetUserURLs(arg0 string) []storage.Link {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserURLs", arg0)
	ret0, _ := ret[0].([]storage.Link)
	return ret0
}

// GetUserURLs indicates an expected call of GetUserURLs.
func (mr *MockLinkStorageMockRecorder) GetUserURLs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserURLs", reflect.TypeOf((*MockLinkStorage)(nil).GetUserURLs), arg0)
}

// PingHandler mocks base method.
func (m *MockLinkStorage) PingHandler(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PingHandler", arg0, arg1)
}

// PingHandler indicates an expected call of PingHandler.
func (mr *MockLinkStorageMockRecorder) PingHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingHandler", reflect.TypeOf((*MockLinkStorage)(nil).PingHandler), arg0, arg1)
}
