package jbtracer

import (
	"math"
)

type Sphere struct {
}

// NewSphere creates a new Sphere
func NewSphere() *Sphere {
	return &Sphere{}
}

// Intersections returns the intersections of the provided Ray
// with the Sphere
func (s *Sphere) Intersections(r *Ray) []*Intersection {
	sphereToRay := r.origin.Subtract(NewPoint(0, 0, 0))
	a := r.direction.Dot(r.direction)
	b := 2 * r.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	i := make([]*Intersection, 0)
	if discriminant < 0 {
		return i
	}

	discriminantRoot := (float32)(math.Sqrt((float64)(discriminant)))
	t1 := (-1*b - discriminantRoot) / (2 * a)
	t2 := (-1*b + discriminantRoot) / (2 * a)

	var object Object = s
	if t1 < t2 {
		i = append(i, NewIntersection(&object, t1))
		i = append(i, NewIntersection(&object, t2))
	} else {
		i = append(i, NewIntersection(&object, t2))
		i = append(i, NewIntersection(&object, t1))

	}
	return i
}
