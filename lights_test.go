package jbtracer

import "fmt"

func pointLight(t1name, c1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c1name)
	}
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t1name)
	}
	light = NewPointLight(c1, t1)
	return nil
}

func pointLightIntensity(c1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c1name)
	}
	expected := c1
	got := light.Intensity
	if !got.Equal(expected) {
		return fmt.Errorf("Expected light.intensity=%v; got %v", expected, got)
	}
	return nil
}

func pointLightPosition(t1name string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t1name)
	}
	expected := t1
	got := light.Position
	if !got.Equal(expected) {
		return fmt.Errorf("Expected light.position=%v; got %v", expected, got)
	}
	return nil
}

func pointLight2(x, y, z, red, green, blue float32) error {
	c1 = &Color{red, green, blue}
	t1 = NewPoint(x, y, z)
	light = NewPointLight(c1, t1)
	return nil
}
