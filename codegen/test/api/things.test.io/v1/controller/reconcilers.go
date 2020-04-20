// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Paint Resource.
// implemented by the user
type PaintReconciler interface {
	ReconcilePaint(obj *things_test_io_v1.Paint) (reconcile.Result, error)
}

// Reconcile deletion events for the Paint Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type PaintDeletionReconciler interface {
	ReconcilePaintDeletion(req reconcile.Request)
}

type PaintReconcilerFuncs struct {
	OnReconcilePaint         func(obj *things_test_io_v1.Paint) (reconcile.Result, error)
	OnReconcilePaintDeletion func(req reconcile.Request)
}

func (f *PaintReconcilerFuncs) ReconcilePaint(obj *things_test_io_v1.Paint) (reconcile.Result, error) {
	if f.OnReconcilePaint == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcilePaint(obj)
}

func (f *PaintReconcilerFuncs) ReconcilePaintDeletion(req reconcile.Request) {
	if f.OnReconcilePaintDeletion == nil {
		return
	}
	f.OnReconcilePaintDeletion(req)
}

// Reconcile and finalize the Paint Resource
// implemented by the user
type PaintFinalizer interface {
	PaintReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	PaintFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizePaint(obj *things_test_io_v1.Paint) error
}

type PaintReconcileLoop interface {
	RunPaintReconciler(ctx context.Context, rec PaintReconciler, predicates ...predicate.Predicate) error
}

type paintReconcileLoop struct {
	loop reconcile.Loop
}

func NewPaintReconcileLoop(name string, mgr manager.Manager) PaintReconcileLoop {
	return &paintReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &things_test_io_v1.Paint{}),
	}
}

func (c *paintReconcileLoop) RunPaintReconciler(ctx context.Context, reconciler PaintReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericPaintReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(PaintFinalizer); ok {
		reconcilerWrapper = genericPaintFinalizer{
			genericPaintReconciler: genericReconciler,
			finalizingReconciler:   finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericPaintHandler implements a generic reconcile.Reconciler
type genericPaintReconciler struct {
	reconciler PaintReconciler
}

func (r genericPaintReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return r.reconciler.ReconcilePaint(obj)
}

func (r genericPaintReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(PaintDeletionReconciler); ok {
		deletionReconciler.ReconcilePaintDeletion(request)
	}
}

// genericPaintFinalizer implements a generic reconcile.FinalizingReconciler
type genericPaintFinalizer struct {
	genericPaintReconciler
	finalizingReconciler PaintFinalizer
}

func (r genericPaintFinalizer) FinalizerName() string {
	return r.finalizingReconciler.PaintFinalizerName()
}

func (r genericPaintFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizePaint(obj)
}