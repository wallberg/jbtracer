package jbtracer

type Object interface {
	Intersections(r *Ray) []*Intersection
	NormalAt(worldPoint *Tuple) *Tuple
	Material() *Material
	Equal(b Object) bool
}
