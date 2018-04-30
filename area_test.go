package main

import "testing"

func TestCalculateDistance(t *testing.T) {
	pt1 := Coordinate{lat: -37.92, long: 145.16}
	pt2 := Coordinate{lat: -37.92, long: 145.17}
	distance := CalculateDistance(pt1, pt2)

	if distance != 0.877183 {
		t.Errorf("Distance was incorrect, got: %f, want: %f.", distance, 0.877183)
	}
}
