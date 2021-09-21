package jbtracer

import "sort"

type World struct {
	Light   *PointLight
	Objects []Object
}

// NewWorld returns a new, empty World
func NewWorld() *World {
	w := &World{}
	w.Objects = make([]Object, 0)
	return w
}

// DefaultWorld returns a new, default populated World
func DefaultWorld() *World {
	w := NewWorld()
	w.Light = NewPointLight(White, NewPoint(-10, 10, -10))

	s := NewSphere()
	s.Transform = Scaling(0.5, 0.5, 0.5)
	w.AddObject(s)

	s = NewSphere()
	s.material.Color = NewColor(0.8, 1.0, 0.6)
	s.material.Diffuse = 0.7
	s.material.Specular = 0.2
	w.AddObject(s)

	return w
}

// AddObject adds a new Object to the World
func (w *World) AddObject(object Object) {
	w.Objects = append(w.Objects, object)
}

// Intersections returns intersections of the Ray with every object
// in the World, sorted in ascending Ray.T order
func (w *World) Intersections(r *Ray) []*Intersection {

	// Accumulate all of the intersections
	is := make([]*Intersection, 0)
	for _, object := range w.Objects {
		is = append(is, object.Intersections(r)...)
	}

	// Sort by ascending T value
	sort.SliceStable(is, func(i, j int) bool {
		return is[i].T < is[j].T
	})

	return is
}
