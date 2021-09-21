package jbtracer

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/cucumber/godog"
)

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

func sphereWith(sph1name string, table *godog.Table) error {
	reTuple := regexp.MustCompile(`^\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`)
	reScalar := regexp.MustCompile(`^(-?\d+(?:\.\d+)?)$`)
	reScaling := regexp.MustCompile(`^scaling\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`)

	rows := len(table.Rows)
	if rows < 1 {
		return fmt.Errorf("sphereWith() requires a table with at least one row")
	}

	cols := len(table.Rows[0].Cells)
	if cols != 2 {
		return fmt.Errorf("sphereWith() requires a table with 2 columns")
	}

	sph1 = NewSphere()
	for _, row := range table.Rows {
		name := row.Cells[0].Value
		value := row.Cells[1].Value

		switch name {
		case "material.color":
			if m := reTuple.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract tuple from %s", value)
			} else {
				var red, green, blue float32
				if f, err := strconv.ParseFloat(m[1], 32); err != nil {
					return err
				} else {
					red = float32(f)
				}
				if f, err := strconv.ParseFloat(m[2], 32); err != nil {
					return err
				} else {
					green = float32(f)
				}
				if f, err := strconv.ParseFloat(m[3], 32); err != nil {
					return err
				} else {
					blue = float32(f)
				}
				sph1.Material().Color = NewColor(red, green, blue)
			}
		case "material.diffuse":
			if m := reScalar.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract scalar from %s", value)
			} else {
				var diffuse float32
				if f, err := strconv.ParseFloat(m[1], 32); err != nil {
					return err
				} else {
					diffuse = float32(f)
				}
				sph1.Material().Diffuse = diffuse
			}
		case "material.specular":
			if m := reScalar.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract scalar from %s", value)
			} else {
				var specular float32
				if f, err := strconv.ParseFloat(m[1], 32); err != nil {
					return err
				} else {
					specular = float32(f)
				}
				sph1.Material().Specular = specular
			}
		case "transform":
			if m := reScaling.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract scaling from %s", value)
			} else {
				var x, y, z float32
				if f, err := strconv.ParseFloat(m[1], 32); err != nil {
					return err
				} else {
					x = float32(f)
				}
				if f, err := strconv.ParseFloat(m[2], 32); err != nil {
					return err
				} else {
					y = float32(f)
				}
				if f, err := strconv.ParseFloat(m[3], 32); err != nil {
					return err
				} else {
					z = float32(f)
				}
				sph1.Transform = Scaling(x, y, z)
			}
		}
	}

	spheres[sph1name] = sph1
	return nil
}
