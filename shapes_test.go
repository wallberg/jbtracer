package jbtracer

import "fmt"

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

func intersect(i1name, sh1name, r1name string) error {
	if sh1, ok = shapes[sh1name]; !ok {
		return fmt.Errorf("Unknown symbol (shape): %s", sh1name)
	}
	if r1, ok = rays[r1name]; !ok {
		return fmt.Errorf("Unknown symbol (ray): %s", r1name)
	}
	intersections[i1name] = Intersections(sh1, r1)
	return nil
}
