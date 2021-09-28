package jbtracer

import "log"

type Shape interface {
	Intersections(*Ray) IntersectionSlice
	NormalAt(worldPoint *Tuple) *Tuple
	Material() *Material
	SetMaterial(*Material)
	Equal(Shape) bool
	Transform() *Matrix
	SetTransform(*Matrix)
}

// Intersections provides the common functionality wrapped around every
// Shape.Intersections() implemention.
func Intersections(s Shape, r *Ray) IntersectionSlice {

	// Instead of transforming the Sphere, apply the inverse
	// of the transform to the Ray
	if inv, err := s.Transform().Inverse(); err != nil {
		log.Fatal(err)
	} else {
		r = r.Transform(inv)
	}

	return s.Intersections(r)
}

// NormalAt provides the common functionality wrapped around every
// Shape.NormalAt() implemention.
func NormalAt(s Shape, worldPoint *Tuple) *Tuple {
	inv, err := s.Transform().Inverse()
	if err != nil {
		log.Fatalf("Matrix s.Transform()=%v is not invertible", s.Transform())
	}

	objectPoint := inv.MultiplyTuple(worldPoint)
	objectNormal := s.NormalAt(objectPoint)
	worldNormal := inv.Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
