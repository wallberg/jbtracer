package jbtracer

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

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

// func aATuple(arg1, arg2, arg3, arg4, arg5, arg6 int) error {
// 	return godog.ErrPending
// }

// func aIsAPoint() error {
// 	return godog.ErrPending
// }

// func aIsAVector() error {
// 	return godog.ErrPending
// }

// func aIsNotAPoint() error {
// 	return godog.ErrPending
// }

// func aIsNotAVector() error {
// 	return godog.ErrPending
// }

// func aTuple(arg1, arg2, arg3, arg4, arg5, arg6, arg7 int) error {
// 	return godog.ErrPending
// }

// func aTuple(arg1, arg2, arg3, arg4, arg5 int) error {
// 	return godog.ErrPending
// }

// func aTuple(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 int) error {
// 	return godog.ErrPending
// }

// func aTuple(arg1, arg2, arg3, arg4 int) error {
// 	return godog.ErrPending
// }

// func aTuple(arg1, arg2, arg3, arg4 int) error {
// 	return godog.ErrPending
// }

// func aTuple(arg1, arg2, arg3, arg4, arg5 int) error {
// 	return godog.ErrPending
// }

// func aTuple(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 int) error {
// 	return godog.ErrPending
// }

// func aVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func aw(arg1, arg2 int) error {
// 	return godog.ErrPending
// }

// func ax(arg1, arg2 int) error {
// 	return godog.ErrPending
// }

// func ay(arg1, arg2 int) error {
// 	return godog.ErrPending
// }

// func az(arg1, arg2 int) error {
// 	return godog.ErrPending
// }

// func bVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func cCColor(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 int) error {
// 	return godog.ErrPending
// }

// func cCColor(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 int) error {
// 	return godog.ErrPending
// }

// func cCColor(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 int) error {
// 	return godog.ErrPending
// }

// func cColor(arg1, arg2, arg3, arg4, arg5, arg6, arg7 int) error {
// 	return godog.ErrPending
// }

// func cColor(arg1, arg2, arg3, arg4, arg5, arg6, arg7 int) error {
// 	return godog.ErrPending
// }

// func cColor(arg1, arg2, arg3, arg4, arg5, arg6 int) error {
// 	return godog.ErrPending
// }

// func cColor(arg1, arg2, arg3, arg4, arg5, arg6 int) error {
// 	return godog.ErrPending
// }

// func cColor(arg1, arg2, arg3, arg4, arg5, arg6 int) error {
// 	return godog.ErrPending
// }

// func cColor(arg1, arg2, arg3, arg4, arg5, arg6 int) error {
// 	return godog.ErrPending
// }

// func cblue(arg1, arg2 int) error {
// 	return godog.ErrPending
// }

// func cgreen(arg1, arg2 int) error {
// 	return godog.ErrPending
// }

// func cred(arg1, arg2 int) error {
// 	return godog.ErrPending
// }

// func crossaBVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func crossbAVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func dotaB(arg1 int) error {
// 	return godog.ErrPending
// }

// func magnitudenorm(arg1 int) error {
// 	return godog.ErrPending
// }

// func magnitudev(arg1 int) error {
// 	return godog.ErrPending
// }

// func magnitudev(arg1 int) error {
// 	return godog.ErrPending
// }

// func nVector(arg1, arg2, arg3, arg4, arg5 int) error {
// 	return godog.ErrPending
// }

// func nVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func normNormalizev() error {
// 	return godog.ErrPending
// }

// func normalizevApproximatelyVector(arg1, arg2, arg3, arg4, arg5, arg6 int) error {
// 	return godog.ErrPending
// }

// func normalizevVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func pPVector(arg1, arg2, arg3, arg4, arg5 int) error {
// 	return godog.ErrPending
// }

// func pPoint(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func pPoint(arg1, arg2, arg3, arg4 int) error {
// 	return godog.ErrPending
// }

// func pPoint(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func pTuple(arg1, arg2, arg3, arg4 int) error {
// 	return godog.ErrPending
// }

// func pVPoint(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func rReflectvN() error {
// 	return godog.ErrPending
// }

// func rVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func vTuple(arg1, arg2, arg3, arg4 int) error {
// 	return godog.ErrPending
// }

// func vVVector(arg1, arg2, arg3, arg4, arg5 int) error {
// 	return godog.ErrPending
// }

// func vVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func vVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func vVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func vVector(arg1, arg2, arg3, arg4 int) error {
// 	return godog.ErrPending
// }

// func zeroVVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

