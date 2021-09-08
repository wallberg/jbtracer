package jbtracer

type Object interface {
	Intersections(r *Ray) []*Intersection
}
