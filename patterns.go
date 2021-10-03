package jbtracer

import (
	"log"
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
