package main

import (
	"fmt"

	t "github.com/wallberg/jbtracer"
)

func main() {

	var material *t.Material

	// Configure the world
	world := t.NewWorld()
	world.Light = t.NewPointLight(t.White, t.NewPoint(-10, 10, -10))

	// Configure the camera
	camera := t.NewCamera(3000, 1500, t.Pi3)
	camera.Transform = t.ViewTransform(
		t.NewPoint(0, 1.5, -5),
		t.NewPoint(0, 1, 0),
		t.NewVector(0, 1, 0),
	)

	// Add the world objects

	// floor
	floor := t.NewSphere()
	floor.Transform = t.Scaling(10, 0.01, 10)
	material = t.NewMaterial()
	material.Color = t.NewColor(1, 0.9, 0.9)
	material.Specular = 0
	floor.SetMaterial(material)
	world.AddObject(floor)

	// left wall
	leftWall := t.NewSphere()
	leftWall.Transform =
		t.Translation(0, 0, 5).Multiply(
			t.Rotation(t.Axis_Y, -1*t.Pi4),
		).Multiply(
			t.Rotation(t.Axis_X, t.Pi2),
		).Multiply(
			t.Scaling(10, 0.01, 10),
		)
	leftWall.SetMaterial(floor.Material())
	world.AddObject(leftWall)

	// right wall
	rightWall := t.NewSphere()
	rightWall.Transform =
		t.Translation(0, 0, 5).Multiply(
			t.Rotation(t.Axis_Y, t.Pi4),
		).Multiply(
			t.Rotation(t.Axis_X, t.Pi2),
		).Multiply(
			t.Scaling(10, 0.01, 10),
		)
	rightWall.SetMaterial(floor.Material())
	world.AddObject(rightWall)

	// middle sphere
	middle := t.NewSphere()
	middle.Transform = t.Translation(-0.5, 1, 0.5)
	material = t.NewMaterial()
	material.Color = t.NewColor(0.1, 1, 0.5)
	material.Diffuse = 0.7
	material.Specular = 0.3
	middle.SetMaterial(material)
	world.AddObject(middle)

	// right sphere
	right := t.NewSphere()
	right.Transform = t.Translation(1.5, 0.5, -0.5).Multiply(
		t.Scaling(0.5, 0.5, 0.5),
	)
	material = t.NewMaterial()
	material.Color = t.NewColor(0.5, 1, 0.1)
	material.Diffuse = 0.7
	material.Specular = 0.3
	right.SetMaterial(material)
	world.AddObject(right)

	// left sphere
	left := t.NewSphere()
	left.Transform = t.Translation(-1.5, 0.33, -0.75).Multiply(
		t.Scaling(0.33, 0.33, 0.33),
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