// func zeroVector(arg1, arg2, arg3 int) error {
// 	return godog.ErrPending
// }

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	// ctx.Step(`^a(\d+) \+ a(\d+) = tuple\((\d+), (\d+), (\d+), (\d+)\)$`, aATuple)
	// ctx.Step(`^a is a point$`, aIsAPoint)
	// ctx.Step(`^a is a vector$`, aIsAVector)
	// ctx.Step(`^a is not a point$`, aIsNotAPoint)
	// ctx.Step(`^a is not a vector$`, aIsNotAVector)
	// ctx.Step(`^a \/ (\d+) = tuple\((\d+)\.(\d+), -(\d+), (\d+)\.(\d+), -(\d+)\)$`, aTuple)
	// ctx.Step(`^a(\d+) ← tuple\((\d+), -(\d+), (\d+), (\d+)\)$`, aTuple)
	// ctx.Step(`^a \* (\d+)\.(\d+) = tuple\((\d+)\.(\d+), -(\d+), (\d+)\.(\d+), -(\d+)\)$`, aTuple)
	// ctx.Step(`^-a = tuple\(-(\d+), (\d+), -(\d+), (\d+)\)$`, aTuple)
	// ctx.Step(`^a ← tuple\((\d+), -(\d+), (\d+), -(\d+)\)$`, aTuple)
	// ctx.Step(`^a(\d+) ← tuple\(-(\d+), (\d+), (\d+), (\d+)\)$`, aTuple)
	// ctx.Step(`^a ← tuple\((\d+)\.(\d+), -(\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, aTuple)
	// ctx.Step(`^a ← vector\((\d+), (\d+), (\d+)\)$`, aVector)
	// ctx.Step(`^a\.w = (\d+)\.(\d+)$`, aw)
	// ctx.Step(`^a\.x = (\d+)\.(\d+)$`, ax)
	// ctx.Step(`^a\.y = -(\d+)\.(\d+)$`, ay)
	// ctx.Step(`^a\.z = (\d+)\.(\d+)$`, az)
	// ctx.Step(`^b ← vector\((\d+), (\d+), (\d+)\)$`, bVector)
	// ctx.Step(`^c(\d+) \* c(\d+) = color\((\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, cCColor)
	// ctx.Step(`^c(\d+) - c(\d+) = color\((\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, cCColor)
	// ctx.Step(`^c(\d+) \+ c(\d+) = color\((\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, cCColor)
	// ctx.Step(`^c \* (\d+) = color\((\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, cColor)
	// ctx.Step(`^c(\d+) ← color\((\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, cColor)
	// ctx.Step(`^c ← color\(-(\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, cColor)
	// ctx.Step(`^c ← color\((\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, cColor)
	// ctx.Step(`^c(\d+) ← color\((\d+)\.(\d+), (\d+), (\d+)\.(\d+)\)$`, cColor)
	// ctx.Step(`^c(\d+) ← color\((\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, cColor)
	// ctx.Step(`^c\.blue = (\d+)\.(\d+)$`, cblue)
	// ctx.Step(`^c\.green = (\d+)\.(\d+)$`, cgreen)
	// ctx.Step(`^c\.red = -(\d+)\.(\d+)$`, cred)
	// ctx.Step(`^cross\(a, b\) = vector\(-(\d+), (\d+), -(\d+)\)$`, crossaBVector)
	// ctx.Step(`^cross\(b, a\) = vector\((\d+), -(\d+), (\d+)\)$`, crossbAVector)
	// ctx.Step(`^dot\(a, b\) = (\d+)$`, dotaB)
	// ctx.Step(`^magnitude\(norm\) = (\d+)$`, magnitudenorm)
	// ctx.Step(`^magnitude\(v\) = (\d+)$`, magnitudev)
	// ctx.Step(`^magnitude\(v\) = √(\d+)$`, magnitudev)
	// ctx.Step(`^n ← vector\(√(\d+)\/(\d+), √(\d+)\/(\d+), (\d+)\)$`, nVector)
	// ctx.Step(`^n ← vector\((\d+), (\d+), (\d+)\)$`, nVector)
	// ctx.Step(`^norm ← normalize\(v\)$`, normNormalizev)
	// ctx.Step(`^normalize\(v\) = approximately vector\((\d+)\.(\d+), (\d+)\.(\d+), (\d+)\.(\d+)\)$`, normalizevApproximatelyVector)
	// ctx.Step(`^normalize\(v\) = vector\((\d+), (\d+), (\d+)\)$`, normalizevVector)
	// ctx.Step(`^p(\d+) - p(\d+) = vector\(-(\d+), -(\d+), -(\d+)\)$`, pPVector)
	// ctx.Step(`^p ← point\((\d+), -(\d+), (\d+)\)$`, pPoint)
	// ctx.Step(`^p(\d+) ← point\((\d+), (\d+), (\d+)\)$`, pPoint)
	// ctx.Step(`^p ← point\((\d+), (\d+), (\d+)\)$`, pPoint)
	// ctx.Step(`^p = tuple\((\d+), -(\d+), (\d+), (\d+)\)$`, pTuple)
	// ctx.Step(`^p - v = point\(-(\d+), -(\d+), -(\d+)\)$`, pVPoint)
	// ctx.Step(`^r ← reflect\(v, n\)$`, rReflectvN)
	// ctx.Step(`^r = vector\((\d+), (\d+), (\d+)\)$`, rVector)
	// ctx.Step(`^v = tuple\((\d+), -(\d+), (\d+), (\d+)\)$`, vTuple)
	// ctx.Step(`^v(\d+) - v(\d+) = vector\(-(\d+), -(\d+), -(\d+)\)$`, vVVector)
	// ctx.Step(`^v ← vector\((\d+), -(\d+), (\d+)\)$`, vVector)
	// ctx.Step(`^v ← vector\((\d+), (\d+), (\d+)\)$`, vVector)
	// ctx.Step(`^v ← vector\(-(\d+), -(\d+), -(\d+)\)$`, vVector)
	// ctx.Step(`^v(\d+) ← vector\((\d+), (\d+), (\d+)\)$`, vVector)
	// ctx.Step(`^zero - v = vector\(-(\d+), (\d+), -(\d+)\)$`, zeroVVector)
	// ctx.Step(`^zero ← vector\((\d+), (\d+), (\d+)\)$`, zeroVector)
}
