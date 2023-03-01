// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mehmetcekirdekci/GolangCRUD/app/product/services (interfaces: BaseService)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
)

// MockBaseService is a mock of BaseService interface.
type MockBaseService struct {
	ctrl     *gomock.Controller
	recorder *MockBaseServiceMockRecorder
}

// MockBaseServiceMockRecorder is the mock recorder for MockBaseService.
type MockBaseServiceMockRecorder struct {
	mock *MockBaseService
}

// NewMockBaseService creates a new mock instance.
func NewMockBaseService(ctrl *gomock.Controller) *MockBaseService {
	mock := &MockBaseService{ctrl: ctrl}
	mock.recorder = &MockBaseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseService) EXPECT() *MockBaseServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBaseService) Create(arg0 types.CreateProductDto) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBaseServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBaseService)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockBaseService) Delete(arg0 int) (*types.Product, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*types.Product)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Delete indicates an expected call of Delete.
func (mr *MockBaseServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBaseService)(nil).Delete), arg0)
}

// GetById mocks base method.
func (m *MockBaseService) GetById(arg0 int) (*types.Product, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*types.Product)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetById indicates an expected call of GetById.
func (mr *MockBaseServiceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBaseService)(nil).GetById), arg0)
}

// Update mocks base method.
func (m *MockBaseService) Update(arg0 types.UpdateProductDto) (*types.Product, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*types.Product)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockBaseServiceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBaseService)(nil).Update), arg0)
}