package main

import (
	"fmt"
	"math"

	t "github.com/wallberg/jbtracer"
)

func main() {

	size := 500
	c := t.NewCanvas(size, size)

	radius := (float32)(size) * 4 / 5 / 2
	twelve := t.NewPoint(0, 0, 1)
	shift := t.NewPoint((float32)(size)/2, 0, (float32)(size)/2)
	color := &t.Color{Red: 1, Green: 1, Blue: 0}

	// Iterate over 12 points of the clock face
	for i := 0; i < 12; i++ {

		// Rotate around the y-axis, multiply by the radius,
		// then translate into second quadrant
		r := t.Rotation(t.Axis_Y, math.Pi/6*(float32)(i))
		dot := r.MultiplyTuple(twelve).Multiply(radius).Add(shift)

		x := int(dot.X)
		y := size - int(dot.Z)

		if x >= 0 && x < size && y >= 0 && y < size {
			c.Grid[x][y] = color
			c.Grid[x+1][y] = color
			c.Grid[x][y+1] = color
			c.Grid[x-1][y] = color
			c.Grid[x][y-1] = color
		}
	}

	// Output the PPM image file
	ppm := c.NewPPM()
	for _, line := range *ppm {
		fmt.Print(line)
	}
}
