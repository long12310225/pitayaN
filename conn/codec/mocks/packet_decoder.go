// Code generated by MockGen. DO NOT EDIT.
// Source: conn/codec/packet_decoder.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	packet "github.com/long12310225/pitayaN/v2/conn/packet"
	reflect "reflect"
)

// MockPacketDecoder is a mock of PacketDecoder interface
type MockPacketDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockPacketDecoderMockRecorder
}

// MockPacketDecoderMockRecorder is the mock recorder for MockPacketDecoder
type MockPacketDecoderMockRecorder struct {
	mock *MockPacketDecoder
}

// NewMockPacketDecoder creates a new mock instance
func NewMockPacketDecoder(ctrl *gomock.Controller) *MockPacketDecoder {
	mock := &MockPacketDecoder{ctrl: ctrl}
	mock.recorder = &MockPacketDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacketDecoder) EXPECT() *MockPacketDecoderMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockPacketDecoder) Decode(data []byte) ([]*packet.Packet, error) {
	ret := m.ctrl.Call(m, "Decode", data)
	ret0, _ := ret[0].([]*packet.Packet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockPacketDecoderMockRecorder) Decode(data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockPacketDecoder)(nil).Decode), data)
}
