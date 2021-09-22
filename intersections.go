package jbtracer

import (
	"fmt"
)

type Intersection struct {
	Object Object
	T      float32
}

type Intersections []*Intersection

type PreparedComputations struct {
	T       float32
	Object  Object
	Point   *Tuple
	EyeV    *Tuple
	NormalV *Tuple
}

// NewIntersection creates a new Intersection
func NewIntersection(object Object, t float32) *Intersection {
	return &Intersection{
		Object: object,
		T:      t,
	}
}

// String returns a string representation of the Intersection
func (a *Intersection) String() string {
	return fmt.Sprintf("%+v", *a)
}

// Equal returns true if the two Intersection objects are the same
func (a *Intersection) Equal(b *Intersection) bool {
	if a == nil || b == nil {
		return false
	} else if a.Object != b.Object {
		return false
	} else if !EqualFloat32(a.T, b.T) {
		return false
	}
	return true
}

// PreparedComputations generates a PreparedComputations object for the intersection
// of the provide Ray with this Intersection
func (i *Intersection) PreparedComputations(r *Ray) *PreparedComputations {

	Point := r.Position(i.T)

	return &PreparedComputations{
		T:       i.T,
		Object:  i.Object,
		Point:   Point,
		EyeV:    r.Direction.Multiply(-1),
		NormalV: i.Object.NormalAt(Point),
	}
}

// String returns a string representation of the Intersection
func (a *Intersections) String() string {
	return fmt.Sprintf("%+v", *a)
}

// Hit returns the smallest positive intersection from the list
func (is Intersections) Hit() *Intersection {
	var hit *Intersection
	for _, i := range is {
		if i.T > 0 && (hit == nil || i.T < hit.T) {
			hit = i
		}
	}
	return hit
}

// Equal returns true if the two Intersections objects are the same
func (a Intersections) Equal(b Intersections) bool {
	if a == nil || b == nil {
		return false
	} else if len(a) != len(b) {
		return false
	}
	for i, aI := range a {
		if !aI.Equal((b)[i]) {
			return false
		}
	}
	return true
}
