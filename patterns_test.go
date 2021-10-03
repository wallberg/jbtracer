package jbtracer

import (
	"fmt"
)

//
// TestPattern
//
type TestPattern struct {
	transform *Matrix
}

func NewTestPattern() *TestPattern {
	return &TestPattern{
		transform: IdentityMatrix(),
	}
}

func (a *TestPattern) Equal(b Pattern) bool {
	if a == nil || b == nil {
		return false
	}

	if _, ok := b.(*TestPattern); !ok {
		return false
	} else {
		return true
	}
}

func (pattern *TestPattern) Transform() *Matrix {
	return pattern.transform
}

func (pattern *TestPattern) SetTransform(transform *Matrix) {
	pattern.transform = transform
}

func (pattern *TestPattern) PatternAt(point *Tuple) *Color {
	return NewColor(point.X, point.Y, point.Z)
}

//
// StripePattern
//

func patternStripePattern(c1name, c2name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (colo) %s", c1name)
	}
	if c2, ok = colors[c2name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c2name)
	}

	pattern = NewStripePattern(c1, c2)
	return nil
}

func patternEqualA(c1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c1name)
	}
	if stripe, ok = pattern.(*StripePattern); !ok {
		return fmt.Errorf("pattern is not a StripePattern")
	}

	got := stripe.A
	expected := c1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected stripe.A=%v; got %v", expected, got)
	}
	return nil
}

func patternEqualB(c1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c1name)
	}
	if stripe, ok = pattern.(*StripePattern); !ok {
		return fmt.Errorf("pattern is not a StripePattern")
	}

	got := stripe.B
	expected := c1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected stripe.B=%v; got %v", expected, got)
	}
	return nil
}

func patternEqualStripeAt(x, y, z float64, c1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c1name)
	}
	if stripe, ok = pattern.(*StripePattern); !ok {
		return fmt.Errorf("pattern is not a StripePattern")
	}
	point := NewPoint(x, y, z)

	got := stripe.PatternAt(point)
	expected := c1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected stripe.PatternAt(%v)=%v; got %v", point, expected, got)
	}
	return nil
}

//
// Pattern
//

func patternPatternAtShape(c1name, sh1name, t1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape) %s", sh1name)
	}
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t1name)
	}

	colors[c1name] = PatternAt(pattern, sh1, t1)
	return nil
}

func patternSetPatternTransform(m1name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}

	pattern.SetTransform(m1)
	return nil
}

func patternEqualTransform(m1name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}

	got := m1
	expected := pattern.Transform()
	if !got.Equal(expected) {
		return fmt.Errorf("Expected pattern.Transform()=%v; got %v", expected, got)
	}
	return nil
}

//
// TestPattern
//

func patternTestPattern() error {
	pattern = NewTestPattern()
	return nil
}
