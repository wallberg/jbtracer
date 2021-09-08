package jbtracer

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
