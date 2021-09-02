package jbtracer

func matrixTranslation(m1name string, x, y, z float32) error {
	matrices[m1name] = Translation(x, y, z)
	return nil
}

func matrixScaling(m1name string, x, y, z float32) error {
	matrices[m1name] = Scaling(x, y, z)
	return nil
}

func matrixRotation(m1name, axis string, radians float32) error {
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
