package jbtracer

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/cucumber/godog"
)

func plane(sh1name string) {
	shapes[sh1name] = NewPlane()
}

func planeWith(sh1name string, table *godog.Table) error {
	// reTuple := regexp.MustCompile(`^\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`)
	reScalar := regexp.MustCompile(`^(-?\d+(?:\.\d+)?)$`)
	reTransform := regexp.MustCompile(`^(translation)\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`)

	rows := len(table.Rows)
	if rows < 1 {
		return fmt.Errorf("planeWith() requires a table with at least one row")
	}

	cols := len(table.Rows[0].Cells)
	if cols != 2 {
		return fmt.Errorf("planeWith() requires a table with 2 columns")
	}

	sh1 = NewPlane()
	var err error
	for _, row := range table.Rows {
		name := row.Cells[0].Value
		value := row.Cells[1].Value

		switch name {
		case "material.reflective":
			if m := reScalar.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract scalar from %s", value)
			} else {
				var reflective float64
				if reflective, err = strconv.ParseFloat(m[1], 64); err != nil {
					return err
				}
				sh1.Material().Reflective = reflective
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
				case "translation":
					sh1.SetTransform(Translation(x, y, z))
				default:
					return fmt.Errorf("Unknown transform %s in planeWith()", m[1])
				}
			}
		default:
			return fmt.Errorf("Unknown field %s in planeWith()", name)
		}
	}

	shapes[sh1name] = sh1
	return nil
}
