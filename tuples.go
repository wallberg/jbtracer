package jbtracer

import (
	"fmt"
	"math"
)

// Epsilon is the expected precision for our floating point operations
const Epsilon = 0.00001

type Tuple struct {
	X, Y, Z float32 // 3D coordinates
	W       float32 // 1.0 when a point, 0.0 when a vector
}

// String returns a string representation of the tuple
func (a *Tuple) String() string {
	var types string
	if a.IsPoint() {
		types = "point"
	} else {
		types = "vector"
	}
	return fmt.Sprintf("x=%+2.5f, y=%+2.5f, z=%+2.5f (%s)", a.X, a.Y, a.Z, types)
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

// Negate negates a tuple
func (a *Tuple) Negate() *Tuple {
	return &Tuple{
		X: a.X * -1.0,
		Y: a.Y * -1.0,
		Z: a.Z * -1.0,
		W: a.W * -1.0,
	}
}

// Multiply multiplies a tuple by a scalar
func (a *Tuple) Multiply(scalar float32) *Tuple {
	return &Tuple{
		X: a.X * scalar,
		Y: a.Y * scalar,
		Z: a.Z * scalar,
		W: a.W * scalar,
	}
}

// Divide divides a tuple by a scalar
func (a *Tuple) Divide(scalar float32) *Tuple {
	return &Tuple{
		X: a.X / scalar,
		Y: a.Y / scalar,
		Z: a.Z / scalar,
		W: a.W / scalar,
	}
}

// Magnitude returns the magnitude (or length) of the tuple
func (a *Tuple) Magnitude() float32 {
	x := (float64)(a.X)
	y := (float64)(a.Y)
	z := (float64)(a.Z)
	w := (float64)(a.W)
	return (float32)(math.Sqrt(x*x + y*y + z*z + w*w))
}

// Normalize returns a normalized unit vector
func (a *Tuple) Normalize() *Tuple {
	m := a.Magnitude()
	return &Tuple{
		X: a.X / m,
		Y: a.Y / m,
		Z: a.Z / m,
		W: a.W / m,
	}
}

// Dot returns the dot product of this vector with the provided vector
func (a *Tuple) Dot(b *Tuple) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

// Cross returns the cross product of this vector with the provided vector
func (a *Tuple) Cross(b *Tuple) *Tuple {
	return NewVector(
		a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X,
	)
}

// EqualFloat determines if two float32 values are the within Epsilon of each other
func EqualFloat32(a, b float32) bool {
	return math.Abs((float64)(a)-(float64)(b)) < Epsilon
}

// Reflect reflects vector v around the normal n
func (v *Tuple) Reflect(n *Tuple) *Tuple {
	return v.Subtract(n.Multiply(2 * v.Dot(n)))
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

type Color struct {
	Red, Green, Blue float32
}

var (
	Black *Color = &Color{0, 0, 0}
	White *Color = &Color{1, 1, 1}
)

func NewColor(red, green, blue float32) *Color {
	return &Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

// Equal determines if two Colors are the same
func (a *Color) Equal(b *Color) bool {
	return EqualFloat32(a.Red, b.Red) && EqualFloat32(a.Green, b.Green) && EqualFloat32(a.Blue, b.Blue)
}

// Add adds one Color to another
func (a *Color) Add(b *Color) *Color {
	return &Color{
		Red:   a.Red + b.Red,
		Green: a.Green + b.Green,
		Blue:  a.Blue + b.Blue,
	}
}

// Add subtracts one Color from another
func (a *Color) Subtract(b *Color) *Color {
	return &Color{
		Red:   a.Red - b.Red,
		Green: a.Green - b.Green,
		Blue:  a.Blue - b.Blue,
	}
}

// Multiply multiplies this Color by another Color
func (a *Color) Multiply(b *Color) *Color {
	return &Color{
		Red:   a.Red * b.Red,
		Green: a.Green * b.Green,
		Blue:  a.Blue * b.Blue,
	}
}

// MultiplyScalar multiplies this Color by a scalar
func (a *Color) MultiplyScalar(scalar float32) *Color {
	return &Color{
		Red:   a.Red * scalar,
		Green: a.Green * scalar,
		Blue:  a.Blue * scalar,
	}
}
