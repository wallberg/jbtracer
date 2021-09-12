package jbtracer

import "math"

const (
	Axis_X = iota
	Axis_Y
	Axis_Z
)

// Translation returns a translation matrix for vector(x, y, z)
func Translation(x, y, z float64) *Matrix {
	a := IdentityMatrix()
	a.Set(0, 3, x)
	a.Set(1, 3, y)
	a.Set(2, 3, z)
	return a
}

// Scaling returns a scaling matrix for vector(x, y, z)
func Scaling(x, y, z float64) *Matrix {
	a := IdentityMatrix()
	a.Set(0, 0, x)
	a.Set(1, 1, y)
	a.Set(2, 2, z)
	return a
}

// Rotation returns the rotation matrix around the x, y, or z axis by
// the provided radians
func Rotation(axis int, radians float64) *Matrix {
	a := IdentityMatrix()
	sin := math.Sin(radians)
	cos := math.Cos(radians)

	switch axis {
	case Axis_X:
		a.Set(1, 1, cos)
		a.Set(1, 2, -1*sin)
		a.Set(2, 1, sin)
		a.Set(2, 2, cos)
	case Axis_Y:
		a.Set(0, 1, cos)
		a.Set(0, 2, sin)
		a.Set(2, 0, -1*sin)
		a.Set(2, 2, cos)
	case Axis_Z:
		a.Set(0, 1, cos)
		a.Set(0, 1, -1*sin)
		a.Set(1, 0, sin)
		a.Set(1, 1, cos)
	}
	return a
}

// Shearing returns the shearing matrix
func Shearing(xY, xZ, yX, yZ, zX, zY float64) *Matrix {
	a := IdentityMatrix()
	a.Set(0, 1, xY)
	a.Set(0, 2, xZ)
	a.Set(1, 0, yX)
	a.Set(1, 2, yZ)
	a.Set(2, 0, zX)
	a.Set(2, 1, zY)
	return a
}
