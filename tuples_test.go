package jbtracer

import (
	"context"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v16"
	"github.com/spf13/pflag"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

var (
	t1, t2, expected, got *Tuple
	ok                    bool
)
var symbols map[string]*Tuple

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                 "tuples",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
}

func isPoint(t1name string) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if !t1.IsPoint() {
		return fmt.Errorf("Expected %s.isPoint()=true; got false", t1name)
	}
	return nil
}

func isVector(t1name string) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if !t1.IsVector() {
		return fmt.Errorf("Expected %s.isVector()=true; got false", t1name)
	}
	return nil
}

func isNotPoint(t1name string) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.IsPoint() {
		return fmt.Errorf("Expected %s.isPoint()=false; got true", t1name)
	}
	return nil
}

func isNotVector(t1name string) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.IsVector() {
		return fmt.Errorf("Expected %s.isVector()=false; got true", t1name)
	}
	return nil
}

func tuple(t1name string, x, y, z, w float32) error {
	symbols[t1name] = &Tuple{X: x, Y: y, Z: z, W: w}
	return nil
}

func equalsTupleX(t1name string, x float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.X != x {
		return fmt.Errorf("Expected %f for %s.x; got %f", x, t1name, t1.X)
	}
	return nil
}

func equalsTupleY(t1name string, y float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.Y != y {
		return fmt.Errorf("Expected %f for %s.y; got %f", y, t1name, t1.Y)
	}
	return nil
}

func equalsTupleZ(t1name string, z float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.Z != z {
		return fmt.Errorf("Expected %f for %s.z; got %f", z, t1name, t1.Z)
	}
	return nil
}

func equalsTupleW(t1name string, w float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t1.W != w {
		return fmt.Errorf("Expected %f for %s.w; got %f", w, t1name, t1.W)
	}
	return nil
}

func point(t1name string, x, y, z float32) error {
	symbols[t1name] = NewPoint(x, y, z)
	return nil
}

func vector(t1name string, x, y, z float32) error {
	symbols[t1name] = NewVector(x, y, z)
	return nil
}

func equalsTuple(t1name string, x, y, z, w float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	if !t1.Equal(expected) {
		return fmt.Errorf("Expected %s=%v; got %v", t1name, expected, t1)
	}
	return nil
}

func equalsTupleAdd(t1name string, t2name string, x, y, z, w float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = symbols[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	got := t1.Add(t2)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s + %s=%v; got %v", t1name, t2name, expected, got)
	}
	return nil
}

func equalsTupleSubtract(t1name string, t2name string, ttype string, x, y, z float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = symbols[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}
	if ttype == "point" {
		expected = NewPoint(x, y, z)
	} else {
		expected = NewVector(x, y, z)
	}
	got = t1.Subtract(t2)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s + %s=%v; got %v", t1name, t2name, expected, got)
	}
	return nil
}

func equalsTupleNegate(t1name string, x, y, z, w float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected = &Tuple{X: x, Y: y, Z: z, W: w}
	got = t1.Negate()
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s=%v; got %v", t1name, expected, got)
	}
	return nil
}

func equalsTupleMultiply(t1name string, scalar float32, x, y, z, w float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	got := t1.Multiply(scalar)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s + %f=%v; got %v", t1name, scalar, expected, got)
	}
	return nil
}

func equalsTupleDivide(t1name string, scalar float32, x, y, z, w float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	expected := &Tuple{X: x, Y: y, Z: z, W: w}
	got := t1.Divide(scalar)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s + %f=%v; got %v", t1name, scalar, expected, got)
	}
	return nil
}

func equalsTupleMagnitude(t1name string, expected float32) error {
	if t1, ok = symbols[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	got := t1.Magnitude()
	if math.Abs((float64)(got)-(float64)(expected)) >= Epsilon {
		return fmt.Errorf("Expected %s=%f; got %f", t1name, expected, got)
	}
	return nil
}

func equalsTupleMagnitudeSquareRoot(t1name string, expected float32) error {
	expected = (float32)(math.Sqrt((float64)(expected)))
	return equalsTupleMagnitude(t1name, expected)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^(\w+) is a point$`, isPoint)
	ctx.Step(`^(\w+) is a vector$`, isVector)
	ctx.Step(`^(\w+) is not a point$`, isNotPoint)
	ctx.Step(`^(\w+) is not a vector$`, isNotVector)
	ctx.Step(`^(\w+) ← tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tuple)
	ctx.Step(`^(\w+)\.x = (-?\d+(?:\.\d+)?)$`, equalsTupleX)
	ctx.Step(`^(\w+)\.y = (-?\d+(?:\.\d+)?)$`, equalsTupleY)
	ctx.Step(`^(\w+)\.z = (-?\d+(?:\.\d+)?)$`, equalsTupleZ)
	ctx.Step(`^(\w+)\.w = (-?\d+(?:\.\d+)?)$`, equalsTupleW)
	ctx.Step(`^(\w+) ← point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, point)
	ctx.Step(`^(\w+) ← vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, vector)
	ctx.Step(`^(\w+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsTuple)
	ctx.Step(`^(\w+) \+ (\w+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsTupleAdd)
	ctx.Step(`^(\w+) - (\w+) = (point|vector)\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsTupleSubtract)
	ctx.Step(`^-(\w+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsTupleNegate)
	ctx.Step(`^(\w+) \* (-?\d+(?:\.\d+)?) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsTupleMultiply)
	ctx.Step(`^(\w+) / (-?\d+(?:\.\d+)?) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsTupleDivide)
	ctx.Step(`^magnitude\((\w+)\) = (-?\d+(?:\.\d+)?)$`, equalsTupleMagnitude)
	ctx.Step(`^magnitude\((\w+)\) = √(-?\d+(?:\.\d+)?)$`, equalsTupleMagnitudeSquareRoot)

	ctx.Before(func(ctx context.Context, sc *messages.Pickle) (context.Context, error) {

		symbols = make(map[string]*Tuple)

		return ctx, nil
	})
}
