// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type PaintSet interface {
	Keys() sets.String
	List() []*things_test_io_v1.Paint
	Map() map[string]*things_test_io_v1.Paint
	Insert(paint ...*things_test_io_v1.Paint)
	Equal(paintSet PaintSet) bool
	Has(paint *things_test_io_v1.Paint) bool
	Delete(paint *things_test_io_v1.Paint)
	Union(set PaintSet) PaintSet
	Difference(set PaintSet) PaintSet
	Intersection(set PaintSet) PaintSet
}

func makeGenericPaintSet(paintList []*things_test_io_v1.Paint) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range paintList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type paintSet struct {
	set sksets.ResourceSet
}

func NewPaintSet(paintList ...*things_test_io_v1.Paint) PaintSet {
	return &paintSet{set: makeGenericPaintSet(paintList)}
}

func (s paintSet) Keys() sets.String {
	return s.set.Keys()
}

func (s paintSet) List() []*things_test_io_v1.Paint {
	var paintList []*things_test_io_v1.Paint
	for _, obj := range s.set.List() {
		paintList = append(paintList, obj.(*things_test_io_v1.Paint))
	}
	return paintList
}

func (s paintSet) Map() map[string]*things_test_io_v1.Paint {
	newMap := map[string]*things_test_io_v1.Paint{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*things_test_io_v1.Paint)
	}
	return newMap
}

func (s paintSet) Insert(
	paintList ...*things_test_io_v1.Paint,
) {
	for _, obj := range paintList {
		s.set.Insert(obj)
	}
}

func (s paintSet) Has(paint *things_test_io_v1.Paint) bool {
	return s.set.Has(paint)
}

func (s paintSet) Equal(
	paintSet PaintSet,
) bool {
	return s.set.Equal(makeGenericPaintSet(paintSet.List()))
}

func (s paintSet) Delete(Paint *things_test_io_v1.Paint) {
	s.set.Delete(Paint)
}

func (s paintSet) Union(set PaintSet) PaintSet {
	return NewPaintSet(append(s.List(), set.List()...)...)
}

func (s paintSet) Difference(set PaintSet) PaintSet {
	newSet := s.set.Difference(makeGenericPaintSet(set.List()))
	return paintSet{set: newSet}
}

func (s paintSet) Intersection(set PaintSet) PaintSet {
	newSet := s.set.Intersection(makeGenericPaintSet(set.List()))
	var paintList []*things_test_io_v1.Paint
	for _, obj := range newSet.List() {
		paintList = append(paintList, obj.(*things_test_io_v1.Paint))
	}
	return NewPaintSet(paintList...)
}

type ClusterResourceSet interface {
	Keys() sets.String
	List() []*things_test_io_v1.ClusterResource
	Map() map[string]*things_test_io_v1.ClusterResource
	Insert(clusterResource ...*things_test_io_v1.ClusterResource)
	Equal(clusterResourceSet ClusterResourceSet) bool
	Has(clusterResource *things_test_io_v1.ClusterResource) bool
	Delete(clusterResource *things_test_io_v1.ClusterResource)
	Union(set ClusterResourceSet) ClusterResourceSet
	Difference(set ClusterResourceSet) ClusterResourceSet
	Intersection(set ClusterResourceSet) ClusterResourceSet
}

func makeGenericClusterResourceSet(clusterResourceList []*things_test_io_v1.ClusterResource) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range clusterResourceList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type clusterResourceSet struct {
	set sksets.ResourceSet
}

func NewClusterResourceSet(clusterResourceList ...*things_test_io_v1.ClusterResource) ClusterResourceSet {
	return &clusterResourceSet{set: makeGenericClusterResourceSet(clusterResourceList)}
}

func (s clusterResourceSet) Keys() sets.String {
	return s.set.Keys()
}

func (s clusterResourceSet) List() []*things_test_io_v1.ClusterResource {
	var clusterResourceList []*things_test_io_v1.ClusterResource
	for _, obj := range s.set.List() {
		clusterResourceList = append(clusterResourceList, obj.(*things_test_io_v1.ClusterResource))
	}
	return clusterResourceList
}

func (s clusterResourceSet) Map() map[string]*things_test_io_v1.ClusterResource {
	newMap := map[string]*things_test_io_v1.ClusterResource{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*things_test_io_v1.ClusterResource)
	}
	return newMap
}

func (s clusterResourceSet) Insert(
	clusterResourceList ...*things_test_io_v1.ClusterResource,
) {
	for _, obj := range clusterResourceList {
		s.set.Insert(obj)
	}
}

func (s clusterResourceSet) Has(clusterResource *things_test_io_v1.ClusterResource) bool {
	return s.set.Has(clusterResource)
}

func (s clusterResourceSet) Equal(
	clusterResourceSet ClusterResourceSet,
) bool {
	return s.set.Equal(makeGenericClusterResourceSet(clusterResourceSet.List()))
}

func (s clusterResourceSet) Delete(ClusterResource *things_test_io_v1.ClusterResource) {
	s.set.Delete(ClusterResource)
}

func (s clusterResourceSet) Union(set ClusterResourceSet) ClusterResourceSet {
	return NewClusterResourceSet(append(s.List(), set.List()...)...)
}

func (s clusterResourceSet) Difference(set ClusterResourceSet) ClusterResourceSet {
	newSet := s.set.Difference(makeGenericClusterResourceSet(set.List()))
	return clusterResourceSet{set: newSet}
}

func (s clusterResourceSet) Intersection(set ClusterResourceSet) ClusterResourceSet {
	newSet := s.set.Intersection(makeGenericClusterResourceSet(set.List()))
	var clusterResourceList []*things_test_io_v1.ClusterResource
	for _, obj := range newSet.List() {
		clusterResourceList = append(clusterResourceList, obj.(*things_test_io_v1.ClusterResource))
	}
	return NewClusterResourceSet(clusterResourceList...)
}
