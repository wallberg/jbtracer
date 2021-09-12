package jbtracer

import "fmt"

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

func intersectionT(i1name string, index int, t float64) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	expected := t
	got := i1[index].t
	if !EqualFloat64(got, expected) {
		return fmt.Errorf("Expected %s[%d].t = %f; got %f", i1name, index, expected, got)
	}
	return nil
}

func intersectionObject(i1name string, index int, sph1name string) error {
	if i1, ok = intersections[i1name]; !ok {
		return fmt.Errorf("Unknown symbol (intersection) %s", i1name)
	}
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere): %s", sph1name)
	}
	expected := sph1
	var object = *(i1[index].object)
	if got, ok := object.(*Sphere); ok {
		if got != expected {
			return fmt.Errorf("Expected %s[%d].object = %v; got %v", i1name, index, expected, got)
		}
	}
	return nil
}

func intersection(i1name string, t float64, o1name string) error {
	if o1, ok = objects[o1name]; !ok {
		return fmt.Errorf("Unknown symbol (object) %s", o1name)
	}
	intersections[i1name] = []*Intersection{NewIntersection(o1, t)}
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
	hit := make(Intersections, 0)
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
