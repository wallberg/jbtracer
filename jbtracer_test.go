package jbtracer

import (
	"context"
	"os"
	"testing"

	"github.com/cucumber/godog"
	godogcolors "github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v16"
	"github.com/spf13/pflag"
)

var opts = godog.Options{Output: godogcolors.Colored(os.Stdout)}

var (
	t1, t2, expected, got *Tuple
	c1, c2                *Color
	c                     *Canvas
	m1, m2, m3            *Matrix
	ppm                   *PPM
	ok                    bool
	tuples                map[string]*Tuple
	colors                map[string]*Color
	matrices              map[string]*Matrix
	identityMatrix        *Matrix
)

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
	identityMatrix = IdentityMatrix()
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                 "github.com/wallberg/jbtracer",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
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
	ctx.Step(`^(\w+) ← normalize\((\w+)\)$`, normalize)
	ctx.Step(`^normalize\((\w+)\) = (?:approximately )?vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsVectorNormalize)
	ctx.Step(`^dot\((\w+), (\w+)\) = (-?\d+(?:\.\d+)?)$`, equalsVectorDot)
	ctx.Step(`^cross\((\w+), (\w+)\) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsVectorCross)
	ctx.Step(`^(\w+) ← color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, color)
	ctx.Step(`^(\w+) ([+\-*]) (\w+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, equalsColorOp)
	ctx.Step(`^(\w+)\.(red|green|blue) = (-?\d+(?:\.\d+)?)$`, equalsColorField)
	ctx.Step(`^c ← canvas\((\d+), (\d+)\)$`, canvas)
	ctx.Step(`^c\.(width|height) = (\d+)$`, equalsCanvasField)
	ctx.Step(`^every pixel of c is color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, allCanvasColors)
	ctx.Step(`^pixel_at\(c, (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\) = (\w+)$`, pixelAt)
	ctx.Step(`^write_pixel\(c, (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (\w+)\)$`, writePixel)
	ctx.Step(`^ppm ← canvas_to_ppm\(c\)$`, canvasToPPM)
	ctx.Step(`^lines (\d+)-(\d+) of ppm are$`, linesOfPPM)
	ctx.Step(`^every pixel of c is set to color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, assignCanvasAllColors)
	ctx.Step(`^ppm ends with a newline character$`, ppmEndsWithANewlineCharacter)
	ctx.Step(`^the following (?:.+ )?matrix (\w+):$`, matrix)
	ctx.Step(`^(\w+)\[(\d+),(\d+)\] = (-?\d+(?:\.\d+)?)$`, matrixCellEquals)
	ctx.Step(`^(\w+) (!?=) (\w+)$`, matrixEquals)
	ctx.Step(`^(\w+) \* (\w+) = matrix (\w+)$`, matrixMultiply)
	ctx.Step(`^(\w+) \* (\w+) = tuple (\w+)$`, matrixMultiplyTuple)
	ctx.Step(`^(\w+) ← transpose\((\w+)\)$`, matrixTranspose)

	ctx.Before(func(ctx context.Context, sc *messages.Pickle) (context.Context, error) {

		// Reset values before each scenario
		tuples = make(map[string]*Tuple)
		colors = make(map[string]*Color)
		matrices = make(map[string]*Matrix)
		matrices["identity_matrix"] = identityMatrix
		c = nil
		ppm = nil

		return ctx, nil
	})
}
