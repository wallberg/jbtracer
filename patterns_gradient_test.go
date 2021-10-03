package jbtracer

import (
	"fmt"
)

func patternGradientPattern(c1name, c2name string) error {
	if c1, ok = colors[c1name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c1name)
	}
	if c2, ok = colors[c2name]; !ok {
		return fmt.Errorf("Unknown symbol (color) %s", c2name)
	}

	pattern = NewGradientPattern(c1, c2)
	return nil
}
