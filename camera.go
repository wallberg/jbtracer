package jbtracer

import (
	"log"
	"math"
)

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
	camera := &Camera{
		Hsize:       hsize,
		Vsize:       vsize,
		FieldOfView: fov,
		Transform:   IdentityMatrix(),
	}

	// Compute HalfHeight, HalfWidth, and PixelSize
	halfView := float32(math.Tan(float64(fov) / 2))
	aspect := float32(hsize) / float32(vsize)
	if aspect >= 1 {
		camera.HalfWidth = halfView
		camera.HalfHeight = halfView / aspect
	} else {
		camera.HalfWidth = halfView * aspect
		camera.HalfHeight = halfView
	}
	camera.PixelSize = (camera.HalfWidth * 2) / float32(hsize)

	return camera
}

// RayForPixel computes the world coordinates at the center of the given pixel and
// then returns a Ray that passes through that point.
func (camera *Camera) RayForPixel(x, y int) *Ray {

	// the offset from the edge of the canvas to the pixel's center
	offsetX := (float32(x) + 0.5) * camera.PixelSize
	offsetY := (float32(y) + 0.5) * camera.PixelSize

	// the untransformed coordinates of the pixel in world space.
	// (remember that the camera looks toward -z, so +x is to the *left*.)
	worldX := camera.HalfWidth - offsetX
	worldY := camera.HalfHeight - offsetY

	// using the camera matrix, transform the canvas point and the origin,
	// and then compute the ray's direction vector.
	// (remember that the canvas is at z=-1)
	var inv *Matrix
	var err error
	if inv, err = camera.Transform.Inverse(); err != nil {
		log.Fatal(err)
	}
	pixel := inv.MultiplyTuple(NewPoint(worldX, worldY, -1))
	origin := inv.MultiplyTuple(NewPoint(0, 0, 0))
	direction := pixel.Subtract(origin).Normalize()

	return NewRay(origin, direction)
}
