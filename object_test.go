package jbtracer

import "fmt"

func objectEqualMaterialColor(c1name, o1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c1name)
	}
	if o1, ok = objects[o1name]; !ok {
		return fmt.Errorf("Unknown symbol (object) %s", o1name)
	}

	expected := o1.Material().Color
	got := c1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected color %s = %v; got %v", c1name, expected, got)
	}
	return nil
}

func objectMaterialAmbient(o1name string, scalar float64) error {
	if o1, ok = objects[o1name]; !ok {
		return fmt.Errorf("Unknown symbol (object) %s", o1name)
	}
	o1.Material().Ambient = scalar
	return nil
}
