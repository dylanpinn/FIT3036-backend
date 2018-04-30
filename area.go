package main

import "github.com/umahmood/haversine"

// PointRect contains the 4 points used to calculate the area of a rectangle.
type PointRect struct {
	North float64
	South float64
	East  float64
	West  float64
}

// Coordinate object containing lattitude and longtude.
type Coordinate struct {
	lat  float64
	long float64
}

// CalculateArea in a PointRect.
func CalculateArea(t PointRect) float64 {
	pt1 := Coordinate{lat: t.North, long: t.West}
	pt2 := Coordinate{lat: t.South, long: t.West}
	distance1 := CalculateDistance(pt1, pt2)
	pt3 := Coordinate{lat: t.South, long: t.East}
	distance2 := CalculateDistance(pt2, pt3)
	area := distance1 * distance2
	return area
}

// CalculateDistance between 2 points.
func CalculateDistance(pt1, pt2 Coordinate) float64 {
	coord1 := haversine.Coord{Lat: pt1.lat, Lon: pt1.long}
	coord2 := haversine.Coord{Lat: pt2.lat, Lon: pt2.long}
	_, km := haversine.Distance(coord1, coord2)
	return km
}
