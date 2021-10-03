package jbtracer

import (
	"fmt"
)

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
