package jbtracer

type PointLight struct {
	Intensity *Color
	Position  *Tuple
}

func NewPointLight(intensity *Color, position *Tuple) *PointLight {
	return &PointLight{
		Intensity: intensity,
		Position:  position,
	}
}
