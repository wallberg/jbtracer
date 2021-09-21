package jbtracer

import "fmt"

func world() error {
	w = NewWorld()
	return nil
}

func worldContainsNoObjects() error {
	expected := 0
	got := len(w.Objects)
	if got != expected {
		return fmt.Errorf("Expected World w to have %d objects; got %d", expected, got)
	}
	return nil
}

func worldHasNoLightSource() error {
	got := w.Light
	if got != nil {
		return fmt.Errorf("Expected World w to have no light source; got %v", got)
	}
	return nil
}

func worldDefault() error {
	w = DefaultWorld()
	return nil
}

func worldContainsSphere(sph1name string) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sph1name)
	}

	for _, object := range w.Objects {
		// Check if the object is a *Sphere
		if sph2, ok := object.(*Sphere); ok {
			if sph1.Equal(sph2) {
				return nil
			}
		}
	}

	return fmt.Errorf("Expected World w to contain sphere %v; it did not", sph1)
}

func worldLight() error {
	if !w.Light.Equal(light) {
		return fmt.Errorf("Expected w.light = light; got not")
	}
	return nil
}

func worldIntersect(i1name, r1name string) error {
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", r1name)
	}

	intersections[i1name] = w.Intersections(r1)
	return nil
}
