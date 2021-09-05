package jbtracer

import "fmt"

func ray(r1name, t1name, t2name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	rays[r1name] = NewRay(t1, t2)
	return nil
}

func rayEqualField(r1name, op, t1name string) error {
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", r1name)
	}
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if op == "origin" {
		if !r1.origin.Equal(t1) {
			return fmt.Errorf("Expected %s.origin = %v; got %v", r1name, t1, r1.origin)
		}
	} else {
		if !r1.direction.Equal(t1) {
			return fmt.Errorf("Expected %s.direction = %v; got %v", r1name, t1, r1.direction)
		}
	}
	return nil
}

func rayPositionEqualPoint(r1name string, t, x, y, z float32) error {
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", r1name)
	}
	expected := NewPoint(x, y, z)
	got := r1.Position(t)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected position(%s,%f) = %v; got %v", r1name, t, expected, got)
	}
	return nil
}
