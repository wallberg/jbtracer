package jbtracer

type Shape interface {
	Intersections(*Ray) Intersections
	NormalAt(worldPoint *Tuple) *Tuple
	Material() *Material
	Equal(Shape) bool
	Transform() *Matrix
}
