// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/long12310225/pitayaN/v2/networkentity (interfaces: NetworkEntity)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	protos "github.com/long12310225/pitayaN/v2/protos"
	net "net"
	reflect "reflect"
)

// MockNetworkEntity is a mock of NetworkEntity interface
type MockNetworkEntity struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkEntityMockRecorder
}

// MockNetworkEntityMockRecorder is the mock recorder for MockNetworkEntity
type MockNetworkEntityMockRecorder struct {
	mock *MockNetworkEntity
}

// NewMockNetworkEntity creates a new mock instance
func NewMockNetworkEntity(ctrl *gomock.Controller) *MockNetworkEntity {
	mock := &MockNetworkEntity{ctrl: ctrl}
	mock.recorder = &MockNetworkEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkEntity) EXPECT() *MockNetworkEntityMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockNetworkEntity) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockNetworkEntityMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNetworkEntity)(nil).Close))
}

// Kick mocks base method
func (m *MockNetworkEntity) Kick(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kick", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Kick indicates an expected call of Kick
func (mr *MockNetworkEntityMockRecorder) Kick(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kick", reflect.TypeOf((*MockNetworkEntity)(nil).Kick), arg0)
}

// Push mocks base method
func (m *MockNetworkEntity) Push(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push
func (mr *MockNetworkEntityMockRecorder) Push(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockNetworkEntity)(nil).Push), arg0, arg1)
}

// RemoteAddr mocks base method
func (m *MockNetworkEntity) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr
func (mr *MockNetworkEntityMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockNetworkEntity)(nil).RemoteAddr))
}

// ResponseMID mocks base method
func (m *MockNetworkEntity) ResponseMID(arg0 context.Context, arg1 uint, arg2 interface{}, arg3 ...bool) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResponseMID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResponseMID indicates an expected call of ResponseMID
func (mr *MockNetworkEntityMockRecorder) ResponseMID(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseMID", reflect.TypeOf((*MockNetworkEntity)(nil).ResponseMID), varargs...)
}

// SendRequest mocks base method
func (m *MockNetworkEntity) SendRequest(arg0 context.Context, arg1, arg2 string, arg3 interface{}) (*protos.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRequest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRequest indicates an expected call of SendRequest
func (mr *MockNetworkEntityMockRecorder) SendRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRequest", reflect.TypeOf((*MockNetworkEntity)(nil).SendRequest), arg0, arg1, arg2, arg3)
}
