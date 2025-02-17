// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MerlinKodo/quic-go (interfaces: Packer)
//
// Generated by this command:
//
//	mockgen -build_flags=-tags=gomock -package quic -self_package github.com/MerlinKodo/quic-go -destination mock_packer_test.go github.com/MerlinKodo/quic-go Packer
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	ackhandler "github.com/MerlinKodo/quic-go/internal/ackhandler"
	protocol "github.com/MerlinKodo/quic-go/internal/protocol"
	qerr "github.com/MerlinKodo/quic-go/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockPacker is a mock of Packer interface.
type MockPacker struct {
	ctrl     *gomock.Controller
	recorder *MockPackerMockRecorder
}

// MockPackerMockRecorder is the mock recorder for MockPacker.
type MockPackerMockRecorder struct {
	mock *MockPacker
}

// NewMockPacker creates a new mock instance.
func NewMockPacker(ctrl *gomock.Controller) *MockPacker {
	mock := &MockPacker{ctrl: ctrl}
	mock.recorder = &MockPackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacker) EXPECT() *MockPackerMockRecorder {
	return m.recorder
}

// AppendPacket mocks base method.
func (m *MockPacker) AppendPacket(arg0 *packetBuffer, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (shortHeaderPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendPacket indicates an expected call of AppendPacket.
func (mr *MockPackerMockRecorder) AppendPacket(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendPacket", reflect.TypeOf((*MockPacker)(nil).AppendPacket), arg0, arg1, arg2)
}

// MaybePackProbePacket mocks base method.
func (m *MockPacker) MaybePackProbePacket(arg0 protocol.EncryptionLevel, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePackProbePacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybePackProbePacket indicates an expected call of MaybePackProbePacket.
func (mr *MockPackerMockRecorder) MaybePackProbePacket(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePackProbePacket", reflect.TypeOf((*MockPacker)(nil).MaybePackProbePacket), arg0, arg1, arg2)
}

// PackAckOnlyPacket mocks base method.
func (m *MockPacker) PackAckOnlyPacket(arg0 protocol.ByteCount, arg1 protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackAckOnlyPacket", arg0, arg1)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackAckOnlyPacket indicates an expected call of PackAckOnlyPacket.
func (mr *MockPackerMockRecorder) PackAckOnlyPacket(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackAckOnlyPacket", reflect.TypeOf((*MockPacker)(nil).PackAckOnlyPacket), arg0, arg1)
}

// PackApplicationClose mocks base method.
func (m *MockPacker) PackApplicationClose(arg0 *qerr.ApplicationError, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackApplicationClose", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackApplicationClose indicates an expected call of PackApplicationClose.
func (mr *MockPackerMockRecorder) PackApplicationClose(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackApplicationClose", reflect.TypeOf((*MockPacker)(nil).PackApplicationClose), arg0, arg1, arg2)
}

// PackCoalescedPacket mocks base method.
func (m *MockPacker) PackCoalescedPacket(arg0 bool, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackCoalescedPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackCoalescedPacket indicates an expected call of PackCoalescedPacket.
func (mr *MockPackerMockRecorder) PackCoalescedPacket(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackCoalescedPacket", reflect.TypeOf((*MockPacker)(nil).PackCoalescedPacket), arg0, arg1, arg2)
}

// PackConnectionClose mocks base method.
func (m *MockPacker) PackConnectionClose(arg0 *qerr.TransportError, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackConnectionClose", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackConnectionClose indicates an expected call of PackConnectionClose.
func (mr *MockPackerMockRecorder) PackConnectionClose(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackConnectionClose", reflect.TypeOf((*MockPacker)(nil).PackConnectionClose), arg0, arg1, arg2)
}

// PackMTUProbePacket mocks base method.
func (m *MockPacker) PackMTUProbePacket(arg0 ackhandler.Frame, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackMTUProbePacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackMTUProbePacket indicates an expected call of PackMTUProbePacket.
func (mr *MockPackerMockRecorder) PackMTUProbePacket(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackMTUProbePacket", reflect.TypeOf((*MockPacker)(nil).PackMTUProbePacket), arg0, arg1, arg2)
}

// SetToken mocks base method.
func (m *MockPacker) SetToken(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToken", arg0)
}

// SetToken indicates an expected call of SetToken.
func (mr *MockPackerMockRecorder) SetToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockPacker)(nil).SetToken), arg0)
}
