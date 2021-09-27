package jbtracer

import "fmt"

type TestShape struct {
	transform *Matrix
	material  *Material
}

func NewTestShape() *TestShape {
	return &TestShape{
		transform: IdentityMatrix(),
		material:  NewMaterial(),
	}
}

func (s *TestShape) Intersections(ray *Ray) Intersections {
	return make(Intersections, 0)
}

func (s *TestShape) NormalAt(worldPoint *Tuple) *Tuple {
	return nil
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
