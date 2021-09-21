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

func (a *PointLight) Equal(b *PointLight) bool {
	return a != nil && b != nil && a.Intensity.Equal(b.Intensity) && a.Position.Equal(b.Position)
}
