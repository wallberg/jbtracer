package jbtracer

import (
	"fmt"
	"math"
)

type Sphere struct {
	transform *Matrix
	material  *Material
}

// NewSphere creates a new Sphere
func NewSphere() *Sphere {
	return &Sphere{
		transform: IdentityMatrix(),
		material:  NewMaterial(),
	}
}

// NewGlassSphere creates a new glass Sphere
func NewGlassSphere() *Sphere {
	sphere := NewSphere()
	sphere.material.Transparency = 1.0
	sphere.material.RefractiveIndex = 1.5
	return sphere
}

// String returns a string representation of the Sphere
func (a *Sphere) String() string {
	return fmt.Sprintf("%+v", *a)
}

// Transform returns the transform for this Sphere
func (a *Sphere) Transform() *Matrix {
	return a.transform
}

// SetMaterial sets the material for this Sphere
func (a *Sphere) SetTransform(transform *Matrix) {
	a.transform = transform
}

// Equal returns whether the two Spheres are the same
func (a *Sphere) Equal(b Shape) bool {
	if sb, ok := b.(*Sphere); !ok {
		return false
	} else {
		return a != nil && sb != nil && a.transform.Equal(sb.transform) && a.material.Equal(sb.material)
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
func (s *Sphere) Intersections(r *Ray) IntersectionSlice {

	sphereToRay := r.Origin.Subtract(NewPoint(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	i := make([]*Intersection, 0)
	if discriminant < 0 {
		return i
	}

	discriminantRoot := math.Sqrt(discriminant)
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
func (s *Sphere) NormalAt(objectPoint *Tuple) *Tuple {

	return objectPoint.Subtract(NewPoint(0, 0, 0))
}
