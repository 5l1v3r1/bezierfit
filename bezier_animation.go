package bezierfit

const evalIterations = 30

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
	var minT, maxT float64
	maxT = 1

	// Basic bisection search gives O(evalIterations) precision.
	// If this were replaced with Newton's method, we could get
	// O(evalIterations^2) precision (which is better).
	for i := 0; i < evalIterations; i++ {
		midT := (minT + maxT) * 0.5
		xVal := b.xForT(midT)
		if xVal < x {
			minT = midT
		} else if xVal > x {
			maxT = midT
		} else {
			maxT = midT
			minT = midT
			break
		}
	}

	return b.yForT((minT + maxT) / 2)
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

func cubicBezierParametricEquation(p0, p1, p2, p3, t float64) float64 {
	tComp := 1 - t
	tComp2 := tComp * tComp
	tComp3 := tComp2 * tComp
	t2 := t * t
	t3 := t2 * t
	return tComp3*p0 + 3*tComp2*t*p1 + 3*tComp*t2*p2 + t3*p3
}
