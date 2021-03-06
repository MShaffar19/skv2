// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	controller "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterPaintReconciler is a mock of MulticlusterPaintReconciler interface
type MockMulticlusterPaintReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPaintReconcilerMockRecorder
}

// MockMulticlusterPaintReconcilerMockRecorder is the mock recorder for MockMulticlusterPaintReconciler
type MockMulticlusterPaintReconcilerMockRecorder struct {
	mock *MockMulticlusterPaintReconciler
}

// NewMockMulticlusterPaintReconciler creates a new mock instance
func NewMockMulticlusterPaintReconciler(ctrl *gomock.Controller) *MockMulticlusterPaintReconciler {
	mock := &MockMulticlusterPaintReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPaintReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPaintReconciler) EXPECT() *MockMulticlusterPaintReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePaint mocks base method
func (m *MockMulticlusterPaintReconciler) ReconcilePaint(clusterName string, obj *v1.Paint) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePaint", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePaint indicates an expected call of ReconcilePaint
func (mr *MockMulticlusterPaintReconcilerMockRecorder) ReconcilePaint(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePaint", reflect.TypeOf((*MockMulticlusterPaintReconciler)(nil).ReconcilePaint), clusterName, obj)
}

// MockMulticlusterPaintDeletionReconciler is a mock of MulticlusterPaintDeletionReconciler interface
type MockMulticlusterPaintDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPaintDeletionReconcilerMockRecorder
}

// MockMulticlusterPaintDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterPaintDeletionReconciler
type MockMulticlusterPaintDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterPaintDeletionReconciler
}

// NewMockMulticlusterPaintDeletionReconciler creates a new mock instance
func NewMockMulticlusterPaintDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterPaintDeletionReconciler {
	mock := &MockMulticlusterPaintDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPaintDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPaintDeletionReconciler) EXPECT() *MockMulticlusterPaintDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePaintDeletion mocks base method
func (m *MockMulticlusterPaintDeletionReconciler) ReconcilePaintDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePaintDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcilePaintDeletion indicates an expected call of ReconcilePaintDeletion
func (mr *MockMulticlusterPaintDeletionReconcilerMockRecorder) ReconcilePaintDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePaintDeletion", reflect.TypeOf((*MockMulticlusterPaintDeletionReconciler)(nil).ReconcilePaintDeletion), clusterName, req)
}

// MockMulticlusterPaintReconcileLoop is a mock of MulticlusterPaintReconcileLoop interface
type MockMulticlusterPaintReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPaintReconcileLoopMockRecorder
}

// MockMulticlusterPaintReconcileLoopMockRecorder is the mock recorder for MockMulticlusterPaintReconcileLoop
type MockMulticlusterPaintReconcileLoopMockRecorder struct {
	mock *MockMulticlusterPaintReconcileLoop
}

// NewMockMulticlusterPaintReconcileLoop creates a new mock instance
func NewMockMulticlusterPaintReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterPaintReconcileLoop {
	mock := &MockMulticlusterPaintReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPaintReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPaintReconcileLoop) EXPECT() *MockMulticlusterPaintReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterPaintReconciler mocks base method
func (m *MockMulticlusterPaintReconcileLoop) AddMulticlusterPaintReconciler(ctx context.Context, rec controller.MulticlusterPaintReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterPaintReconciler", varargs...)
}

// AddMulticlusterPaintReconciler indicates an expected call of AddMulticlusterPaintReconciler
func (mr *MockMulticlusterPaintReconcileLoopMockRecorder) AddMulticlusterPaintReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterPaintReconciler", reflect.TypeOf((*MockMulticlusterPaintReconcileLoop)(nil).AddMulticlusterPaintReconciler), varargs...)
}

// MockMulticlusterClusterResourceReconciler is a mock of MulticlusterClusterResourceReconciler interface
type MockMulticlusterClusterResourceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterResourceReconcilerMockRecorder
}

