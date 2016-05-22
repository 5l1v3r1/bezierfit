package bezierfit

import (
	"math"
	"testing"
)

func TestBestFit(t *testing.T) {
	p1 := Point{0.3, 0.2}
	p2 := Point{0.7, 0.8}
	bestFit := BestFit([]Point{p1, p2})
	for _, point := range []Point{p1, p2} {
		actual := bestFit.Eval(point.X)
		if math.Abs(actual-point.Y) > 1e-4 {
			t.Errorf("expected bestFit(%f) = %f but got %f", point.X, point.Y,
				actual)
		}
	}
}
