// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	discovery_multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/discovery.multicluster.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AwsDiscoveryDirective Resource across clusters.
// implemented by the user
type MulticlusterAwsDiscoveryDirectiveReconciler interface {
	ReconcileAwsDiscoveryDirective(clusterName string, obj *discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryDirective) (reconcile.Result, error)
}

// Reconcile deletion events for the AwsDiscoveryDirective Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterAwsDiscoveryDirectiveDeletionReconciler interface {
	ReconcileAwsDiscoveryDirectiveDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterAwsDiscoveryDirectiveReconcilerFuncs struct {
	OnReconcileAwsDiscoveryDirective         func(clusterName string, obj *discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryDirective) (reconcile.Result, error)
	OnReconcileAwsDiscoveryDirectiveDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterAwsDiscoveryDirectiveReconcilerFuncs) ReconcileAwsDiscoveryDirective(clusterName string, obj *discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryDirective) (reconcile.Result, error) {
	if f.OnReconcileAwsDiscoveryDirective == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAwsDiscoveryDirective(clusterName, obj)
}

func (f *MulticlusterAwsDiscoveryDirectiveReconcilerFuncs) ReconcileAwsDiscoveryDirectiveDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileAwsDiscoveryDirectiveDeletion == nil {
		return nil
	}
	return f.OnReconcileAwsDiscoveryDirectiveDeletion(clusterName, req)
}

type MulticlusterAwsDiscoveryDirectiveReconcileLoop interface {
	// AddMulticlusterAwsDiscoveryDirectiveReconciler adds a MulticlusterAwsDiscoveryDirectiveReconciler to the MulticlusterAwsDiscoveryDirectiveReconcileLoop.
	AddMulticlusterAwsDiscoveryDirectiveReconciler(ctx context.Context, rec MulticlusterAwsDiscoveryDirectiveReconciler, predicates ...predicate.Predicate)
}

type multiclusterAwsDiscoveryDirectiveReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterAwsDiscoveryDirectiveReconcileLoop) AddMulticlusterAwsDiscoveryDirectiveReconciler(ctx context.Context, rec MulticlusterAwsDiscoveryDirectiveReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericAwsDiscoveryDirectiveMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterAwsDiscoveryDirectiveReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterAwsDiscoveryDirectiveReconcileLoop {
	return &multiclusterAwsDiscoveryDirectiveReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryDirective{})}
}

type genericAwsDiscoveryDirectiveMulticlusterReconciler struct {
	reconciler MulticlusterAwsDiscoveryDirectiveReconciler
}

func (g genericAwsDiscoveryDirectiveMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterAwsDiscoveryDirectiveDeletionReconciler); ok {
		return deletionReconciler.ReconcileAwsDiscoveryDirectiveDeletion(cluster, req)
	}
	return nil
}

func (g genericAwsDiscoveryDirectiveMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryDirective)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AwsDiscoveryDirective handler received event for %T", object)
	}
	return g.reconciler.ReconcileAwsDiscoveryDirective(cluster, obj)
}
