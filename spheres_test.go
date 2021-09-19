package jbtracer

import "fmt"

func rayPointVector(r1name string, xp, yp, zp, xv, yv, zv float32) error {
	p := NewPoint(xp, yp, zp)
	v := NewVector(xv, yv, zv)
	rays[r1name] = NewRay(p, v)
	return nil
}

func sphere(s1name string) error {
	sph1 = NewSphere()
	spheres[s1name] = sph1

	var object Object = sph1
	objects[s1name] = &object

	return nil
}

func sphereEqualTransform(sph1name, m1name string) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sph1name)
	}
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}

	expected := m1
	got := sph1.Transform
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s.transform = %v; got %v", sph1name, expected, got)
	}
	return nil

}

func sphereTransform(sph1name, m1name string) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sph1name)
	}
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}

	sph1.Transform = m1
	// if sph1.Transform, err = m1.Inverse(); err != nil {
	// 	return fmt.Errorf("Matrix %s is not invertible", m1name)
	// }
	return nil

}

func sphereNormalAt(t1name, sph1name string, x, y, z float32) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sph1name)
	}

	tuples[t1name] = sph1.NormalAt(NewPoint(x, y, z))
	return nil

}

func sphereMaterial(mat1name, sph1name string) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sph1name)
	}
	materials[mat1name] = sph1.material
	return nil
}

func sphereMaterial2(sph1name, mat1name string) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sph1name)
	}
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	sph1.material = mat1
	return nil
}
