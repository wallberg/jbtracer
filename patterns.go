package jbtracer

import (
	"log"
	"math"
)

//
// Pattern
//
type Pattern interface {
	Transform() *Matrix
	SetTransform(transform *Matrix)
	PatternAt(point *Tuple) *Color
	Equal(b Pattern) bool
}

func PatternAt(pattern Pattern, object Shape, worldPoint *Tuple) *Color {
	var (
		shapeInv   *Matrix
		err        error
		patternInv *Matrix
	)

	if shapeInv, err = object.Transform().Inverse(); err != nil {
		log.Fatalf("Matrix shape.Transform()=%v is not invertible", object.Transform())
	}
	objectPoint := shapeInv.MultiplyTuple(worldPoint)

	if patternInv, err = pattern.Transform().Inverse(); err != nil {
		log.Fatalf("Matrix pattern.Transform()=%v is not invertible", pattern.Transform())
	}
	patternPoint := patternInv.MultiplyTuple(objectPoint)

	return pattern.PatternAt(patternPoint)
}

//
// StripePattern
//
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
