// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/cluster_upgrader (interfaces: ClusterUpgrader)

// Package mocks is a generated GoMock package.
package mocks

import (
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	maintenance "github.com/openshift/managed-upgrade-operator/pkg/maintenance"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClusterUpgrader is a mock of ClusterUpgrader interface
type MockClusterUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockClusterUpgraderMockRecorder
}

// MockClusterUpgraderMockRecorder is the mock recorder for MockClusterUpgrader
type MockClusterUpgraderMockRecorder struct {
	mock *MockClusterUpgrader
}

// NewMockClusterUpgrader creates a new mock instance
func NewMockClusterUpgrader(ctrl *gomock.Controller) *MockClusterUpgrader {
	mock := &MockClusterUpgrader{ctrl: ctrl}
	mock.recorder = &MockClusterUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterUpgrader) EXPECT() *MockClusterUpgraderMockRecorder {
	return m.recorder
}

// UpgradeCluster mocks base method
func (m *MockClusterUpgrader) UpgradeCluster(arg0 client.Client, arg1 maintenance.Maintenance, arg2 *v1alpha1.UpgradeConfig, arg3 logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeCluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpgradeCluster indicates an expected call of UpgradeCluster
func (mr *MockClusterUpgraderMockRecorder) UpgradeCluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeCluster", reflect.TypeOf((*MockClusterUpgrader)(nil).UpgradeCluster), arg0, arg1, arg2, arg3)
}