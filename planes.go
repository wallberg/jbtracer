package jbtracer

import (
	"fmt"
	"math"
)

type Plane struct {
	transform *Matrix
	material  *Material
}

// NewPlane creates a new Plane
func NewPlane() *Plane {
	return &Plane{
		transform: IdentityMatrix(),
		material:  NewMaterial(),
	}
}

// String returns a string representation of the Plane
func (a *Plane) String() string {
	return fmt.Sprintf("%+v", *a)
}

// Transform returns the transform for this Plane
func (a *Plane) Transform() *Matrix {
	return a.transform
}

// SetTransform sets the transform for this Plane
func (a *Plane) SetTransform(transform *Matrix) {
	a.transform = transform
}

// Equal returns whether the two Planes are the same
func (a *Plane) Equal(b Shape) bool {
	if sb, ok := b.(*Plane); !ok {
		return false
	} else {
		return a != nil && sb != nil && a.transform.Equal(sb.transform) && a.material.Equal(sb.material)
	}
}

// Material returns the material for this Plane
func (a *Plane) Material() *Material {
	return a.material
}

// SetMaterial sets the material for this Plane
func (a *Plane) SetMaterial(material *Material) {
	a.material = material
}

// Intersections returns the intersections of the provided Ray
// with the Plane
func (plane *Plane) Intersections(r *Ray) IntersectionSlice {

	if math.Abs(r.Direction.Y) < Epsilon {
		return nil
	}

	t := -1 * r.Origin.Y / r.Direction.Y
	return IntersectionSlice{NewIntersection(plane, t)}
}

// NormalAt returns the surface normal to the Plane at the given Point.
func (plane *Plane) NormalAt(objectPoint *Tuple) *Tuple {

	return NewVector(0, 1, 0)
}
