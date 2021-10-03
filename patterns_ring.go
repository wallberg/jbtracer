package jbtracer

import (
	"math"
)

// RingPattern creates a circular stripe, based on x,z radius
type RingPattern struct {
	A, B      *Color
	transform *Matrix
}

func NewRingPattern(a, b *Color) *RingPattern {
	return &RingPattern{
		A:         a,
		B:         b,
		transform: IdentityMatrix(),
	}
}

func (a *RingPattern) Equal(b Pattern) bool {
	if a == nil || b == nil {
		return false
	}

	if bStripe, ok := b.(*RingPattern); !ok {
		return false
	} else {
		return a.A.Equal(bStripe.A) && a.B.Equal(bStripe.B)
	}
}

func (pattern *RingPattern) Transform() *Matrix {
	return pattern.transform
}

func (pattern *RingPattern) SetTransform(transform *Matrix) {
	pattern.transform = transform
}

func (pattern *RingPattern) PatternAt(point *Tuple) *Color {
	radius := math.Sqrt(point.X*point.X + point.Z*point.Z)
	if int(math.Floor(radius))%2 == 0 {
		return pattern.A
	} else {
		return pattern.B
	}
}
