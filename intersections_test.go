package jbtracer

import "fmt"

func intersect(i1name, sph1name, r1name string) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere): %s", sph1name)
	}
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol (ray): %s", r1name)
	}
	intersections[i1name] = Intersections(sph1, r1)
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

func intersectionsT(i1name string, index int, t float64) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	expected := t
	got := i1[index].T
	if !EqualFloat64(got, expected) {
		return fmt.Errorf("Expected %s[%d].t = %f; got %f", i1name, index, expected, got)
	}
	return nil
}

func intersectionT(i1name string, t float64) error {
	return intersectionsT(i1name, 0, t)
}

func intersectionsObject(i1name string, index int, o1name string) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersections) %s", i1name)
	}
	if sh1, ok = shapes[o1name]; !ok {
		return fmt.Errorf("Unknown symbol (object): %s", o1name)
	}
	expected := sh1
	got := i1[index].Object
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s[%d].object = %v; got %v", i1name, index, expected, got)
	}
	return nil
}

func intersectionObject(i1name, o1name string) error {
	return intersectionsObject(i1name, 0, o1name)
}

func intersection(i1name string, t float64, o1name string) error {
	if sh1, ok = shapes[o1name]; !ok {
		return fmt.Errorf("Unknown symbol (object) %s", o1name)
	}
	intersections[i1name] = []*Intersection{NewIntersection(sh1, t)}
	return nil
}

func intersectionConcat4(i1name, i2name, i3name, i4name, i5name string) error {
	if i2, ok = intersections[i2name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i2name)
	}
	if i3, ok = intersections[i3name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i3name)
	}
	if i4, ok = intersections[i4name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i4name)
	}
	if i5, ok = intersections[i5name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i5name)
	}
	is := append(i2, i3...)
	is = append(is, i4...)
	is = append(is, i5...)
	intersections[i1name] = is
	return nil
}

func intersectionConcat(i1name, i2name, i3name string) error {
	if i2, ok = intersections[i2name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i2name)
	}
	if i3, ok = intersections[i3name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i3name)
	}
	intersections[i1name] = append(i2, i3...)
	return nil
}

func intersectionHits(i1name, i2name string) error {
	if i2, ok = intersections[i2name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i2name)
	}
	hit := make(IntersectionSlice, 0)
	i := i2.Hit()
	if i != nil {
		hit = append(hit, i)
	}
	intersections[i1name] = hit
	return nil
}

func intersectionEqual(i1name, i2name string) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	if i2, ok = intersections[i2name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i2name)
	}
	expected := i2
	got := i1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s = %v; got %v", i1name, expected, got)
	}
	return nil
}

func intersectionEmpty(i1name string) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	if len(i1) != 0 {
		return fmt.Errorf("Expected %s is empty; got len()=%d", i1name, len(i1))
	}
	return nil
}

func comp(i1name, r1name string) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol (ray): %s", r1name)
	}
	comps = i1[0].PreparedComputations(r1)
	return nil
}

func compEqualObject(i1name string) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	expected := i1[0].Object
	got := comps.Object
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s[0].object = %v; got %v", i1name, expected, got)
	}
	return nil
}

func compEqualT(i1name string) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	expected := i1[0].T
	got := comps.T
	if got != expected {
		return fmt.Errorf("Expected %s[0].t = %v; got %v", i1name, expected, got)
	}
	return nil
}

func compEqualPoint(x, y, z float64) error {
	expected := NewPoint(x, y, z)
	got := comps.Point
	if !got.Equal(expected) {
		return fmt.Errorf("Expected comps.point = %v; got %v", expected, got)
	}
	return nil
}

func compEqualEyeV(x, y, z float64) error {
	expected := NewVector(x, y, z)
	got := comps.EyeV
	if !got.Equal(expected) {
		return fmt.Errorf("Expected comps.eyev = %v; got %v", expected, got)
	}
	return nil
}

func compEqualNormalV(x, y, z float64) error {
	expected := NewVector(x, y, z)
	got := comps.NormalV
	if !got.Equal(expected) {
		return fmt.Errorf("Expected comps.normalv = %v; got %v", expected, got)
	}
	return nil
}

func compEqualInside(bool string) error {
	expected := (bool == "true")
	got := comps.Inside
	if got != expected {
		return fmt.Errorf("Expected comps.inside = %v; got %v", expected, got)
	}
	return nil
}

func compOverPointZLessThanEpsilon() error {
	expected := true
	got := comps.OverPoint.Z < -1*Epsilon/2
	if got != expected {
		return fmt.Errorf("Expected comps.over_point.z < -EPSILON/2 = %v; got %v", expected, got)
	}
	return nil
}

func compPointZGreaterThanOverPointZ() error {
	expected := true
	got := comps.Point.Z > comps.OverPoint.Z
	if got != expected {
		return fmt.Errorf("Expected comps.point. z > comps.over_point.z = %v; got %v", expected, got)
	}
	return nil
}
