// Code generated by MockGen. DO NOT EDIT.
// Source: src/application/interfaces/iproduct_service.go

// Package mock_interfaces is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/xStrato/ports-adapters-go-sample/src/application/interfaces"
)

// MockIProductService is a mock of IProductService interface.
type MockIProductService struct {
	ctrl     *gomock.Controller
	recorder *MockIProductServiceMockRecorder
}

// MockIProductServiceMockRecorder is the mock recorder for MockIProductService.
type MockIProductServiceMockRecorder struct {
	mock *MockIProductService
}

// NewMockIProductService creates a new mock instance.
func NewMockIProductService(ctrl *gomock.Controller) *MockIProductService {
	mock := &MockIProductService{ctrl: ctrl}
	mock.recorder = &MockIProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductService) EXPECT() *MockIProductServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIProductService) Create(name string, price float32) (interfaces.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, price)
	ret0, _ := ret[0].(interfaces.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIProductServiceMockRecorder) Create(name, price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIProductService)(nil).Create), name, price)
}

// Disable mocks base method.
func (m *MockIProductService) Disable(product interfaces.IProduct) (interfaces.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", product)
	ret0, _ := ret[0].(interfaces.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disable indicates an expected call of Disable.
func (mr *MockIProductServiceMockRecorder) Disable(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockIProductService)(nil).Disable), product)
}

// Enable mocks base method.
func (m *MockIProductService) Enable(product interfaces.IProduct) (interfaces.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", product)
	ret0, _ := ret[0].(interfaces.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enable indicates an expected call of Enable.
func (mr *MockIProductServiceMockRecorder) Enable(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockIProductService)(nil).Enable), product)
}

// Get mocks base method.
func (m *MockIProductService) Get(id string) (interfaces.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(interfaces.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIProductServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIProductService)(nil).Get), id)
}
