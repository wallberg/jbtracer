package jbtracer

import (
	"math"
)

// GradientPattern blends two colors, linearly interpolating from
// A to B as the x coordinate changes from one integer value to the next.
// eg, [0,1) or [1,2)
type GradientPattern struct {
	A, B      *Color
	transform *Matrix
}

func NewGradientPattern(a, b *Color) *GradientPattern {
	return &GradientPattern{
		A:         a,
		B:         b,
		transform: IdentityMatrix(),
	}
}

func (a *GradientPattern) Equal(b Pattern) bool {
	if a == nil || b == nil {
		return false
	}

	if bStripe, ok := b.(*GradientPattern); !ok {
		return false
	} else {
		return a.A.Equal(bStripe.A) && a.B.Equal(bStripe.B)
	}
}

func (pattern *GradientPattern) Transform() *Matrix {
	return pattern.transform
}

func (pattern *GradientPattern) SetTransform(transform *Matrix) {
	pattern.transform = transform
}

func (pattern *GradientPattern) PatternAt(point *Tuple) *Color {
	distance := pattern.B.Subtract(pattern.A)
	fraction := point.X - math.Floor(point.X)

	return pattern.A.Add(distance.MultiplyScalar(fraction))
}
