package jbtracer

import "fmt"

func camera(hsize, vsize int, fov float32) error {
	cam = NewCamera(hsize, vsize, fov)
	return nil
}

func cameraEqualHsize(hsize int) error {
	expected := hsize
	got := cam.Hsize
	if got != expected {
		return fmt.Errorf("Expected c.hsize = %v; got %v", expected, got)
	}
	return nil
}

func cameraEqualVsize(vsize int) error {
	expected := vsize
	got := cam.Vsize
	if got != expected {
		return fmt.Errorf("Expected c.vsize = %v; got %v", expected, got)
	}
	return nil
}

func cameraEqualFOV(fov float32) error {
	expected := fov
	got := cam.FieldOfView
	if got != expected {
		return fmt.Errorf("Expected c.field_of_view = %v; got %v", expected, got)
	}
	return nil
}

func cameraEqualTransform(m1name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}

	expected := m1
	got := cam.Transform

	if !got.Equal(expected) {
		return fmt.Errorf("Expected c.transform = %v; got %v", expected, got)
	}
	return nil
}

func cameraEqualPixelSize(pixelSize float32) error {
	expected := pixelSize
	got := cam.PixelSize
	if got != expected {
		return fmt.Errorf("Expected c.pixel_size = %v; got %v", expected, got)
	}
	return nil
}

func cameraTransform(m1name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}
	cam.Transform = m1
	return nil
}

func cameraRayForPixel(r1name string, x, y int) {
	rays[r1name] = cam.RayForPixel(x, y)
}

func render() error {
	c = cam.Render(w)
	return nil
}
