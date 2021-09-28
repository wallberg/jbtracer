package jbtracer

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/cucumber/godog"
)

func rayPointVector(r1name string, xp, yp, zp, xv, yv, zv float64) error {
	p := NewPoint(xp, yp, zp)
	v := NewVector(xv, yv, zv)
	rays[r1name] = NewRay(p, v)
	return nil
}

func sphere(s1name string) error {
	sph1 = NewSphere()
	spheres[s1name] = sph1
	shapes[s1name] = sph1

	return nil
}

func sphereNormalAt(t1name, sph1name string, x, y, z float64) error {
	if sph1, ok = spheres[sph1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sph1name)
	}

	tuples[t1name] = NormalAt(sph1, NewPoint(x, y, z))
	return nil

}

func sphereWith(sph1name string, table *godog.Table) error {
	reTuple := regexp.MustCompile(`^\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`)
	reScalar := regexp.MustCompile(`^(-?\d+(?:\.\d+)?)$`)
	reTransform := regexp.MustCompile(`^(scaling|translation)\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`)

	rows := len(table.Rows)
	if rows < 1 {
		return fmt.Errorf("sphereWith() requires a table with at least one row")
	}

	cols := len(table.Rows[0].Cells)
	if cols != 2 {
		return fmt.Errorf("sphereWith() requires a table with 2 columns")
	}

	sph1 = NewSphere()
	var err error
	for _, row := range table.Rows {
		name := row.Cells[0].Value
		value := row.Cells[1].Value

		switch name {
		case "material.color":
			if m := reTuple.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract tuple from %s", value)
			} else {
				var red, green, blue float64
				if red, err = strconv.ParseFloat(m[1], 64); err != nil {
					return err
				}
				if green, err = strconv.ParseFloat(m[2], 64); err != nil {
					return err
				}
				if blue, err = strconv.ParseFloat(m[3], 64); err != nil {
					return err
				}
				sph1.Material().Color = NewColor(red, green, blue)
			}
		case "material.diffuse":
			if m := reScalar.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract scalar from %s", value)
			} else {
				var diffuse float64
				if diffuse, err = strconv.ParseFloat(m[1], 64); err != nil {
					return err
				}
				sph1.Material().Diffuse = diffuse
			}
		case "material.specular":
			if m := reScalar.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract scalar from %s", value)
			} else {
				var specular float64
				if specular, err = strconv.ParseFloat(m[1], 64); err != nil {
					return err
				}
				sph1.Material().Specular = specular
			}
		case "transform":
			if m := reTransform.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract transform from %s", value)
			} else {
				var x, y, z float64
				if x, err = strconv.ParseFloat(m[2], 64); err != nil {
					return err
				}
				if y, err = strconv.ParseFloat(m[3], 64); err != nil {
					return err
				}
				if z, err = strconv.ParseFloat(m[4], 64); err != nil {
					return err
				}
				switch m[1] {
				case "scaling":
					sph1.SetTransform(Scaling(x, y, z))
				case "translation":
					sph1.SetTransform(Translation(x, y, z))
				}
			}
		}
	}

	spheres[sph1name] = sph1
	shapes[sph1name] = sph1
	return nil
}
