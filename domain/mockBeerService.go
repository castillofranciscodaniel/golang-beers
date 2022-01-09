// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/castillofranciscodaniel/golang-beers/domain (interfaces: BeerService)

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBeerService is a mock of BeerService interface.
type MockBeerService struct {
	ctrl     *gomock.Controller
	recorder *MockBeerServiceMockRecorder
}

// MockBeerServiceMockRecorder is the mock recorder for MockBeerService.
type MockBeerServiceMockRecorder struct {
	mock *MockBeerService
}

// NewMockBeerService creates a new mock instance.
func NewMockBeerService(ctrl *gomock.Controller) *MockBeerService {
	mock := &MockBeerService{ctrl: ctrl}
	mock.recorder = &MockBeerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeerService) EXPECT() *MockBeerServiceMockRecorder {
	return m.recorder
}

// BoxPrice mocks base method.
func (m *MockBeerService) BoxPrice(arg0 int64, arg1 string, arg2 int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BoxPrice", arg0, arg1, arg2)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BoxPrice indicates an expected call of BoxPrice.
func (mr *MockBeerServiceMockRecorder) BoxPrice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BoxPrice", reflect.TypeOf((*MockBeerService)(nil).BoxPrice), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockBeerService) Get() ([]Beer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]Beer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBeerServiceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBeerService)(nil).Get))
}

// GetById mocks base method.
func (m *MockBeerService) GetById(arg0 int64) (*Beer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*Beer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBeerServiceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBeerService)(nil).GetById), arg0)
}

// Post mocks base method.
func (m *MockBeerService) Post(arg0 Beer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockBeerServiceMockRecorder) Post(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockBeerService)(nil).Post), arg0)
}