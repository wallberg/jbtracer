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
	camera := t.NewCamera(300, 150, t.Pi3)
	camera.Transform = t.ViewTransform(
		t.NewPoint(0, 1.5, -5),
		t.NewPoint(0, 1, 0),
		t.NewVector(0, 1, 0),
	)

	// Add the world objects
	color1 := t.NewColor(1, 1, 0)
	color2 := t.NewColor(0, 0, 1)
	color3 := t.NewColor(1, 0, 0)

	// floor
	floor := t.NewPlane()
	material = t.NewMaterial()
	material.Pattern = t.NewStripePattern(t.Black, t.White)
	material.Specular = 0
	floor.SetMaterial(material)
	world.AddObject(floor)

	// middle sphere
	middle := t.NewSphere()
	middle.SetTransform(t.Translation(-0.5, 1, 0.5))
	material = t.NewMaterial()
	material.Pattern = t.NewCheckersPattern(color1, color2)
	material.Pattern.SetTransform(t.Scaling(0.3, 0.3, 0.3))
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
	material.Pattern = t.NewGradientPattern(color2, color3)
	material.Pattern.SetTransform(
		t.Translation(0.66, 0, 0).Multiply(
			t.Scaling(2, 2, 2),
		),
	)
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
	material.Pattern = t.NewRingPattern(color3, color1)
	material.Pattern.SetTransform(t.Scaling(0.05, 0.05, 0.05))
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
