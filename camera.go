package jbtracer

import "math"

type Camera struct {
	Hsize       int
	Vsize       int
	FieldOfView float32
	HalfHeight  float32
	HalfWidth   float32
	PixelSize   float32
	Transform   *Matrix
}

// NewCamera returns a new Camera, with default Transform
func NewCamera(hsize, vsize int, fov float32) *Camera {
	c := &Camera{
		Hsize:       hsize,
		Vsize:       vsize,
		FieldOfView: fov,
		Transform:   IdentityMatrix(),
	}

	// Compute HalfHeight, HalfWidth, and PixelSize
	halfView := float32(math.Tan(float64(fov) / 2))
	aspect := float32(hsize) / float32(vsize)
	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2) / float32(hsize)

	return c
}
