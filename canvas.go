package jbtracer

import (
	"fmt"
	"math"
	"strings"
)

type Canvas struct {
	Width, Height int
	Grid          [][]*Color
}

type PPM []string

// NewCanvas returns a new canvas of width, height with an
// initialized grid of pixels
func NewCanvas(width, height int) *Canvas {
	c := &Canvas{Width: width, Height: height}

	// Initialize the grid of pixels
	c.Grid = make([][]*Color, width)
	for x := range c.Grid {
		c.Grid[x] = make([]*Color, height)
		for y := range c.Grid[x] {
			c.Grid[x][y] = &Color{0, 0, 0}
		}
	}

	return c
}

// SetColorAll sets every pixel to the provided color
func (c *Canvas) SetColorAll(color *Color) {
	for x, column := range c.Grid {
		for y := range column {
			c.Grid[x][y] = color
		}
	}
}

// NewPPM creates a new PPM image from this Canvas
func (c *Canvas) NewPPM() *PPM {
	ppm := make(PPM, 3)

	// Add the header
	ppm[0] = "P3\n"
	ppm[1] = fmt.Sprintf("%d %d\n", c.Width, c.Height)
	ppm[2] = "255\n"

	// Add the grid of pixels
	var line strings.Builder

	// Iterate over the rows
	for y := 0; y < c.Height; y++ {
		// Iterate over the columns
		for x := 0; x < c.Width; x++ {
			color := c.Grid[x][y]
			// Iteratve over RGB values of this pixel
			for _, valueFloat32 := range []float32{color.Red, color.Green, color.Blue} {

				// Scale the color value
				valueInt := int(math.Round((float64)(valueFloat32 * 255)))

				// Clamp to [0,255]
				if valueInt < 0 {
					valueInt = 0
				} else if valueInt > 255 {
					valueInt = 255
				}

				// Convert to a string
				valueString := fmt.Sprintf("%d", valueInt)

				// Check if this color value will fit on this line
				if line.Len()+len(valueString)+1 > 70 {
					// Add this line to PPM
					line.WriteString("\n")
					ppm = append(ppm, line.String())

					// Start a new line
					line.Reset()
				}

				// Add this color value to the line
				if line.Len() != 0 {
					line.WriteString(" ")
				}
				line.WriteString(valueString)
			}
		}

		// Always end a line when we reach the end of the row
		line.WriteString("\n")
		ppm = append(ppm, line.String())
		line.Reset()
	}

	// Add the last line to PPM
	line.WriteString("\n")
	ppm = append(ppm, line.String())

	return &ppm
}
