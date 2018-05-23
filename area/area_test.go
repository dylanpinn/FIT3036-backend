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

func TestIsPointInsideBounds(t *testing.T) {
	// Outside both
	pt1 := Coordinate{lat: -37.92, long: 145.16}
	rect := PointRect{North: -36.90, South: -37.0, East: 145.10, West: 145.0}
	result := IsPointInsideBounds(pt1, rect)

	if result == true {
		t.Errorf("IsPointInsideBounds was incorrect, got: %t, want: %t", result,
			false)
	}

	// Inside lat
	rect = PointRect{North: -36.90, South: -37.5, East: 145.10, West: 145.0}
	result = IsPointInsideBounds(pt1, rect)

	if result == true {
		t.Errorf("IsPointInsideBounds was incorrect, got: %t, want: %t", result,
			false)
	}

	// Inside long
	rect = PointRect{North: -36.90, South: -37.0, East: 145.20, West: 145.0}
	result = IsPointInsideBounds(pt1, rect)

	if result == true {
		t.Errorf("IsPointInsideBounds was incorrect, got: %t, want: %t", result,
			false)
	}

	// Inside both
	rect = PointRect{North: -36.80, South: -38.0, East: 145.20, West: 145.0}
	result = IsPointInsideBounds(pt1, rect)

	if result == false {
		t.Errorf("IsPointInsideBounds was incorrect, got: %t, want: %t", result,
			true)
	}
}
