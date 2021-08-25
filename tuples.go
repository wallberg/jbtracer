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

// IsPoint returns true if this Tuple is a point
func (a *Tuple) IsPoint() bool {
	return math.Abs((float64)(a.W)-1.0) < Epsilon
}

// IsVector returns true if this Tuple is a vector
func (a *Tuple) IsVector() bool {
	return a.W < Epsilon
}

// Equal determines if two Tuples are the same
func (a *Tuple) Equal(b *Tuple) bool {
	return EqualFloat32(a.X, b.X) && EqualFloat32(a.Y, b.Y) && EqualFloat32(a.Z, b.Z) && EqualFloat32(a.W, b.W)
}

// Add adds one tuple to another
func (a *Tuple) Add(b *Tuple) *Tuple {
	return &Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

// Add subtracts one tuple from another
func (a *Tuple) Subtract(b *Tuple) *Tuple {
	return &Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

// Negate negates a single tuple
func (a *Tuple) Negate() *Tuple {
	return &Tuple{
		X: a.X * -1.0,
		Y: a.Y * -1.0,
		Z: a.Z * -1.0,
		W: a.W * -1.0,
	}
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
