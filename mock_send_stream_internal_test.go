// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/metacubex/quic-go (interfaces: SendStreamI)
//
// Generated by this command:
//
//	mockgen -build_flags=-tags=gomock -package quic -self_package github.com/metacubex/quic-go -destination mock_send_stream_internal_test.go github.com/metacubex/quic-go SendStreamI
//
// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	reflect "reflect"
	time "time"

	ackhandler "github.com/metacubex/quic-go/internal/ackhandler"
	protocol "github.com/metacubex/quic-go/internal/protocol"
	qerr "github.com/metacubex/quic-go/internal/qerr"
	wire "github.com/metacubex/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockSendStreamI is a mock of SendStreamI interface.
type MockSendStreamI struct {
	ctrl     *gomock.Controller
	recorder *MockSendStreamIMockRecorder
}

// MockSendStreamIMockRecorder is the mock recorder for MockSendStreamI.
type MockSendStreamIMockRecorder struct {
	mock *MockSendStreamI
}

// NewMockSendStreamI creates a new mock instance.
func NewMockSendStreamI(ctrl *gomock.Controller) *MockSendStreamI {
	mock := &MockSendStreamI{ctrl: ctrl}
	mock.recorder = &MockSendStreamIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendStreamI) EXPECT() *MockSendStreamIMockRecorder {
	return m.recorder
}

// CancelWrite mocks base method.
func (m *MockSendStreamI) CancelWrite(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelWrite", arg0)
}

// CancelWrite indicates an expected call of CancelWrite.
func (mr *MockSendStreamIMockRecorder) CancelWrite(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWrite", reflect.TypeOf((*MockSendStreamI)(nil).CancelWrite), arg0)
}

// Close mocks base method.
func (m *MockSendStreamI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSendStreamIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSendStreamI)(nil).Close))
}

// Context mocks base method.
func (m *MockSendStreamI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSendStreamIMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSendStreamI)(nil).Context))
}

// SetWriteDeadline mocks base method.
func (m *MockSendStreamI) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockSendStreamIMockRecorder) SetWriteDeadline(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockSendStreamI)(nil).SetWriteDeadline), arg0)
}

// StreamID mocks base method.
func (m *MockSendStreamI) StreamID() protocol.StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockSendStreamIMockRecorder) StreamID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockSendStreamI)(nil).StreamID))
}

// Write mocks base method.
func (m *MockSendStreamI) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockSendStreamIMockRecorder) Write(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSendStreamI)(nil).Write), arg0)
}

// closeForShutdown mocks base method.
func (m *MockSendStreamI) closeForShutdown(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeForShutdown", arg0)
}

// closeForShutdown indicates an expected call of closeForShutdown.
func (mr *MockSendStreamIMockRecorder) closeForShutdown(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeForShutdown", reflect.TypeOf((*MockSendStreamI)(nil).closeForShutdown), arg0)
}

// handleStopSendingFrame mocks base method.
func (m *MockSendStreamI) handleStopSendingFrame(arg0 *wire.StopSendingFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleStopSendingFrame", arg0)
}

// handleStopSendingFrame indicates an expected call of handleStopSendingFrame.
func (mr *MockSendStreamIMockRecorder) handleStopSendingFrame(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStopSendingFrame", reflect.TypeOf((*MockSendStreamI)(nil).handleStopSendingFrame), arg0)
}

// hasData mocks base method.
func (m *MockSendStreamI) hasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// hasData indicates an expected call of hasData.
func (mr *MockSendStreamIMockRecorder) hasData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hasData", reflect.TypeOf((*MockSendStreamI)(nil).hasData))
}

// popStreamFrame mocks base method.
func (m *MockSendStreamI) popStreamFrame(arg0 protocol.ByteCount, arg1 protocol.VersionNumber) (ackhandler.StreamFrame, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "popStreamFrame", arg0, arg1)
	ret0, _ := ret[0].(ackhandler.StreamFrame)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// popStreamFrame indicates an expected call of popStreamFrame.
func (mr *MockSendStreamIMockRecorder) popStreamFrame(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "popStreamFrame", reflect.TypeOf((*MockSendStreamI)(nil).popStreamFrame), arg0, arg1)
}

// updateSendWindow mocks base method.
func (m *MockSendStreamI) updateSendWindow(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateSendWindow", arg0)
}

// updateSendWindow indicates an expected call of updateSendWindow.
func (mr *MockSendStreamIMockRecorder) updateSendWindow(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateSendWindow", reflect.TypeOf((*MockSendStreamI)(nil).updateSendWindow), arg0)
}
