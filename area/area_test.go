package area

import (
	"math"
	"testing"
)

const TOLERANCE = 0.00001

func TestCalculateDistance(t *testing.T) {
	pt1 := Coordinate{lat: -37.92, long: 145.16}
	pt2 := Coordinate{lat: -37.92, long: 145.17}
	distance := CalculateDistance(pt1, pt2)

	if diff := math.Abs(distance - 0.877183); diff > TOLERANCE {
		t.Errorf("Distance was incorrect, got: %f, want: %f.", distance, 0.877183)
	}
}
