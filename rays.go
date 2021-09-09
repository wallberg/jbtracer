package jbtracer

import "fmt"

type Ray struct {
	origin, direction *Tuple
}

func NewRay(origin, direction *Tuple) *Ray {
	return &Ray{
		origin:    origin,
		direction: direction,
	}
}

func (a *Ray) String() string {
	return fmt.Sprintf("origin=%v, direction=%v", a.origin, a.direction)
}

func (a *Ray) Position(t float32) *Tuple {
	return a.direction.Multiply(t).Add(a.origin)
}

// Transform transforms the Ray by the provided Matrix
func (a *Ray) Transform(transform *Matrix) *Ray {
	return NewRay(
		transform.MultiplyTuple(a.origin),
		transform.MultiplyTuple(a.direction),
	)
}
