package jbtracer

type Object interface {
	Intersections(r *Ray) []*Intersection
}

type Intersection struct {
	object *Object
	t      float32
}

// NewIntersection creates a new Intersection
func NewIntersection(object *Object, t float32) *Intersection {
	return &Intersection{
		object: object,
		t:      t,
	}
}
