package jbtracer

import (
	"math"
)

// CheckersPattern creates a 3D pattern of alternating cube colors
type CheckersPattern struct {
	A, B      *Color
	transform *Matrix
}

func NewCheckersPattern(a, b *Color) *CheckersPattern {
	return &CheckersPattern{
		A:         a,
		B:         b,
		transform: IdentityMatrix(),
	}
}

func (a *CheckersPattern) Equal(b Pattern) bool {
	if a == nil || b == nil {
		return false
	}

	if bStripe, ok := b.(*CheckersPattern); !ok {
		return false
	} else {
		return a.A.Equal(bStripe.A) && a.B.Equal(bStripe.B)
	}
}

func (pattern *CheckersPattern) Transform() *Matrix {
	return pattern.transform
}

func (pattern *CheckersPattern) SetTransform(transform *Matrix) {
	pattern.transform = transform
}

func (pattern *CheckersPattern) PatternAt(point *Tuple) *Color {
	floor := math.Floor(point.X) + math.Floor(point.Y) + math.Floor(point.Z)
	if int(floor)%2 == 0 {
		return pattern.A
	} else {
		return pattern.B
	}
}
