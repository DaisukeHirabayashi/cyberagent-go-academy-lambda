// Code generated by MockGen. DO NOT EDIT.
// Source: client/corona.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCoronaClientInterface is a mock of CoronaClientInterface interface.
type MockCoronaClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCoronaClientInterfaceMockRecorder
}

// MockCoronaClientInterfaceMockRecorder is the mock recorder for MockCoronaClientInterface.
type MockCoronaClientInterfaceMockRecorder struct {
	mock *MockCoronaClientInterface
}

// NewMockCoronaClientInterface creates a new mock instance.
func NewMockCoronaClientInterface(ctrl *gomock.Controller) *MockCoronaClientInterface {
	mock := &MockCoronaClientInterface{ctrl: ctrl}
	mock.recorder = &MockCoronaClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoronaClientInterface) EXPECT() *MockCoronaClientInterfaceMockRecorder {
	return m.recorder
}

// GetMedicalSystem mocks base method.
func (m *MockCoronaClientInterface) GetMedicalSystem(time *string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMedicalSystem", time)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMedicalSystem indicates an expected call of GetMedicalSystem.
func (mr *MockCoronaClientInterfaceMockRecorder) GetMedicalSystem(time interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMedicalSystem", reflect.TypeOf((*MockCoronaClientInterface)(nil).GetMedicalSystem), time)
}
