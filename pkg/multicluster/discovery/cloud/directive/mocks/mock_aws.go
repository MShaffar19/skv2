// Code generated by MockGen. DO NOT EDIT.
// Source: ./aws.go

// Package mock_directive is a generated GoMock package.
package mock_directive

import (
	context "context"
	reflect "reflect"

	credentials "github.com/aws/aws-sdk-go/aws/credentials"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/skv2/pkg/api/discovery.multicluster.solo.io/v1alpha1"
	directive "github.com/solo-io/skv2/pkg/multicluster/discovery/cloud/directive"
)

// MockAwsDiscoveryResolver is a mock of AwsDiscoveryResolver interface
type MockAwsDiscoveryResolver struct {
	ctrl     *gomock.Controller
	recorder *MockAwsDiscoveryResolverMockRecorder
}

// MockAwsDiscoveryResolverMockRecorder is the mock recorder for MockAwsDiscoveryResolver
type MockAwsDiscoveryResolverMockRecorder struct {
	mock *MockAwsDiscoveryResolver
}

// NewMockAwsDiscoveryResolver creates a new mock instance
func NewMockAwsDiscoveryResolver(ctrl *gomock.Controller) *MockAwsDiscoveryResolver {
	mock := &MockAwsDiscoveryResolver{ctrl: ctrl}
	mock.recorder = &MockAwsDiscoveryResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAwsDiscoveryResolver) EXPECT() *MockAwsDiscoveryResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method
func (m *MockAwsDiscoveryResolver) Resolve(ctx context.Context, creds *credentials.Credentials, awsDirectives []*v1alpha1.AwsDiscoveryDirectiveSpec) directive.AwsResources {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", ctx, creds, awsDirectives)
	ret0, _ := ret[0].(directive.AwsResources)
	return ret0
}

// Resolve indicates an expected call of Resolve
func (mr *MockAwsDiscoveryResolverMockRecorder) Resolve(ctx, creds, awsDirectives interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockAwsDiscoveryResolver)(nil).Resolve), ctx, creds, awsDirectives)
}
