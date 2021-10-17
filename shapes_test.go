package jbtracer

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/cucumber/godog"
)

type TestShape struct {
	transform *Matrix
	material  *Material
	savedRay  *Ray
}

func NewTestShape() *TestShape {
	return &TestShape{
		transform: IdentityMatrix(),
		material:  NewMaterial(),
	}
}

func (s *TestShape) Intersections(ray *Ray) IntersectionSlice {
	s.savedRay = ray
	return make(IntersectionSlice, 0)
}

func (s *TestShape) NormalAt(objectPoint *Tuple) *Tuple {
	return NewVector(
		objectPoint.X,
		objectPoint.Y,
		objectPoint.Z,
	)
}

func (s *TestShape) Material() *Material {
	return s.material
}

func (s *TestShape) SetMaterial(material *Material) {
	s.material = material
}

func (a *TestShape) Equal(b Shape) bool {
	if sb, ok := b.(*TestShape); !ok {
		return false
	} else {
		return a != nil && sb != nil && a.transform.Equal(sb.transform) && a.material.Equal(sb.material)
	}
}

func (s *TestShape) Transform() *Matrix {
	return s.transform
}

func (s *TestShape) SetTransform(transform *Matrix) {
	s.transform = transform
}

func objectEqualMaterialColor(c1name, o1name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c1name)
	}
	if sh1, ok = shapes[o1name]; !ok {
		return fmt.Errorf("Unknown symbol (object) %s", o1name)
	}

	expected := sh1.Material().Color
	got := c1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected color %s = %v; got %v", c1name, expected, got)
	}
	return nil
}

func objectMaterialAmbient(o1name string, scalar float64) error {
	if sh1, ok = shapes[o1name]; !ok {
		return fmt.Errorf("Unknown symbol (object) %s", o1name)
	}
	sh1.Material().Ambient = scalar
	return nil
}

func shape() error {
	shapes["s"] = NewTestShape()
	return nil
}

func shapeEqualTransform(sh1name, m1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape) %s", sh1name)
	}
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}

	expected := m1
	got := sh1.Transform()
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s.transform = %v; got %v", sh1name, expected, got)
	}
	return nil

}

func shapeTransform(sh1name, m1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sh1name)
	}
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", m1name)
	}

	sh1.SetTransform(m1)
	return nil

}

func shapeMaterial(mat1name, sh1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape) %s", sh1name)
	}
	materials[mat1name] = sh1.Material()
	return nil
}

func shapeMaterial2(sh1name, mat1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (sphere) %s", sh1name)
	}
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	sh1.SetMaterial(mat1)
	return nil
}

func shapeEqualMaterial(sh1name, mat1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape) %s", sh1name)
	}
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (matrix) %s", mat1name)
	}

	expected := mat1
	got := sh1.Material()
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s.material = %v; got %v", sh1name, expected, got)
	}
	return nil

}

func shapeNormalAt(t1name, sh1name string, x, y, z float64) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape) %s", sh1name)
	}

	tuples[t1name] = NormalAt(sh1, NewPoint(x, y, z))
	return nil

}

func shapeLocalNormalAt(t1name, sh1name string, x, y, z float64) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape) %s", sh1name)
	}

	tuples[t1name] = sh1.NormalAt(NewPoint(x, y, z))
	return nil

}

func shapeEqualSavedRayOrigin(sh1name string, x, y, z float64) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape) %s", sh1name)
	}

	if tsh, ok := sh1.(*TestShape); !ok {
		return fmt.Errorf("Expected shape %s to be a TestShape", sh1name)
	} else {
		expected := NewPoint(x, y, z)
		got := tsh.savedRay.Origin
		if !got.Equal(expected) {
			return fmt.Errorf("Expected saved_ray.origin = %v; got %v", expected, got)
		}
	}
	return nil
}

func shapeEqualSavedRayDirection(sh1name string, x, y, z float64) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape) %s", sh1name)
	}

	if tsh, ok := sh1.(*TestShape); !ok {
		return fmt.Errorf("Expected shape %s to be a TestShape", sh1name)
	} else {
		expected := NewVector(x, y, z)
		got := tsh.savedRay.Direction
		if !got.Equal(expected) {
			return fmt.Errorf("Expected saved_ray.direction = %v; got %v", expected, got)
		}
	}
	return nil
}

func shapeIntersect(i1name, sh1name, r1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape): %s", sh1name)
	}
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol (ray): %s", r1name)
	}
	intersections[i1name] = Intersections(sh1, r1)
	return nil
}

func shapeLocalIntersect(i1name, sh1name, r1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape): %s", sh1name)
	}
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol (ray): %s", r1name)
	}
	intersections[i1name] = sh1.Intersections(r1)
	return nil
}

func shapeHas(sh1name string, table *godog.Table) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape): %s", sh1name)
	}
	return shapeSettings(sh1, table)
}

func shapeSettings(sh1 Shape, table *godog.Table) error {
	// reTuple := regexp.MustCompile(`^\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`)
	reScalar := regexp.MustCompile(`^(-?\d+(?:\.\d+)?)$`)
	reTransform := regexp.MustCompile(`^(translation)\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`)

	rows := len(table.Rows)
	if rows < 1 {
		return fmt.Errorf("shapeSettings() requires a table with at least one row")
	}

	cols := len(table.Rows[0].Cells)
	if cols != 2 {
		return fmt.Errorf("shapeSettings() requires a table with 2 columns")
	}

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
		case "material.transparency":
			if m := reScalar.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract scalar from %s", value)
			} else {
				var transparency float64
				if transparency, err = strconv.ParseFloat(m[1], 64); err != nil {
					return err
				}
				sh1.Material().Transparency = transparency
			}
		case "material.refractive_index":
			if m := reScalar.FindStringSubmatch(value); m == nil {
				return fmt.Errorf("Unable to extract scalar from %s", value)
			} else {
				var refractiveIndex float64
				if refractiveIndex, err = strconv.ParseFloat(m[1], 64); err != nil {
					return err
				}
				sh1.Material().RefractiveIndex = refractiveIndex
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
					return fmt.Errorf("Unknown transform %s in shapeSettings()", m[1])
				}
			}
		default:
			return fmt.Errorf("Unknown field %s in shapeSettings()", name)
		}
	}

	return nil
}
