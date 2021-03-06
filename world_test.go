package jbtracer

import (
	"fmt"
	"time"
)

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

func worldPointLight(x, y, z, red, green, blue float64) error {
	c1 = &Color{red, green, blue}
	t1 = NewPoint(x, y, z)
	w.Light = NewPointLight(c1, t1)
	return nil
}

func worldShadeHit(c1name string) error {
	colors[c1name] = w.ShadeHit(comps, DefaultReflectedDepth)
	return nil
}

func worldToObject(o1name, sindex string) error {
	var index int
	if sindex == "first" {
		index = 0
	} else {
		index = 1
	}
	shapes[o1name] = w.Objects[index]
	return nil
}

func worldColorAt(c1name, r1name string) error {
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", r1name)
	}
	colors[c1name] = w.ColorAt(r1, DefaultReflectedDepth)
	return nil
}

func worldIsShadowed(t1name, flag string) error {
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t1name)
	}
	expected := true
	if flag == "false" {
		expected = false
	}
	got := w.IsShadowed(t1)
	if got != expected {
		return fmt.Errorf("Expected point %s in shadow = %v; got %v", t1name, expected, got)
	}
	return nil
}

func worldAddObject(o1name string) error {
	if sh1, ok = shapes[o1name]; !ok {
		return fmt.Errorf("Unknown symbol (object) %s", o1name)
	}
	w.AddObject(sh1)
	return nil
}

func worldReflectedColor(c1name string) error {
	colors[c1name] = w.ReflectedColor(comps, DefaultReflectedDepth)
	return nil
}

func worldReflectedColorDepth(c1name string, depth int) error {
	colors[c1name] = w.ReflectedColor(comps, depth)
	return nil
}

func worldColorAtTerminates(r1name string) error {
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", r1name)
	}

	// Run ColorAt in a goroutine
	c := make(chan bool)
	go func() {
		w.ColorAt(r1, DefaultReflectedDepth)
		c <- true
	}()

	var n time.Duration = 1
	timer := time.NewTimer(n * time.Second)
	select {
	case <-c:
		// ColorAt terminated successfully
	case <-timer.C:
		return fmt.Errorf("w.ColorAt() did not terminate within %d second(s)", n)
	}
	return nil
}
