package jbtracer

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v16"
	"github.com/spf13/pflag"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

var (
	a, p, v *Tuple
)

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
}

func aIsAPoint() error {
	if a.W != 1.0 {
		return fmt.Errorf("Expected tuple a=%v is a point; got not a point", a)
	}
	return nil
}

func aIsAVector() error {
	if a.W != 0.0 {
		return fmt.Errorf("Expected tuple a=%v is a vector; got not a vector", a)
	}
	return nil
}

func aIsNotAPoint() error {
	if a.W == 1.0 {
		return fmt.Errorf("Expected tuple a=%v is not a point; got a point", a)
	}
	return nil
}

func aIsNotAVector() error {
	if a.W == 0.0 {
		return fmt.Errorf("Expected tuple a=%v is not a vector; got a vector", a)
	}
	return nil
}

func aTuple(x, y, z, w float32) error {
	if w != 0.0 && w != 1.0 {
		return fmt.Errorf("Expected w to be 0.0 or 1.0; got %f", w)
	}
	a = &Tuple{X: x, Y: y, Z: z, W: w}
	return nil
}

func aw(w float32) error {
	a.W = w
	return nil
}

func ax(x float32) error {
	a.X = x
	return nil
}

func ay(y float32) error {
	a.Y = y
	return nil
}

func az(z float32) error {
	a.Z = z
	return nil
}

func pPoint(x, y, z float32) error {
	p = NewPoint(x, y, z)
	return nil
}

func pTuple(x, y, z, w float32) error {
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	if !p.Equal(expected) {
		return fmt.Errorf("Expected p=%v; got %v", expected, p)
	}
	return nil
}

func vVector(x, y, z float32) error {
	v = NewVector(x, y, z)
	return nil
}

func vTuple(x, y, z, w float32) error {
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	if !v.Equal(expected) {
		return fmt.Errorf("Expected p=%v; got %v", expected, v)
	}
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a is a point$`, aIsAPoint)
	ctx.Step(`^a is a vector$`, aIsAVector)
	ctx.Step(`^a is not a point$`, aIsNotAPoint)
	ctx.Step(`^a is not a vector$`, aIsNotAVector)
	ctx.Step(`^a ← tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, aTuple)
	ctx.Step(`^a\.w = (-?\d+(?:\.\d+)?)$`, aw)
	ctx.Step(`^a\.x = (-?\d+(?:\.\d+)?)$`, ax)
	ctx.Step(`^a\.y = (-?\d+(?:\.\d+)?)$`, ay)
	ctx.Step(`^a\.z = (-?\d+(?:\.\d+)?)$`, az)
	ctx.Step(`^p ← point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, pPoint)
	ctx.Step(`^p = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, pTuple)
	ctx.Step(`^v = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, vTuple)
	ctx.Step(`^v ← vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, vVector)

	ctx.Before(func(ctx context.Context, sc *messages.Pickle) (context.Context, error) {
		a = nil
		p = nil
		v = nil

		return ctx, nil
	})
}
