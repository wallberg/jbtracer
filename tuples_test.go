package jbtracer

import (
	"fmt"
	"strconv"
)

func isPoint(t1name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if !t1.IsPoint() {
		return fmt.Errorf("Expected %s.isPoint()=true; got false", t1name)
	}
	return nil
}

func isVector(t1name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if !t1.IsVector() {
		return fmt.Errorf("Expected %s.isVector()=true; got false", t1name)
	}
	return nil
}

func isNotPoint(t1name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.IsPoint() {
		return fmt.Errorf("Expected %s.isPoint()=false; got true", t1name)
	}
	return nil
}

func isNotVector(t1name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.IsVector() {
		return fmt.Errorf("Expected %s.isVector()=false; got true", t1name)
	}
	return nil
}

func tuple(t1name string, x, y, z, w float64) error {
	tuples[t1name] = &Tuple{X: x, Y: y, Z: z, W: w}
	return nil
}

func equalsTupleX(t1name string, x float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.X != x {
		return fmt.Errorf("Expected %f for %s.x; got %f", x, t1name, t1.X)
	}
	return nil
}

func equalsTupleY(t1name string, y float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.Y != y {
		return fmt.Errorf("Expected %f for %s.y; got %f", y, t1name, t1.Y)
	}
	return nil
}

func equalsTupleZ(t1name string, z float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.Z != z {
		return fmt.Errorf("Expected %f for %s.z; got %f", z, t1name, t1.Z)
	}
	return nil
}

func equalsTupleW(t1name string, w float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.W != w {
		return fmt.Errorf("Expected %f for %s.w; got %f", w, t1name, t1.W)
	}
	return nil
}

func point(t1name string, x, y, z float64) error {
	tuples[t1name] = NewPoint(x, y, z)
	return nil
}

func vector(t1name string, x, y, z float64) error {
	tuples[t1name] = NewVector(x, y, z)
	return nil
}

func equalsTuple(t1name string, x, y, z, w float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	if !t1.Equal(expected) {
		return fmt.Errorf("Expected %s=%v; got %v", t1name, expected, t1)
	}
	return nil
}

func equalsTupleAdd(t1name string, t2name string, x, y, z, w float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	got := t1.Add(t2)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s + %s=%v; got %v", t1name, t2name, expected, got)
	}
	return nil
}

func equalsTupleSubtract(t1name string, t2name string, ttype string, x, y, z float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	if ttype == "point" {
		expected = NewPoint(x, y, z)
	} else {
		expected = NewVector(x, y, z)
	}
	got = t1.Subtract(t2)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s + %s=%v; got %v", t1name, t2name, expected, got)
	}
	return nil
}

func equalsTupleNegate(t1name string, x, y, z, w float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected = &Tuple{X: x, Y: y, Z: z, W: w}
	got = t1.Negate()
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s=%v; got %v", t1name, expected, got)
	}
	return nil
}

func equalsTupleMultiply(t1name string, scalar float64, x, y, z, w float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	got := t1.Multiply(scalar)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s + %f=%v; got %v", t1name, scalar, expected, got)
	}
	return nil
}

func equalsTupleDivide(t1name string, scalar float64, x, y, z, w float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	got := t1.Divide(scalar)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s + %f=%v; got %v", t1name, scalar, expected, got)
	}
	return nil
}

func equalsTupleMagnitude(t1name string, expected float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	got := t1.Magnitude()
	if !EqualFloat64(got, expected) {
		return fmt.Errorf("Expected %s=%f; got %f", t1name, expected, got)
	}
	return nil
}

func normalize(t1name string, t2name string) error {
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	tuples[t1name] = t2.Normalize()
	return nil
}

func equalsVectorNormalize(t1name string, x, y, z float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected = NewVector(x, y, z)
	got = t1.Normalize()
	if !got.Equal(expected) {
		return fmt.Errorf("Expected normalize(%s)=%v; got %v", t1name, expected, got)
	}
	return nil
}

func equalsVectorDot(t1name string, t2name string, expected float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	got := t1.Dot(t2)
	if got != expected {
		return fmt.Errorf("Expected dot(%s, %s)=%v; got %v", t1name, t2name, expected, got)
	}
	return nil
}

func equalsVectorCross(t1name string, t2name string, x, y, z float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	expected = NewVector(x, y, z)
	got = t1.Cross(t2)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected cross(%s, %s)=%v; got %v", t1name, t2name, expected, got)
	}
	return nil
}

func color(t1name string, red, green, blue float64) error {
	colors[t1name] = &Color{red, green, blue}
	return nil
}

func equalsColorOp(c1name string, op string, c2name string, red, green, blue float64) error {

	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", c1name)
	}

	var scalar float64
	if c2, ok = colors[c2name]; !ok {
		// Not a known symbol, see if this is a scalar
		if f, err := strconv.ParseFloat(c2name, 64); err == nil {
			// It's a scalar
			scalar = f
			if op != "*" {
				return fmt.Errorf("Can't perform %s operation on scalar", op)
			}
		} else {
			// Not a scalar
			return fmt.Errorf("Unknown symbol %s", c2name)
		}
	}

	expected := &Color{red, green, blue}
	var got *Color

	if c2 != nil {
		switch op {
		case "+":
			got = c1.Add(c2)
		case "-":
			got = c1.Subtract(c2)
		case "*":
			got = c1.Multiply(c2)
		}
	} else {
		got = c1.MultiplyScalar(scalar)

	}
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s %s %s=%v; got %v", c1name, op, c2name, expected, got)
	}
	return nil
}

func equalsColorField(c1name string, field string, expected float64) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", c1name)
	}

	var got float64
	switch field {
	case "red":
		got = c1.Red
	case "green":
		got = c1.Green
	case "blue":
		got = c1.Blue
	}

	if got != expected {
		return fmt.Errorf("Expected %s.%s = %f; got %f", c1name, field, expected, got)
	}
	return nil
}

func vectorEqual(t1name string, x, y, z float64) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected := NewVector(x, y, z)
	got := t1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s=%v; got %v", t1name, expected, got)
	}
	return nil
}

func tupleEqualNormalize(t1name, t2name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t2name)
	}
	expected := t2.Normalize()
	got := t1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s=%v; got %v", t1name, expected, got)
	}
	return nil
}
