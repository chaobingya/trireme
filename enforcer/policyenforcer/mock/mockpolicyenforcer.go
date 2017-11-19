// Code generated by MockGen. DO NOT EDIT.
// Source: enforcer/policyenforcer/interfaces.go

// Package mockpolicyenforcer is a generated GoMock package.
package mockpolicyenforcer

import (
	reflect "reflect"

	fqconfig "github.com/aporeto-inc/trireme-lib/enforcer/utils/fqconfig"
	policy "github.com/aporeto-inc/trireme-lib/policy"
	gomock "github.com/golang/mock/gomock"
)

// MockEnforcer is a mock of Enforcer interface
// nolint
type MockEnforcer struct {
	ctrl     *gomock.Controller
	recorder *MockEnforcerMockRecorder
}

// MockEnforcerMockRecorder is the mock recorder for MockEnforcer
// nolint
type MockEnforcerMockRecorder struct {
	mock *MockEnforcer
}

// NewMockEnforcer creates a new mock instance
// nolint
func NewMockEnforcer(ctrl *gomock.Controller) *MockEnforcer {
	mock := &MockEnforcer{ctrl: ctrl}
	mock.recorder = &MockEnforcerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockEnforcer) EXPECT() *MockEnforcerMockRecorder {
	return m.recorder
}

// Enforce mocks base method
// nolint
func (m *MockEnforcer) Enforce(contextID string, puInfo *policy.PUInfo) error {
	ret := m.ctrl.Call(m, "Enforce", contextID, puInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enforce indicates an expected call of Enforce
// nolint
func (mr *MockEnforcerMockRecorder) Enforce(contextID, puInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enforce", reflect.TypeOf((*MockEnforcer)(nil).Enforce), contextID, puInfo)
}

// Unenforce mocks base method
// nolint
func (m *MockEnforcer) Unenforce(contextID string) error {
	ret := m.ctrl.Call(m, "Unenforce", contextID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unenforce indicates an expected call of Unenforce
// nolint
func (mr *MockEnforcerMockRecorder) Unenforce(contextID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unenforce", reflect.TypeOf((*MockEnforcer)(nil).Unenforce), contextID)
}

// GetFilterQueue mocks base method
// nolint
func (m *MockEnforcer) GetFilterQueue() *fqconfig.FilterQueue {
	ret := m.ctrl.Call(m, "GetFilterQueue")
	ret0, _ := ret[0].(*fqconfig.FilterQueue)
	return ret0
}

// GetFilterQueue indicates an expected call of GetFilterQueue
// nolint
func (mr *MockEnforcerMockRecorder) GetFilterQueue() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilterQueue", reflect.TypeOf((*MockEnforcer)(nil).GetFilterQueue))
}

// Start mocks base method
// nolint
func (m *MockEnforcer) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
// nolint
func (mr *MockEnforcerMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEnforcer)(nil).Start))
}

// Stop mocks base method
// nolint
func (m *MockEnforcer) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
// nolint
func (mr *MockEnforcerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockEnforcer)(nil).Stop))
}