package jbtracer

import "math"

type Material struct {
	Color     *Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() *Material {
	return &Material{
		Color:     White,
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200,
	}
}

func (a *Material) Equal(b *Material) bool {
	if a == nil || b == nil {
		return false
	} else if !a.Color.Equal(b.Color) {
		return false
	} else if !EqualFloat64(a.Ambient, b.Ambient) {
		return false
	} else if !EqualFloat64(a.Diffuse, b.Diffuse) {
		return false
	} else if !EqualFloat64(a.Specular, b.Specular) {
		return false
	} else if !EqualFloat64(a.Shininess, b.Shininess) {
		return false
	}
	return true
}

// Lighting uses the Phong Reflection Model to simulate the lighting on a single
// point on the Material
func (material *Material) Lighting(light *PointLight, point, eyev, normalv *Tuple, inShadow bool) *Color {

	// combine the surface color with the light's color/intensity
	effectiveColor := material.Color.Multiply(light.Intensity)

	// find the direction to the light source
	lightv := light.Position.Subtract(point).Normalize()

	// compute the ambient contribution
	ambient := effectiveColor.MultiplyScalar(material.Ambient)

	// if we are in shadow diffuse and specular do not contribute
	if inShadow {
		return ambient
	}

	// light_dot_normal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the
	// light is on the other side of the surface; again diffuse and Specular
	// do not contribute
	lightDotNormal := lightv.Dot(normalv)
	if lightDotNormal < 0 {
		return ambient
	}

	// compute the diffuse contribution
	diffuse := effectiveColor.MultiplyScalar(material.Diffuse).MultiplyScalar(lightDotNormal)

	// reflect_dot_eye represents the cosine of the angle between the
	// reflection vector and the eye vector. A negative number means the
	// light reflects away from the eye; specular does not contribute.
	reflectv := lightv.Multiply(-1).Reflect(normalv)
	reflectDotEye := reflectv.Dot(eyev)
	if reflectDotEye <= 0 {
		return ambient.Add(diffuse)
	}

	// compute the specular contribution
	factor := math.Pow(reflectDotEye, material.Shininess)
	specular := light.Intensity.MultiplyScalar(material.Specular).MultiplyScalar(factor)

	// Add the three contributions together to get the final shading
	return ambient.Add(diffuse).Add(specular)
}
