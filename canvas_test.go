package jbtracer

import (
	"fmt"
	"strings"
)

func canvas(width, height int) error {
	c = NewCanvas(width, height)
	return nil
}

func equalsCanvasField(field string, expected int) error {
	var got int
	switch field {
	case "width":
		got = c.Width
	case "height":
		got = c.Height
	}

	if got != expected {
		return fmt.Errorf("Expected c.%s = %d; got %d", field, expected, got)
	}
	return nil
}

func allCanvasColors(red, green, blue float64) error {

	expected := &Color{red, green, blue}

	for x, column := range c.Grid {
		for y, got := range column {
			if !got.Equal(expected) {
				return fmt.Errorf("Expected pixel(%d,%d) to be %v; got %v", x, y, expected, c.Grid[x][y])
			}
		}
	}
	return nil
}

func writePixel(x, y int, c1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", c1name)
	}
	c.Grid[x][y] = c1
	return nil
}

func pixelAt(x, y int, c1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", c1name)
	}
	expected := c1
	got := c.Grid[x][y]
	if !got.Equal(expected) {
		return fmt.Errorf("Expected pixel(%d,%d) = %v; got %v", x, y, expected, got)
	}
	return nil
}

func canvasToPPM() error {
	ppm = c.NewPPM()
	return nil
}

func linesOfPPM(start, stop int, expected string) error {
	// Build the lines from ppm
	var b strings.Builder
	for _, line := range (*ppm)[start-1 : stop] {
		b.WriteString(line)
	}
	got := b.String()
	if got != expected {
		return fmt.Errorf("Expected lines %d-%d:\n%sgot\n%s", start, stop, expected, got)
	}
	return nil
}

func assignCanvasAllColors(red, green, blue float64) error {
	c.SetColorAll(&Color{red, green, blue})
	return nil
}

func ppmEndsWithANewlineCharacter() error {
	line := (*ppm)[len(*ppm)-1]
	got := line[len(line)-1]
	if got != '\n' {
		return fmt.Errorf("Expected ppm to end with a newline character; got %c", got)
	}
	return nil
}
