package jbtracer

import "fmt"

func matrixTranslation(m1name string, x, y, z float64) error {
	matrices[m1name] = Translation(x, y, z)
	return nil
}

func matrixScaling(m1name string, x, y, z float64) error {
	matrices[m1name] = Scaling(x, y, z)
	return nil
}

func matrixRotation(m1name, axis string, radians float64) error {
	var axisInt int
	switch axis {
	case "x":
		axisInt = Axis_X
	case "y":
		axisInt = Axis_Y
	case "z":
		axisInt = Axis_Z
	}
	matrices[m1name] = Rotation(axisInt, radians)
	return nil
}

func matrixShearing(m1name string, xY, xZ, yX, yZ, zX, zY float64) error {
	matrices[m1name] = Shearing(xY, xZ, yX, yZ, zX, zY)
	return nil
}

func tupleMatrixAssign(t1name, m1name, t2name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	tuples[t1name] = m1.MultiplyTuple(t2)
	return nil
}

func tupleEqual(t1name, t2name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	expected := t1
	got := t2
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s = %v; got %v", t2name, expected, got)
	}
	return nil
}

func matrixViewTransform(m1name, t1name, t2name, t3name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t2name)
	}
	if t3, ok = tuples[t3name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t3name)
	}
	matrices[m1name] = ViewTransform(t1, t2, t3)
	return nil
}
