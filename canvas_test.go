package jbtracer

import "fmt"

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

func allCanvasColors(red, green, blue float32) error {

	expected := &Color{red, green, blue}

	for x := range c.Grid {
		for y := range c.Grid[x] {
			if c.Grid[x][y].Equal(expected) {
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
