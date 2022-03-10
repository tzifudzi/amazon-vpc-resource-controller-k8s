// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-resource-controller-k8s/pkg/condition (interfaces: Conditions)

// Package mock_condition is a generated GoMock package.
package mock_condition

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConditions is a mock of Conditions interface.
type MockConditions struct {
	ctrl     *gomock.Controller
	recorder *MockConditionsMockRecorder
}

// MockConditionsMockRecorder is the mock recorder for MockConditions.
type MockConditionsMockRecorder struct {
	mock *MockConditions
}

// NewMockConditions creates a new mock instance.
func NewMockConditions(ctrl *gomock.Controller) *MockConditions {
	mock := &MockConditions{ctrl: ctrl}
	mock.recorder = &MockConditionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConditions) EXPECT() *MockConditionsMockRecorder {
	return m.recorder
}

// IsOldVPCControllerDeploymentPresent mocks base method.
func (m *MockConditions) IsOldVPCControllerDeploymentPresent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOldVPCControllerDeploymentPresent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOldVPCControllerDeploymentPresent indicates an expected call of IsOldVPCControllerDeploymentPresent.
func (mr *MockConditionsMockRecorder) IsOldVPCControllerDeploymentPresent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOldVPCControllerDeploymentPresent", reflect.TypeOf((*MockConditions)(nil).IsOldVPCControllerDeploymentPresent))
}

// IsPodSGPEnabled mocks base method.
func (m *MockConditions) IsPodSGPEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPodSGPEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPodSGPEnabled indicates an expected call of IsPodSGPEnabled.
func (mr *MockConditionsMockRecorder) IsPodSGPEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPodSGPEnabled", reflect.TypeOf((*MockConditions)(nil).IsPodSGPEnabled))
}

// IsWindowsIPAMEnabled mocks base method.
func (m *MockConditions) IsWindowsIPAMEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWindowsIPAMEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsWindowsIPAMEnabled indicates an expected call of IsWindowsIPAMEnabled.
func (mr *MockConditionsMockRecorder) IsWindowsIPAMEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWindowsIPAMEnabled", reflect.TypeOf((*MockConditions)(nil).IsWindowsIPAMEnabled))
}

// WaitTillPodDataStoreSynced mocks base method.
func (m *MockConditions) WaitTillPodDataStoreSynced() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WaitTillPodDataStoreSynced")
}

// WaitTillPodDataStoreSynced indicates an expected call of WaitTillPodDataStoreSynced.
func (mr *MockConditionsMockRecorder) WaitTillPodDataStoreSynced() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitTillPodDataStoreSynced", reflect.TypeOf((*MockConditions)(nil).WaitTillPodDataStoreSynced))
}
