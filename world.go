package jbtracer

import "sort"

type World struct {
	Light   *PointLight
	Objects []Shape
}

// NewWorld returns a new, empty World
func NewWorld() *World {
	w := &World{}
	w.Objects = make([]Shape, 0)
	return w
}

// DefaultWorld returns a new, default populated World
func DefaultWorld() *World {
	w := NewWorld()
	w.Light = NewPointLight(White, NewPoint(-10, 10, -10))

	s := NewSphere()
	s.material.Color = NewColor(0.8, 1.0, 0.6)
	s.material.Diffuse = 0.7
	s.material.Specular = 0.2
	w.AddObject(s)

	s = NewSphere()
	s.SetTransform(Scaling(0.5, 0.5, 0.5))
	w.AddObject(s)

	return w
}

// AddObject adds a new Object to the World
func (w *World) AddObject(object Shape) {
	w.Objects = append(w.Objects, object)
}

// Intersections returns intersections of the Ray with every object
// in the World, sorted in ascending Ray.T order
func (w *World) Intersections(r *Ray) IntersectionSlice {

	// Accumulate all of the intersections
	is := make([]*Intersection, 0)
	for _, object := range w.Objects {
		is = append(is, Intersections(object, r)...)
	}

	// Sort by ascending T value
	sort.SliceStable(is, func(i, j int) bool {
		return is[i].T < is[j].T
	})

	return is
}

// ShadeHit returns the Color at the Intersection encapsulated by a PreparedComputations
// in the given World
func (w *World) ShadeHit(comps *PreparedComputations) *Color {

	shadowed := w.IsShadowed(comps.OverPoint)

	return comps.Object.Material().Lighting(
		w.Light,
		comps.Object,
		comps.OverPoint,
		comps.EyeV,
		comps.NormalV,
		shadowed,
	)
}

// ColorAt returns the Color at the Point where the provided ray
// intersects this World
func (w *World) ColorAt(r *Ray) *Color {

	// Get the intersections with the World
	xs := w.Intersections(r)

	// Get the hit
	if hit := xs.Hit(); hit == nil {
		return Black
	} else {
		// Return the Color at the intersection
		comps := hit.PreparedComputations(r)
		return w.ShadeHit(comps)
	}
}

// IsShadowed determines if a Point in this World is in shadow
func (w *World) IsShadowed(point *Tuple) bool {
	v := w.Light.Position.Subtract(point)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := NewRay(point, direction)
	xs := w.Intersections(r)

	if hit := xs.Hit(); hit != nil && hit.T < distance {
		return true
	} else {
		return false
	}
}
