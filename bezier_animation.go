package bezierfit

import "math"

const evalXPrecision = 1e-8

type Point struct {
	X float64
	Y float64
}

// A BezierAnimation defines an animation using a
// bezier curve starting from (0,0) and going to
// (1,1) with control points P1 and P2.
// The x coordinates of the control points must be
// bounded in the range [0, 1].
type BezierAnimation struct {
	P1 Point
	P2 Point
}

// Eval approximates the y value for a given x
// value along the bezier animation.
// The given x value should be in the range [0,1].
func (b *BezierAnimation) Eval(x float64) float64 {
	var approx, counter float64
	var approxVal, counterVal float64

	if x < 0.5 {
		counter = 1
		counterVal = 1
	} else {
		approx = 1
		approxVal = 1
	}

	// Use a hybrid of Newton's method and bisection.
	for math.Abs(approxVal-x) > evalXPrecision {
		bisectionPoint := (approx + counter) / 2
		newtonPoint := approx - (b.xForT(approx)-x)/(b.xPrimeForT(approx))

		var usePoint float64
		if math.Abs(approx-newtonPoint) < math.Abs(approx-bisectionPoint) {
			usePoint = newtonPoint
		} else {
			usePoint = bisectionPoint
		}
		useValue := b.xForT(usePoint)

		if (useValue < x) == (approxVal < x) {
			approx = usePoint
			approxVal = useValue
		} else {
			counter = usePoint
			counterVal = useValue
		}

		if math.Abs(approx-x) > math.Abs(counter-x) {
			counter, approx = approx, counter
			counterVal, approxVal = approxVal, counterVal
		}
	}

	return b.yForT(approx)
}

// xForT evaluates the Bezier curve's underlying
// parametric equation for the x coordinate.
func (b *BezierAnimation) xForT(t float64) float64 {
	return cubicBezierParametricEquation(0, b.P1.X, b.P2.X, 1, t)
}

// yForT evaluates the Bezier curve's underlying
// parametric equation for the y coordinate.
func (b *BezierAnimation) yForT(t float64) float64 {
	return cubicBezierParametricEquation(0, b.P1.Y, b.P2.Y, 1, t)
}

// xPrimeForT evaluates the derivative of the
// underlying parametric equation for x with
// respect to t.
func (b *BezierAnimation) xPrimeForT(t float64) float64 {
	return cubicBezierDerivative(0, b.P1.X, b.P2.X, 1, t)
}

func cubicBezierParametricEquation(p0, p1, p2, p3, t float64) float64 {
	tComp := 1 - t
	tComp2 := tComp * tComp
	tComp3 := tComp2 * tComp
	t2 := t * t
	t3 := t2 * t
	return tComp3*p0 + 3*tComp2*t*p1 + 3*tComp*t2*p2 + t3*p3
}

func cubicBezierDerivative(p0, p1, p2, p3, t float64) float64 {
	tComp := 1 - t
	tComp2 := tComp * tComp
	t2 := t * t
	return 3*tComp2*(p1-p0) + 6*tComp*t*(p2-p1) + 3*t2*(p3-p2)
}