// MockMulticlusterClusterResourceReconcilerMockRecorder is the mock recorder for MockMulticlusterClusterResourceReconciler
type MockMulticlusterClusterResourceReconcilerMockRecorder struct {
	mock *MockMulticlusterClusterResourceReconciler
}

// NewMockMulticlusterClusterResourceReconciler creates a new mock instance
func NewMockMulticlusterClusterResourceReconciler(ctrl *gomock.Controller) *MockMulticlusterClusterResourceReconciler {
	mock := &MockMulticlusterClusterResourceReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterResourceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterResourceReconciler) EXPECT() *MockMulticlusterClusterResourceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileClusterResource mocks base method
func (m *MockMulticlusterClusterResourceReconciler) ReconcileClusterResource(clusterName string, obj *v1.ClusterResource) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterResource", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileClusterResource indicates an expected call of ReconcileClusterResource
func (mr *MockMulticlusterClusterResourceReconcilerMockRecorder) ReconcileClusterResource(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterResource", reflect.TypeOf((*MockMulticlusterClusterResourceReconciler)(nil).ReconcileClusterResource), clusterName, obj)
}

// MockMulticlusterClusterResourceDeletionReconciler is a mock of MulticlusterClusterResourceDeletionReconciler interface
type MockMulticlusterClusterResourceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterResourceDeletionReconcilerMockRecorder
}

// MockMulticlusterClusterResourceDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterClusterResourceDeletionReconciler
type MockMulticlusterClusterResourceDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterClusterResourceDeletionReconciler
}

// NewMockMulticlusterClusterResourceDeletionReconciler creates a new mock instance
func NewMockMulticlusterClusterResourceDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterClusterResourceDeletionReconciler {
	mock := &MockMulticlusterClusterResourceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterResourceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterResourceDeletionReconciler) EXPECT() *MockMulticlusterClusterResourceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileClusterResourceDeletion mocks base method
func (m *MockMulticlusterClusterResourceDeletionReconciler) ReconcileClusterResourceDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterResourceDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileClusterResourceDeletion indicates an expected call of ReconcileClusterResourceDeletion
func (mr *MockMulticlusterClusterResourceDeletionReconcilerMockRecorder) ReconcileClusterResourceDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterResourceDeletion", reflect.TypeOf((*MockMulticlusterClusterResourceDeletionReconciler)(nil).ReconcileClusterResourceDeletion), clusterName, req)
}

// MockMulticlusterClusterResourceReconcileLoop is a mock of MulticlusterClusterResourceReconcileLoop interface
type MockMulticlusterClusterResourceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterResourceReconcileLoopMockRecorder
}

// MockMulticlusterClusterResourceReconcileLoopMockRecorder is the mock recorder for MockMulticlusterClusterResourceReconcileLoop
type MockMulticlusterClusterResourceReconcileLoopMockRecorder struct {
	mock *MockMulticlusterClusterResourceReconcileLoop
}

// NewMockMulticlusterClusterResourceReconcileLoop creates a new mock instance
func NewMockMulticlusterClusterResourceReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterClusterResourceReconcileLoop {
	mock := &MockMulticlusterClusterResourceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterResourceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterResourceReconcileLoop) EXPECT() *MockMulticlusterClusterResourceReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterClusterResourceReconciler mocks base method
func (m *MockMulticlusterClusterResourceReconcileLoop) AddMulticlusterClusterResourceReconciler(ctx context.Context, rec controller.MulticlusterClusterResourceReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterClusterResourceReconciler", varargs...)
}

// AddMulticlusterClusterResourceReconciler indicates an expected call of AddMulticlusterClusterResourceReconciler
func (mr *MockMulticlusterClusterResourceReconcileLoopMockRecorder) AddMulticlusterClusterResourceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterClusterResourceReconciler", reflect.TypeOf((*MockMulticlusterClusterResourceReconcileLoop)(nil).AddMulticlusterClusterResourceReconciler), varargs...)
}
