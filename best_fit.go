package bezierfit

import (
	"math"
	"math/rand"
)

const (
	bestFitAttempts   = 10
	derivativeEpsilon = 1e-4
)

var descentStepSizes = []float64{4, 1, 0.3, 0.1}

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
	var gradient [4]float64
	paramPtrs := []*float64{&anim.P1.X, &anim.P1.Y, &anim.P2.X, &anim.P2.Y}
	oldError := meanSquaredError(points, anim)

	stepSizes := descentStepSizes

	for {
		var oldValues [4]float64
		for i, ptr := range paramPtrs {
			// Using this two-point formula gives O(derivativeEpsilon^2)
			// error instead of the traditional derivative formula's
			// O(derivativeEpsilon) error (and is thus more accurate).
			centerVal := *ptr
			oldValues[i] = centerVal
			*ptr += derivativeEpsilon
			rightErr := meanSquaredError(points, anim)
			*ptr = centerVal - derivativeEpsilon
			leftErr := meanSquaredError(points, anim)
			*ptr = centerVal
			gradient[i] = (rightErr - leftErr) / (derivativeEpsilon * 2)
		}
		var bestError float64
		var bestStepSize float64
		for i, stepSize := range stepSizes {
			for j, g := range gradient {
				*paramPtrs[j] = oldValues[j] - g*stepSize
				if j%2 == 0 {
					*paramPtrs[j] = math.Max(0, math.Min(1, *paramPtrs[j]))
				}
			}
			err := meanSquaredError(points, anim)
			if i == 0 || err < bestError {
				bestError = err
				bestStepSize = stepSize
			}
		}
		if bestError > oldError {
			for i := range paramPtrs {
				*paramPtrs[i] = oldValues[i]
			}
			return
		}
		for i, g := range gradient {
			*paramPtrs[i] = oldValues[i] - g*bestStepSize
			if i%2 == 0 {
				*paramPtrs[i] = math.Max(0, math.Min(1, *paramPtrs[i]))
			}
		}
		oldError = bestError
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
