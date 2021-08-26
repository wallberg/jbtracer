package jbtracer

type Canvas struct {
	Width, Height int
	Grid          [][]*Color
}

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
