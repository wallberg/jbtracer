package jbtracer

import "fmt"

func rayPointVector(r1name string, xp, yp, zp, xv, yv, zv float32) error {
	p := NewPoint(xp, yp, zp)
	v := NewVector(xv, yv, zv)
	rays[r1name] = NewRay(p, v)
	return nil
}

func sphere(s1name string) error {
	spheres[s1name] = NewSphere()
	return nil
}

func intersect(i1name, sph1name, r1name string) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere): %s", sph1name)
	}
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol (ray): %s", r1name)
	}
	intersections[i1name] = sph1.Intersections(r1)
	return nil
}

func intersectionCount(i1name string, count int) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	expected := count
	got := len(i1)
	if got != expected {
		return fmt.Errorf("Expected %s.count = %d; got %d", i1name, expected, got)
	}
	return nil
}

func intersectionIndex(i1name string, index int, t float32) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	expected := t
	got := i1[index].t
	if !EqualFloat32(got, expected) {
		return fmt.Errorf("Expected %s[%d].t = %f; got %f", i1name, index, expected, got)
	}
	return nil
}
