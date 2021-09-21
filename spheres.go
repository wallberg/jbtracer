package jbtracer

import (
	"fmt"
	"log"
	"math"
)

type Sphere struct {
	Transform *Matrix
	material  *Material
}

// NewSphere creates a new Sphere
func NewSphere() *Sphere {
	return &Sphere{
		Transform: IdentityMatrix(),
		material:  NewMaterial(),
	}
}

// String returns a string representation of the Sphere
func (a *Sphere) String() string {
	return fmt.Sprintf("%+v", *a)
}

// Equal returns whether the two Spheres are the same
func (a *Sphere) Equal(b Object) bool {
	if sb, ok := b.(*Sphere); !ok {
		return false
	} else {
		return a != nil && sb != nil && a.Transform.Equal(sb.Transform) && a.material.Equal(sb.material)
	}
}

// Material returns the material for this Sphere
func (a *Sphere) Material() *Material {
	return a.material
}

// SetMaterial sets the material for this Sphere
func (a *Sphere) SetMaterial(material *Material) {
	a.material = material
}

// Intersections returns the intersections of the provided Ray
// with the Sphere
func (s *Sphere) Intersections(r *Ray) []*Intersection {

	// Instead of transforming the Sphere, apply the inverse
	// of the transform to the Ray
	if inv, err := s.Transform.Inverse(); err != nil {
		log.Fatal(err)
	} else {
		r = r.Transform(inv)
	}

	sphereToRay := r.Origin.Subtract(NewPoint(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	i := make([]*Intersection, 0)
	if discriminant < 0 {
		return i
	}

	discriminantRoot := (float32)(math.Sqrt((float64)(discriminant)))
	t1 := (-1*b - discriminantRoot) / (2 * a)
	t2 := (-1*b + discriminantRoot) / (2 * a)

	if t1 < t2 {
		i = append(i, NewIntersection(s, t1))
		i = append(i, NewIntersection(s, t2))
	} else {
		i = append(i, NewIntersection(s, t2))
		i = append(i, NewIntersection(s, t1))

	}
	return i
}

// NormalAt returns the surface normal to the sphere at the given Point.
func (s *Sphere) NormalAt(worldPoint *Tuple) *Tuple {
	transformInverse, err := s.Transform.Inverse()
	if err != nil {
		log.Fatalf("Matrix s.Transform=%v is not invertible", s.Transform)
	}

	objectPoint := transformInverse.MultiplyTuple(worldPoint)
	objectNormal := objectPoint.Subtract(NewPoint(0, 0, 0))
	worldNormal := transformInverse.Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}
