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

var opts = godog.Options{
	Output: godogcolors.Colored(os.Stdout),
	Format: "progress",
}

var (
	t1, t2, t3, expected, got *Tuple
	c1, c2                    *Color
	c                         *Canvas
	m1, m2, m3                *Matrix
	s1, s2                    float64
	r1, r2                    *Ray
	sph1                      *Sphere
	i1, i2, i3, i4, i5        IntersectionSlice
	sh1                       Shape
	mat1, mat2                *Material
	ppm                       *PPM
	ok                        bool
	tuples                    map[string]*Tuple
	colors                    map[string]*Color
	matrices                  map[string]*Matrix
	identityMatrix            *Matrix
	scalars                   map[string]float64
	rays                      map[string]*Ray
	spheres                   map[string]*Sphere
	intersections             map[string][]*Intersection
	shapes                    map[string]Shape
	light                     *PointLight
	materials                 map[string]*Material
	w                         *World
	comps                     *PreparedComputations
	cam                       *Camera
	inShadow                  bool
	pattern                   Pattern
	stripe                    *StripePattern
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
	// tuples
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
	ctx.Step(`^pixel_at\(c, (\d+), (\d+)\) = (\w+)$`, pixelAt)
	ctx.Step(`^write_pixel\(c, (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (\w+)\)$`, writePixel)
	ctx.Step(`^ppm ← canvas_to_ppm\(c\)$`, canvasToPPM)
	ctx.Step(`^lines (\d+)-(\d+) of ppm are$`, linesOfPPM)
	ctx.Step(`^every pixel of c is set to color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, assignCanvasAllColors)
	ctx.Step(`^ppm ends with a newline character$`, ppmEndsWithANewlineCharacter)
	ctx.Step(`^(\w+) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, vectorEqual)
	ctx.Step(`^(\w+) = normalize\((\w+)\)$`, tupleEqualNormalize)
	ctx.Step(`^(\w+) ← reflect\((\w+), (\w+)\)$`, vectorReflect)
	ctx.Step(`^(\w+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, colorEqual)
	ctx.Step(`^(\w+) = tuple (\w+)$`, tupleEqual)
	ctx.Step(`^(\w+) = color (\w+)$`, colorEqual2)

	// matrices
	ctx.Step(`^the following (?:.+ )?matrix (\w+):$`, matrix)
	ctx.Step(`^(\w+)\[(\d+),(\d+)\] = (-?\d+(?:\.\d+)?)$`, matrixCellEqual)
	ctx.Step(`^(\w+) (!?=) matrix (\w+)$`, matrixEqual)
	ctx.Step(`^(\w+) \* (\w+) = matrix (\w+)$`, matrixMultiply)
	ctx.Step(`^(\w+) \* (\w+) = tuple (\w+)$`, matrixMultiplyTuple)
	ctx.Step(`^(\w+) ← (\w+) \* (\w+)$`, matrixMultiplyAssign)
	ctx.Step(`^(\w+) ← transpose\((\w+)\)$`, matrixTranspose)
	ctx.Step(`^(\w+) ← determinant\((\w+)\)$`, matrixDeterminant)
	ctx.Step(`^(\w+) = scalar ((?:\w|\-)+)$`, scalarEqual)
	ctx.Step(`^(\w+) ← scalar\((-?\d+(?:\.\d+)?)\)$`, scalar)
	ctx.Step(`^(\w+) ← submatrix\((\w+), (\d+), (\d+)\)$`, matrixSubmatrix)
	ctx.Step(`^(\w+) ← minor\((\w+), (\d+), (\d+)\)$`, matrixMinor)
	ctx.Step(`^(\w+) ← cofactor\((\w+), (\d+), (\d+)\)$`, matrixCofactor)
	ctx.Step(`^(\w+) (is(?: not)?) invertible$`, matrixInvertible)
	ctx.Step(`^(\w+) ← inverse\((\w+)\)$`, matrixInverse)

	// transformations
	ctx.Step(`^(\w+) ← translation\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, matrixTranslation)
	ctx.Step(`^(\w+) ← scaling\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, matrixScaling)
	ctx.Step(`^(\w+) ← rotation_([xyz])\((-?\d+(?:\.\d+)?)\)$`, matrixRotation)
	ctx.Step(`^(\w+) ← shearing\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, matrixShearing)
	ctx.Step(`^(\w+) ← (\w+) \* tuple (\w+)$`, tupleMatrixAssign)
	ctx.Step(`^(\w+) ← view_transform\((\w+), (\w+), (\w+)\)$`, matrixViewTransform)

	// rays
	ctx.Step(`^(\w+) ← ray\((\w+), (\w+)\)$`, ray)
	ctx.Step(`^(\w+)\.(origin|direction) = (\w+)$`, rayEqualField)
	ctx.Step(`^position\((\w+), (-?\d+(?:\.\d+)?)\) = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, rayPositionEqualPoint)
	ctx.Step(`^(\w+) ← transform\((\w+), (\w+)\)$`, transform)
	ctx.Step(`^([a-zA-Z0-9_]+)\.direction = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, rayEqualDirectionVector)
	ctx.Step(`^([a-zA-Z0-9_]+)\.origin = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, rayEqualOriginPoint)

	// spheres
	ctx.Step(`^(\w+) ← ray\(point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\), vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, rayPointVector)
	ctx.Step(`^(\w+) ← sphere\(\)$`, sphere)
	ctx.Step(`^(\w+) ← glass_sphere\(\)$`, sphereGlass)
	ctx.Step(`^(\w+)\.count = (\d+)$`, intersectionCount)
	ctx.Step(`^(\w+)\[(\d+)\].object = (\w+)$`, intersectionsObject)
	ctx.Step(`^(\w+) ← sphere\(\) with:$`, sphereWith)

	// intersections
	ctx.Step(`^(\w+) ← intersection\((-?\d+(?:\.\d+)?), (\w+)\)$`, intersection)
	ctx.Step(`^(\w+) ← intersections\((\w+), (\w+)\)$`, intersectionConcat)
	ctx.Step(`^(\w+) ← intersections\((\w+), (\w+), (\w+), (\w+)\)$`, intersectionConcat4)
	ctx.Step(`^(\w+)\[(\d+)\](?:\.t)? = (-?\d+(?:\.\d+)?)$`, intersectionsT)
	ctx.Step(`^(\w+) ← hit\((\w+)\)$`, intersectionHits)
	ctx.Step(`^(\w+) = intersection (\w+)$`, intersectionEqual)
	ctx.Step(`^(\w+) is nothing$`, intersectionEmpty)
	ctx.Step(`^(\w+)\.object = (\w+)$`, intersectionObject)
	ctx.Step(`^(\w+)\.t = (-?\d+(?:\.\d+)?)$`, intersectionT)
	ctx.Step(`^comps ← prepare_computations\((\w+), (\w+)\)$`, comp)
	ctx.Step(`^comps\.object = (\w+)\.object$`, compEqualObject)
	ctx.Step(`^comps\.t = (\w+)\.t$`, compEqualT)
	ctx.Step(`^comps\.point = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, compEqualPoint)
	ctx.Step(`^comps\.eyev = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, compEqualEyeV)
	ctx.Step(`^comps\.normalv = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, compEqualNormalV)
	ctx.Step(`^comps\.inside = (true|false)$`, compEqualInside)
	ctx.Step(`^comps\.over_point\.z < -EPSILON\/2$`, compOverPointZLessThanEpsilon)
	ctx.Step(`^comps\.point\.z > comps\.over_point\.z$`, compPointZGreaterThanOverPointZ)
	ctx.Step(`^comps\.reflectv = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, compEqualReflectV)

	// lights
	ctx.Step(`^light ← point_light\((\w+), (\w+)\)$`, pointLight)
	ctx.Step(`^light\.intensity = (\w+)$`, pointLightIntensity)
	ctx.Step(`^light\.position = (\w+)$`, pointLightPosition)
	ctx.Step(`^light ← point_light\(point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\), color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, pointLight2)

	// materials
	ctx.Step(`^(\w+) ← material\(\)$`, material)
	ctx.Step(`^(\w+)\.color = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, materialEqualColor)
	ctx.Step(`^(\w+)\.ambient = (-?\d+(?:\.\d+)?)$`, materialEqualAmbient)
	ctx.Step(`^(\w+)\.diffuse = (-?\d+(?:\.\d+)?)$`, materialEqualDiffuse)
	ctx.Step(`^(\w+)\.shininess = (-?\d+(?:\.\d+)?)$`, materialEqualShininess)
	ctx.Step(`^(\w+)\.specular = (-?\d+(?:\.\d+)?)$`, materialEqualSpecular)
	ctx.Step(`^(\w+)\.reflective = (-?\d+(?:\.\d+)?)$`, materialEqualReflective)
	ctx.Step(`^(\w+)\.transparency = (-?\d+(?:\.\d+)?)$`, materialEqualTransparency)
	ctx.Step(`^(\w+)\.refractive_index = (-?\d+(?:\.\d+)?)$`, materialEqualRefractiveIndex)
	ctx.Step(`^(\w+) = material (\w+)$`, materialEqual)
	ctx.Step(`^(\w+)\.ambient ← (-?\d+(?:\.\d+)?)$`, materialAmbient)
	ctx.Step(`^(\w+) ← lighting\((\w+), light, (\w+), (\w+), (\w+), in_shadow\)$`, lighting)
	ctx.Step(`^in_shadow ← (true|false)$`, materialInShadow)
	ctx.Step(`^(\w+)\.diffuse ← (-?\d+(?:\.\d+)?)$`, materialDiffuse)
	ctx.Step(`^(\w+)\.pattern ← pattern$`, materialPattern)
	ctx.Step(`^(\w+)\.specular ← (-?\d+(?:\.\d+)?)$`, materialSpecular)

	// world
	ctx.Step(`^w ← world\(\)$`, world)
	ctx.Step(`^w contains no objects$`, worldContainsNoObjects)
	ctx.Step(`^w has no light source$`, worldHasNoLightSource)
	ctx.Step(`^w ← default_world\(\)$`, worldDefault)
	ctx.Step(`^w contains (\w+)$`, worldContainsSphere)
	ctx.Step(`^w\.light = light$`, worldLight)
	ctx.Step(`^(\w+) ← intersect_world\(w, (\w+)\)$`, worldIntersect)
	ctx.Step(`^(\w+) ← shade_hit\(w, comps\)$`, worldShadeHit)
	ctx.Step(`^(\w+) ← the (first|second) object in w$`, worldToObject)
	ctx.Step(`^w\.light ← point_light\(point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\), color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, worldPointLight)
	ctx.Step(`^(\w+) ← color_at\(w, (\w+)\)$`, worldColorAt)
	ctx.Step(`^(\w+) = (\w+)\.material\.color$`, objectEqualMaterialColor)
	ctx.Step(`^(\w+)\.material\.ambient ← (-?\d+(?:\.\d+)?)$`, objectMaterialAmbient)
	ctx.Step(`^is_shadowed\(w, (\w+)\) is (true|false)$`, worldIsShadowed)
	ctx.Step(`^(\w+) is added to w$`, worldAddObject)
	ctx.Step(`^(\w+) ← reflected_color\(w, comps\)$`, worldReflectedColor)
	ctx.Step(`^(\w+) ← reflected_color\(w, comps, (\d+)\)$`, worldReflectedColorDepth)
	ctx.Step(`^color_at\(w, (\w+)\) should terminate successfully$`, worldColorAtTerminates)

	// camera
	ctx.Step(`^c ← camera\((\d+), (\d+), (-?\d+(?:\.\d+)?)\)$`, camera)
	ctx.Step(`^c\.hsize = (\d+)$`, cameraEqualHsize)
	ctx.Step(`^c\.vsize = (\d+)$`, cameraEqualVsize)
	ctx.Step(`^c\.field_of_view = (-?\d+(?:\.\d+)?)$`, cameraEqualFOV)
	ctx.Step(`^c\.transform = (\w+)$`, cameraEqualTransform)
	ctx.Step(`^c\.pixel_size = (-?\d+(?:\.\d+)?)$`, cameraEqualPixelSize)
	ctx.Step(`^c\.transform ← (\w+)$`, cameraTransform)
	ctx.Step(`^(\w+) ← ray_for_pixel\(c, (\d+), (\d+)\)$`, cameraRayForPixel)
	ctx.Step(`^c ← render\(c, w\)$`, render)

	// shapes
	ctx.Step(`^s ← test_shape\(\)$`, shape)
	ctx.Step(`^(s)\.transform = (\w+)$`, shapeEqualTransform)
	ctx.Step(`^set_transform\((\w+), (\w+)\)$`, shapeTransform)
	ctx.Step(`^(\w+) ← (\w+)\.material$`, shapeMaterial)
	ctx.Step(`^(\w+)\.material ← (\w+)$`, shapeMaterial2)
	ctx.Step(`^(\w+)\.material = (\w+)$`, shapeEqualMaterial)
	ctx.Step(`^(\w+) ← normal_at\((\w+), point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, shapeNormalAt)
	ctx.Step(`^(s)\.saved_ray\.origin = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, shapeEqualSavedRayOrigin)
	ctx.Step(`^(s)\.saved_ray\.direction = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, shapeEqualSavedRayDirection)
	ctx.Step(`^(\w+) ← intersect\((\w+), (\w+)\)$`, shapeIntersect)
	ctx.Step(`^(\w+) ← local_normal_at\((\w+), point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, shapeLocalNormalAt)
	ctx.Step(`^(\w+) ← local_intersect\((\w+), (\w+)\)$`, shapeLocalIntersect)

	// planes
	ctx.Step(`^(\w+) ← plane\(\)$`, plane)
	ctx.Step(`^(\w+) ← plane\(\) with:$`, planeWith)

	// patterns
	ctx.Step(`^pattern ← stripe_pattern\((\w+), (\w+)\)$`, patternStripePattern)
	ctx.Step(`^pattern\.a = (\w+)$`, patternEqualA)
	ctx.Step(`^pattern\.b = (\w+)$`, patternEqualB)
	ctx.Step(`^stripe_at\(pattern, point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\) = (\w+)$`, patternEqualStripeAt)
	ctx.Step(`^(\w+) ← stripe_at_object\(pattern, (\w+), (\w+)\)$`, patternPatternAtShape)
	ctx.Step(`^set_pattern_transform\(pattern, (\w+)\)$`, patternSetPatternTransform)
	ctx.Step(`^pattern ← test_pattern\(\)$`, patternTestPattern)
	ctx.Step(`^pattern\.transform = (\w+)$`, patternEqualTransform)
	ctx.Step(`^(\w+) ← pattern_at_shape\(pattern, (\w+), (\w+)\)$`, patternPatternAtShape)
	ctx.Step(`^pattern ← gradient_pattern\((\w+), (\w+)\)$`, patternGradientPattern)
	ctx.Step(`^pattern_at\(pattern, point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\) = (\w+)$`, patternEqualPatternAt)
	ctx.Step(`^pattern ← ring_pattern\((\w+), (\w+)\)$`, patternRingPattern)
	ctx.Step(`^pattern ← checkers_pattern\((\w+), (\w+)\)$`, patternCheckersPattern)

	// Executed before each scenario
	ctx.Before(func(ctx context.Context, sc *messages.Pickle) (context.Context, error) {

		// Reset all values
		tuples = make(map[string]*Tuple)
		colors = make(map[string]*Color)
		matrices = make(map[string]*Matrix)
		matrices["identity_matrix"] = identityMatrix
		scalars = make(map[string]float64)
		rays = make(map[string]*Ray)
		spheres = make(map[string]*Sphere)
		intersections = make(map[string][]*Intersection)
		shapes = make(map[string]Shape)
		light = nil
		c = nil
		ppm = nil
		materials = make(map[string]*Material)
		w = nil
		cam = nil
		inShadow = false
		pattern = nil
		stripe = nil

		return ctx, nil
	})
}

func BenchmarkFloatOps(b *testing.B) {

	b.Run("FloatOps", func(b *testing.B) {

		// world settings
		rayOrigin := NewPoint(0, 0, -5)
		var wallZ float64 = 10
		var wallSize float64 = 7
		wallZHalf := wallSize / 2
		color := &Color{Red: 1, Green: 0, Blue: 0}

		// canvas settings
		canvasPixels := 1000
		pixelSize := wallSize / float64(canvasPixels)
		c := NewCanvas(canvasPixels, canvasPixels)

		sphere := NewSphere()
		transform := Rotation(Axis_Y, 0.78539)
		transform = transform.Multiply(Scaling(0.4, 1, 1))
		transform = transform.Multiply(Translation(0.4, 0, 0))
		sphere.SetTransform(transform)

		// Iterate over canvas points
		for y := 0; y < canvasPixels; y++ {
			worldY := wallZHalf - pixelSize*float64(y)

			for x := 0; x < canvasPixels; x++ {
				worldX := -1*wallZHalf + pixelSize*float64(x)

				// Create a ray from the light source to the canvas point
				position := NewPoint(worldX, worldY, wallZ)
				vector := position.Subtract(rayOrigin).Normalize()
				ray := NewRay(rayOrigin, vector)

				// Determine if the ray intersects the sphere
				var xs IntersectionSlice = sphere.Intersections(ray)
				if hit := xs.Hit(); hit != nil {
					c.Grid[x][y] = color
				}
			}
		}

	})
}
