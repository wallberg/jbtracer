package jbtracer

import (
	"log"
	"math"
)

type Camera struct {
	Hsize       int
	Vsize       int
	FieldOfView float64
	HalfHeight  float64
	HalfWidth   float64
	PixelSize   float64
	Transform   *Matrix
}

// NewCamera returns a new Camera, with default Transform
func NewCamera(hsize, vsize int, fov float64) *Camera {
	camera := &Camera{
		Hsize:       hsize,
		Vsize:       vsize,
		FieldOfView: fov,
		Transform:   IdentityMatrix(),
	}

	// Compute HalfHeight, HalfWidth, and PixelSize
	halfView := math.Tan(fov / 2)
	aspect := float64(hsize) / float64(vsize)
	if aspect >= 1 {
		camera.HalfWidth = halfView
		camera.HalfHeight = halfView / aspect
	} else {
		camera.HalfWidth = halfView * aspect
		camera.HalfHeight = halfView
	}
	camera.PixelSize = (camera.HalfWidth * 2) / float64(hsize)

	return camera
}

// RayForPixel computes the world coordinates at the center of the given pixel and
// then returns a Ray that passes through that point.
func (camera *Camera) RayForPixel(x, y int) *Ray {

	// the offset from the edge of the canvas to the pixel's center
	offsetX := (float64(x) + 0.5) * camera.PixelSize
	offsetY := (float64(y) + 0.5) * camera.PixelSize

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

// Render uses this Camera to render an image of the provided World
func (camera *Camera) Render(world *World) *Canvas {

	image := NewCanvas(camera.Hsize, camera.Vsize)

	for y := 0; y < camera.Vsize; y++ {
		for x := 0; x < camera.Hsize; x++ {
			ray := camera.RayForPixel(x, y)
			color := world.ColorAt(ray)
			image.Grid[x][y] = color
		}
	}

	return image
}
