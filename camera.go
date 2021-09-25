package jbtracer

type Camera struct {
	Hsize       int
	Vsize       int
	FieldOfView float32
	Transform   *Matrix
}

// NewCamera returns a new Camera, with default Transform
func NewCamera(hsize, vsize int, fov float32) *Camera {
	return &Camera{
		Hsize:       hsize,
		Vsize:       vsize,
		FieldOfView: fov,
		Transform:   IdentityMatrix(),
	}
}
