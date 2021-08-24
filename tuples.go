package jbtracer

import (
	"math"
)

// Epsilon is the expected precision for our floating point operations
const Epsilon = 0.00001

type Tuple struct {
	X, Y, Z float32 // 3D coordinates
	W       float32 // 1.0 when a point, 0.0 when a vector
}

// Equal determines if two Tuples are the same
func (a *Tuple) Equal(b *Tuple) bool {
	return EqualFloat32(a.X, b.X) && EqualFloat32(a.Y, b.Y) && EqualFloat32(a.Z, b.Z) && EqualFloat32(a.W, b.W)
}

// EqualFloat determines if two float32 values are the within Epsilon of each other
func EqualFloat32(a, b float32) bool {
	return math.Abs((float64)(a)-(float64)(b)) < Epsilon
}

// NewPoint creates a new Tuple of type point
func NewPoint(X, Y, Z float32) *Tuple {
	point := &Tuple{
		X: X,
		Y: Y,
		Z: Z,
		W: 1.0,
	}

	return point
}

// NewVector creates a new Tuple of type vector
func NewVector(X, Y, Z float32) *Tuple {
	vector := &Tuple{
		X: X,
		Y: Y,
		Z: Z,
		W: 0.0,
	}

	return vector
}
