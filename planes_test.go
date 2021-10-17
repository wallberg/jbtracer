package jbtracer

import (
	"github.com/cucumber/godog"
)

func plane(sh1name string) {
	shapes[sh1name] = NewPlane()
}

func planeWith(sh1name string, table *godog.Table) error {

	sh1 = NewPlane()
	if err := shapeSettings(sh1, table); err != nil {
		return nil
	}

	shapes[sh1name] = sh1
	return nil
}
