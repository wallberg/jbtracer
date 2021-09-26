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
		if !r1.Origin.Equal(t1) {
			return fmt.Errorf("Expected %s.origin = %v; got %v", r1name, t1, r1.Origin)
		}
	} else {
		if !r1.Direction.Equal(t1) {
			return fmt.Errorf("Expected %s.direction = %v; got %v", r1name, t1, r1.Direction)
		}
	}
	return nil
}

func rayPositionEqualPoint(r1name string, t, x, y, z float64) error {
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

func transform(r1name, r2name, m1name string) error {
	if r2, ok = rays[r2name]; !ok {
		return fmt.Errorf("Unknown symbol (ray) %s", r2name)
	}
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}
	rays[r1name] = r2.Transform(m1)
	return nil
}

func rayEqualOriginPoint(r1name string, x, y, z float64) error {
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", r1name)
	}
	expected := NewPoint(x, y, z)
	got := r1.Origin
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s.origin = %v; got %v", r1name, expected, got)
	}
	return nil
}

func rayEqualDirectionVector(r1name string, x, y, z float64) error {
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol (ray) %s", r1name)
	}
	expected := NewVector(x, y, z)
	got := r1.Direction
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s.direction = %v; got %v", r1name, expected, got)
	}
	return nil
}
