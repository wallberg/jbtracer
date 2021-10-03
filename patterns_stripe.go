package jbtracer

import (
	"math"
)

type StripePattern struct {
	A, B      *Color
	transform *Matrix
}

func NewStripePattern(a, b *Color) *StripePattern {
	return &StripePattern{
		A:         a,
		B:         b,
		transform: IdentityMatrix(),
	}
}

func (a *StripePattern) Equal(b Pattern) bool {
	if a == nil || b == nil {
		return false
	}

	if bStripe, ok := b.(*StripePattern); !ok {
		return false
	} else {
		return a.A.Equal(bStripe.A) && a.B.Equal(bStripe.B)
	}
}

func (pattern *StripePattern) Transform() *Matrix {
	return pattern.transform
}

func (pattern *StripePattern) SetTransform(transform *Matrix) {
	pattern.transform = transform
}

func (pattern *StripePattern) PatternAt(point *Tuple) *Color {
	if int(math.Floor(point.X))%2 == 0 {
		return pattern.A
	} else {
		return pattern.B
	}
}
