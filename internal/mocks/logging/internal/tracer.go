// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MerlinKodo/quic-go/internal/mocks/logging (interfaces: Tracer)
//
// Generated by this command:
//
//	mockgen -build_flags=-tags=gomock -package internal -destination internal/tracer.go github.com/MerlinKodo/quic-go/internal/mocks/logging Tracer
//
// Package internal is a generated GoMock package.
package internal

import (
	net "net"
	reflect "reflect"

	protocol "github.com/MerlinKodo/quic-go/internal/protocol"
	wire "github.com/MerlinKodo/quic-go/internal/wire"
	logging "github.com/MerlinKodo/quic-go/logging"
	gomock "go.uber.org/mock/gomock"
)

// MockTracer is a mock of Tracer interface.
type MockTracer struct {
	ctrl     *gomock.Controller
	recorder *MockTracerMockRecorder
}

// MockTracerMockRecorder is the mock recorder for MockTracer.
type MockTracerMockRecorder struct {
	mock *MockTracer
}

// NewMockTracer creates a new mock instance.
func NewMockTracer(ctrl *gomock.Controller) *MockTracer {
	mock := &MockTracer{ctrl: ctrl}
	mock.recorder = &MockTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTracer) EXPECT() *MockTracerMockRecorder {
	return m.recorder
}

// DroppedPacket mocks base method.
func (m *MockTracer) DroppedPacket(arg0 net.Addr, arg1 logging.PacketType, arg2 protocol.ByteCount, arg3 logging.PacketDropReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DroppedPacket", arg0, arg1, arg2, arg3)
}

// DroppedPacket indicates an expected call of DroppedPacket.
func (mr *MockTracerMockRecorder) DroppedPacket(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DroppedPacket", reflect.TypeOf((*MockTracer)(nil).DroppedPacket), arg0, arg1, arg2, arg3)
}

// SentPacket mocks base method.
func (m *MockTracer) SentPacket(arg0 net.Addr, arg1 *wire.Header, arg2 protocol.ByteCount, arg3 []logging.Frame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentPacket", arg0, arg1, arg2, arg3)
}

// SentPacket indicates an expected call of SentPacket.
func (mr *MockTracerMockRecorder) SentPacket(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentPacket", reflect.TypeOf((*MockTracer)(nil).SentPacket), arg0, arg1, arg2, arg3)
}

// SentVersionNegotiationPacket mocks base method.
func (m *MockTracer) SentVersionNegotiationPacket(arg0 net.Addr, arg1, arg2 protocol.ArbitraryLenConnectionID, arg3 []protocol.VersionNumber) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentVersionNegotiationPacket", arg0, arg1, arg2, arg3)
}

// SentVersionNegotiationPacket indicates an expected call of SentVersionNegotiationPacket.
func (mr *MockTracerMockRecorder) SentVersionNegotiationPacket(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentVersionNegotiationPacket", reflect.TypeOf((*MockTracer)(nil).SentVersionNegotiationPacket), arg0, arg1, arg2, arg3)
}
