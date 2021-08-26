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
		position: jbtracer.NewPoint(0, 1, 0),
		velocity: jbtracer.NewVector(1, 1, 0),
	}

	e := &environment{
		gravity: jbtracer.NewVector(0, -0.1, 0),
		wind:    jbtracer.NewVector(-0.01, 0, 0),
	}

	for t := 0; p.position.Y > 0; t++ {
		tick(p, e)
		fmt.Printf("t=%d; position=%v\n", t, p.position)
	}
}

func tick(p *projectile, e *environment) {
	p.position = p.position.Add(p.velocity)
	p.velocity = p.velocity.Add(e.gravity).Add(e.wind)
}
