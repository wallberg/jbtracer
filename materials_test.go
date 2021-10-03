package jbtracer

import "fmt"

func material(m1name string) error {
	materials[m1name] = NewMaterial()
	return nil
}

func materialEqualColor(mat1name string, red, green, blue float64) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", mat1name)
	}
	expected := &Color{red, green, blue}
	got := mat1.Color
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s.color=%v; got %v", mat1name, expected, got)
	}
	return nil
}

func materialEqualAmbient(mat1name string, scalar float64) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	expected := scalar
	got := mat1.Ambient
	if got != expected {
		return fmt.Errorf("Expected %s.Ambient=%v; got %v", mat1name, expected, got)
	}
	return nil
}

func materialEqualDiffuse(mat1name string, scalar float64) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	expected := scalar
	got := mat1.Diffuse
	if got != expected {
		return fmt.Errorf("Expected %s.Diffuse=%v; got %v", mat1name, expected, got)
	}
	return nil
}

func materialEqualSpecular(mat1name string, scalar float64) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	expected := scalar
	got := mat1.Specular
	if got != expected {
		return fmt.Errorf("Expected %s.Specular=%v; got %v", mat1name, expected, got)
	}
	return nil
}

func materialEqualShininess(mat1name string, scalar float64) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	expected := scalar
	got := mat1.Shininess
	if got != expected {
		return fmt.Errorf("Expected %s.Shininess=%v; got %v", mat1name, expected, got)
	}
	return nil
}

func materialAmbient(mat1name string, scalar float64) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	mat1.Ambient = scalar
	return nil
}

func materialEqual(mat1name, mat2name string) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	if mat2, ok = materials[mat2name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat2name)
	}
	expected := mat2
	got := mat1
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s=%v; got %v", mat1name, expected, got)
	}
	return nil
}

func lighting(c1name, mat1name, t1name, t2name, t3name string) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t2name)
	}
	if t3, ok = tuples[t3name]; !ok {
		return fmt.Errorf("Unknown symbol (tuple) %s", t3name)
	}
	colors[c1name] = mat1.Lighting(light, NewSphere(), t1, t2, t3, inShadow)
	return nil
}

func materialInShadow(f string) error {
	if f == "true" {
		inShadow = true
	} else {
		inShadow = false
	}
	return nil
}

func materialDiffuse(mat1name string, scalar float64) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	mat1.Diffuse = scalar
	return nil
}

func materialSpecular(mat1name string, scalar float64) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	mat1.Specular = scalar
	return nil
}

func materialPattern(mat1name string) error {
	if mat1, ok = materials[mat1name]; !ok {
		return fmt.Errorf("Unknown symbol (material) %s", mat1name)
	}
	mat1.Pattern = pattern
	return nil
}
