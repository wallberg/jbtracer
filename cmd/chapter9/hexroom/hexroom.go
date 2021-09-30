package main

import (
	"fmt"
	"math"

	t "github.com/wallberg/jbtracer"
)

func main() {

	var material *t.Material

	// Configure the world
	world := t.NewWorld()
	world.Light = t.NewPointLight(t.White, t.NewPoint(-10, 10, -10))

	// Configure the camera
	camera := t.NewCamera(200, 100, t.Pi/1.5)
	camera.Transform = t.ViewTransform(
		t.NewPoint(0, 3, -10),
		t.NewPoint(0, -1, 0),
		t.NewVector(0, 1, 1),
	)

	// Add the world objects

	// floor
	floor := t.NewPlane()
	material = t.NewMaterial()
	// material.Color = t.NewColor(1, 0.9, 0.9)
	material.Specular = 0
	floor.SetMaterial(material)
	world.AddObject(floor)

	// walls
	const wallRadius = 30
	for i := 0; i < 6; i++ {
		wall := t.NewPlane()
		angle := math.Pi / 3 * float64(i)

		wall.SetTransform(
			t.Rotation(t.Axis_Y, angle).Multiply(
				t.Translation(0, 0, wallRadius),
			).Multiply(
				t.Rotation(t.Axis_X, t.Degrees90),
			),
		)
		material = t.NewMaterial()
		c := float64(i) / 6
		material.Color = t.NewColor(c, c, c)
		wall.SetMaterial(material)

		world.AddObject(wall)
	}

	// middle sphere
	middle := t.NewSphere()
	middle.SetTransform(t.Translation(-0.5, 1, 0.5))
	material = t.NewMaterial()
	material.Color = t.NewColor(0.1, 1, 0.5)
	material.Diffuse = 0.7
	material.Specular = 0.3
	middle.SetMaterial(material)
	world.AddObject(middle)

	// right sphere
	right := t.NewSphere()
	right.SetTransform(
		t.Translation(1.5, 0.5, -0.5).Multiply(
			t.Scaling(0.5, 0.5, 0.5),
		),
	)
	material = t.NewMaterial()
	material.Color = t.NewColor(0.5, 1, 0.1)
	material.Diffuse = 0.7
	material.Specular = 0.3
	right.SetMaterial(material)
	world.AddObject(right)

	// left sphere
	left := t.NewSphere()
	left.SetTransform(
		t.Translation(-1.5, 0.33, -0.75).Multiply(
			t.Scaling(0.33, 0.33, 0.33)),
	)
	material = t.NewMaterial()
	material.Color = t.NewColor(1, 0.8, 0.1)
	material.Diffuse = 0.7
	material.Specular = 0.3
	left.SetMaterial(material)
	world.AddObject(left)

	// Render the result to a canvas
	canvas := camera.Render(world)

	// Output the PPM image file
	ppm := canvas.NewPPM()
	for _, line := range *ppm {
		fmt.Print(line)
	}
}
