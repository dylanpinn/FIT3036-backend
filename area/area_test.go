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

func TestCalculateArea(t *testing.T) {
	rect := PointRect{North: -37.9072244235794, South: -37.9162075764206, East: 145.13289004553383, West: 145.12150395446622}
	area := CalculateArea(rect)
	const expectedArea = 0.9977023751531141

	if diff := math.Abs(area - expectedArea); diff > TOLERANCE {
		t.Errorf("Area was incorrect, got: %f, want: %f.", area, expectedArea)
	}
}

func TestCalculateRoadArea(t *testing.T) {
	rect := PointRect{North: -37.9072244235794, South: -37.9162075764206, East: 145.13289004553383, West: 145.12150395446622}
	area := CalculateRoadArea(rect)
	const expectedArea = 0.19777044148142264

	if diff := math.Abs(area - expectedArea); diff > TOLERANCE {
		t.Errorf("Area was incorrect, got: %f, want: %f.", area, expectedArea)
	}
}
