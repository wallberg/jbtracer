package jbtracer

import "math"

const (
	Axis_X = iota
	Axis_Y
	Axis_Z
	Degrees180 = 3.1415926536 // pi
	Degrees90  = 1.5707963268 // pi/2
	Degrees60  = 1.0471975512 // pi/3
	Degrees45  = 0.7853981634 // pi/4
	Degrees30  = 0.5235987756 // pi/6
	Degrees10  = 0.1745329252 // pi/18
	Pi         = 3.1415926536 // pi
	Pi2        = 1.5707963268 // pi/2
	Pi3        = 1.0471975512 // pi/3
	Pi4        = 0.7853981634 // pi/4
	Pi6        = 0.5235987756 // pi/6
	Pi18       = 0.1745329252 // pi/18
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
		a.Set(0, 0, cos)
		a.Set(0, 2, sin)
		a.Set(2, 0, -1*sin)
		a.Set(2, 2, cos)
	case Axis_Z:
		a.Set(0, 0, cos)
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

// ViewTransform returns the tranformation matrix that orients the world relative to the
// eye
func ViewTransform(from, to, up *Tuple) *Matrix {
	forward := to.Subtract(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)

	orientation := IdentityMatrix()
	orientation.Set(0, 0, left.X)
	orientation.Set(0, 1, left.Y)
	orientation.Set(0, 2, left.Z)
	orientation.Set(1, 0, trueUp.X)
	orientation.Set(1, 1, trueUp.Y)
	orientation.Set(1, 2, trueUp.Z)
	orientation.Set(2, 0, -1*forward.X)
	orientation.Set(2, 1, -1*forward.Y)
	orientation.Set(2, 2, -1*forward.Z)

	return orientation.Multiply(Translation(-1*from.X, -1*from.Y, -1*from.Z))
}
