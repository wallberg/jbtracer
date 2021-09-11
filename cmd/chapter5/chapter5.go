package main

import (
	"fmt"

	t "github.com/wallberg/jbtracer"
)

func main() {

	// world settings
	rayOrigin := t.NewPoint(0, 0, -5)
	var wallZ float32 = 10
	var wallSize float32 = 7
	wallZHalf := wallSize / 2
	color := &t.Color{Red: 1, Green: 0, Blue: 0}

	// canvas settings
	canvasPixels := 500
	pixelSize := wallSize / float32(canvasPixels)
	c := t.NewCanvas(canvasPixels, canvasPixels)

	sphere := t.NewSphere()
	transform := t.Rotation(t.Axis_Y, 0.78539)
	transform = transform.Multiply(t.Scaling(0.4, 1, 1))
	transform = transform.Multiply(t.Translation(0.4, 0, 0))
	sphere.Transform = transform

	// Iterate over canvas points
	for y := 0; y < canvasPixels; y++ {
		worldY := wallZHalf - pixelSize*float32(y)

		for x := 0; x < canvasPixels; x++ {
			worldX := -1*wallZHalf + pixelSize*float32(x)

			// Create a ray from the light source to the canvas point
			position := t.NewPoint(worldX, worldY, wallZ)
			vector := position.Subtract(rayOrigin).Normalize()
			ray := t.NewRay(rayOrigin, vector)

			// Determine if the ray intersects the sphere
			var xs t.Intersections = sphere.Intersections(ray)
			if hit := xs.Hit(); hit != nil {
				c.Grid[x][y] = color
			}
		}
	}

	// Output the PPM image file
	ppm := c.NewPPM()
	for _, line := range *ppm {
		fmt.Print(line)
	}
}
