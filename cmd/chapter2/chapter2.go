package main

import (
	"fmt"

	"github.com/wallberg/jbtracer"
)

type projectile struct {
	position, velocity *jbtracer.Tuple
}

type environment struct {
	gravity, wind *jbtracer.Tuple
}

func main() {
	p := &projectile{
		position: jbtracer.NewVector(0, 1, 0),
		velocity: jbtracer.NewPoint(1, 1.8, 0).Normalize().Multiply(11.25),
	}

	e := &environment{
		gravity: jbtracer.NewVector(0, -0.1, 0),
		wind:    jbtracer.NewVector(-0.01, 0, 0),
	}

	width := 900
	height := 550
	c := jbtracer.NewCanvas(width, height)

	color := &jbtracer.Color{Red: 1, Green: 1, Blue: 0}

	for t := 0; p.position.Y > 0; t++ {

		x := int(p.position.X)
		y := height - int(p.position.Y)

		if x >= 0 && x < width && y >= 0 && y < height {
			c.Grid[x][y] = color
		}

		tick(p, e)
	}

	ppm := c.NewPPM()
	for _, line := range *ppm {
		fmt.Print(line)
	}
}

func tick(p *projectile, e *environment) {
	p.position = p.position.Add(p.velocity)
	p.velocity = p.velocity.Add(e.gravity).Add(e.wind)
}
