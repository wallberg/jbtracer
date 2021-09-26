package jbtracer

import "fmt"

type Ray struct {
	Origin, Direction *Tuple
}

func NewRay(origin, direction *Tuple) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (a *Ray) String() string {
	return fmt.Sprintf("origin=%v, direction=%v", a.Origin, a.Direction)
}

// Position returns the Point at time t along the Ray
func (a *Ray) Position(t float64) *Tuple {
	return a.Direction.Multiply(t).Add(a.Origin)
}

// Transform transforms the Ray by the provided Matrix
func (a *Ray) Transform(transform *Matrix) *Ray {
	return NewRay(
		transform.MultiplyTuple(a.Origin),
		transform.MultiplyTuple(a.Direction),
	)
}
