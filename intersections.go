package jbtracer

import (
	"fmt"
)

type Intersection struct {
	Object Shape
	T      float64
}

type IntersectionSlice []*Intersection

type PreparedComputations struct {
	T          float64
	Object     Shape
	Point      *Tuple
	EyeV       *Tuple
	NormalV    *Tuple
	Inside     bool
	OverPoint  *Tuple
	ReflectV   *Tuple
	N1         float64
	N2         float64
	UnderPoint *Tuple
}

// NewIntersection creates a new Intersection
func NewIntersection(object Shape, t float64) *Intersection {
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
	} else if !EqualFloat64(a.T, b.T) {
		return false
	}
	return true
}

// PreparedComputations generates a PreparedComputations object for the
// intersection (hit) of the provide Ray with this Intersection
func (hit *Intersection) PreparedComputations(r *Ray, xs IntersectionSlice) *PreparedComputations {

	Point := r.Position(hit.T)

	comps := &PreparedComputations{
		T:       hit.T,
		Object:  hit.Object,
		Point:   Point,
		EyeV:    r.Direction.Multiply(-1),
		NormalV: NormalAt(hit.Object, Point),
	}

	if comps.NormalV.Dot(comps.EyeV) < 0 {
		comps.Inside = true
		comps.NormalV = comps.NormalV.Multiply(-1)
	}

	comps.OverPoint = comps.Point.Add(comps.NormalV.Multiply(Epsilon))
	comps.UnderPoint = comps.Point.Subtract(comps.NormalV.Multiply(Epsilon))
	comps.ReflectV = r.Direction.Reflect(comps.NormalV)

	// Compute N1 and N2
	if xs == nil {
		xs = IntersectionSlice{hit}
	}

	var containers []Shape

	for _, i := range xs {
		if i == hit {
			if len(containers) == 0 {
				comps.N1 = 1.0
			} else {
				comps.N1 = containers[len(containers)-1].Material().RefractiveIndex
			}
		}

		// Look for i in containers
		found := false
		for j := 0; j < len(containers) && !found; j++ {
			if i.Object == containers[j] {
				found = true

				// remove i.object from containers
				copy(containers[j:], containers[j+1:])
				containers[len(containers)-1] = nil
				containers = containers[:len(containers)-1]

			}
		}
		if !found {
			// append i.object onto containers
			containers = append(containers, i.Object)
		}

		if i == hit {
			if len(containers) == 0 {
				comps.N2 = 1.0
			} else {
				comps.N2 = containers[len(containers)-1].Material().RefractiveIndex
			}

			// terminate loop
			return comps
		}
	}

	return comps
}

// String returns a string representation of the Intersection
func (a *IntersectionSlice) String() string {
	return fmt.Sprintf("%+v", *a)
}

// Hit returns the smallest positive intersection from the list
// TODO: determine if we can assume the slice is sorted
func (is IntersectionSlice) Hit() *Intersection {
	var hit *Intersection
	for _, i := range is {
		if i.T > 0 && (hit == nil || i.T < hit.T) {
			hit = i
		}
	}
	return hit
}

// Equal returns true if the two Intersections objects are the same
func (a IntersectionSlice) Equal(b IntersectionSlice) bool {
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
