package bezierfit

import (
	"math"
	"math/rand"
)

const (
	bestFitAttempts   = 10
	derivativeEpsilon = 1e-4
	descentStepSize   = 0.1
)

// BestFit approximates the BezierAnimation of best
// fit (in a least squares sense) for the list of
// points.
// If two or fewer points are passed, the returned
// animation should pass through the given points.
func BestFit(points []Point) *BezierAnimation {
	var minError float64
	var bestFit *BezierAnimation

	// Perform gradient descent multiple times from
	// multiple starting positions to minimize the
	// chances of getting caught in local minima.
	for i := 0; i < bestFitAttempts; i++ {
		anim := &BezierAnimation{
			P1: Point{rand.Float64(), rand.Float64()},
			P2: Point{rand.Float64(), rand.Float64()},
		}
		gradientDescent(points, anim)
		err := meanSquaredError(points, anim)
		if err < minError || i == 0 {
			minError = err
			bestFit = anim
		}
	}

	return bestFit
}

func gradientDescent(points []Point, anim *BezierAnimation) {
	paramPtrs := []*float64{&anim.P1.X, &anim.P1.Y, &anim.P2.X, &anim.P2.Y}
	gradient := make([]float64, 4)
	oldError := meanSquaredError(points, anim)
	for {
		for i, ptr := range paramPtrs {
			// Using this two-point formula gives O(derivativeEpsilon^2)
			// error instead of the traditional derivative formula's
			// O(derivativeEpsilon) error (and is thus more accurate).
			centerVal := *ptr
			*ptr += derivativeEpsilon
			rightErr := meanSquaredError(points, anim)
			*ptr = centerVal - derivativeEpsilon
			leftErr := meanSquaredError(points, anim)
			*ptr = centerVal
			gradient[i] = (rightErr - leftErr) / (derivativeEpsilon * 2)
		}
		for i, g := range gradient {
			*paramPtrs[i] -= g * descentStepSize
		}
		newErr := meanSquaredError(points, anim)
		if newErr > oldError {
			break
		}
		oldError = newErr
	}
}

func meanSquaredError(points []Point, anim *BezierAnimation) float64 {
	var e float64
	for _, point := range points {
		actual := anim.Eval(point.X)
		e += math.Pow(actual-point.Y, 2)
	}
	return e
}
