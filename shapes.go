package jbtracer

import "log"

type Shape interface {
	Intersections(*Ray) Intersections
	NormalAt(worldPoint *Tuple) *Tuple
	Material() *Material
	SetMaterial(*Material)
	Equal(Shape) bool
	Transform() *Matrix
	SetTransform(*Matrix)
}

// IntersectionsCommon provides the functionality common to every Shape which
// implements Intersections(), wrapped around the call to local() which
// is provided by the implementing type.
func IntersectionsCommon(s Shape, r *Ray, local func(r *Ray) Intersections) Intersections {

	// Instead of transforming the Sphere, apply the inverse
	// of the transform to the Ray
	if inv, err := s.Transform().Inverse(); err != nil {
		log.Fatal(err)
	} else {
		r = r.Transform(inv)
	}

	return local(r)
}

// NormalAtCommon provides the functionality common to every Shape which
// implements NormalAt(), wrapped around the call to locat() which is provided
// by the implementing type.
func NormalAtCommon(s Shape, worldPoint *Tuple, local func(objectPoint *Tuple) (objectNormal *Tuple)) *Tuple {
	inv, err := s.Transform().Inverse()
	if err != nil {
		log.Fatalf("Matrix s.Transform()=%v is not invertible", s.Transform())
	}

	objectPoint := inv.MultiplyTuple(worldPoint)
	objectNormal := local(objectPoint)
	worldNormal := inv.Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
