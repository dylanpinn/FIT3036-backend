package area

import (
	"fmt"
	"strings"

	"github.com/serjvanilla/go-overpass"
	"github.com/umahmood/haversine"
)

// LaneWidth is the Default Lane width in Australia.
const LaneWidth = 0.0035

// PointRect contains the 4 points used to calculate the area of a rectangle.
type PointRect struct {
	North float64
	South float64
	East  float64
	West  float64
}

// Coordinate object containing latitude and longitude.
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

// CalculateRoadArea in a PointRect.
func CalculateRoadArea(t PointRect) float64 {
	client := overpass.New()

	query := buildQuery(t)

	result, _ := client.Query(query)
	sumArea := sumArea(result)
	return sumArea
}

func sumArea(osmData overpass.Result) float64 {
	total := 0.0
	for _, v := range osmData.Ways {
		total += calculateAreaOfWay(v)
	}

	return total
}

func calculateAreaOfWay(way *overpass.Way) float64 {
	wayNodes := way.Nodes
	distance := 0.0
	noOfWays := len(wayNodes)
	for k, v := range wayNodes {
		if k != noOfWays-1 {
			nextNode := wayNodes[k+1]
			pt1 := Coordinate{v.Lat, v.Lon}
			pt2 := Coordinate{nextNode.Lat, nextNode.Lon}
			distance += CalculateDistance(pt1, pt2)
		}
	}

	// TODO: Use supplied distance.
	area := distance * LaneWidth * 2
	return area
}

func buildQuery(t PointRect) string {
	// (lat_min, lon_min, lat_max, lon_max)
	var query strings.Builder

	query.WriteString("[out:json][timeout:25];")
	query.WriteString("(")
	query.WriteString("node['highway']['highway'!='footway']")
	query.WriteString("['highway'!='pedestrian']['-highway'!='path']")
	fmt.Fprintf(&query, "(%f,%f,%f,%f);", t.South, t.West, t.North, t.East)
	query.WriteString("way['highway']['highway'!='footway']")
	query.WriteString("['highway'!='pedestrian']['-highway'!='path']")
	fmt.Fprintf(&query, "(%f,%f,%f,%f);", t.South, t.West, t.North, t.East)
	query.WriteString("relation['highway']['highway'!='footway']")
	query.WriteString("['highway'!='pedestrian']['-highway'!='path']")
	fmt.Fprintf(&query, "(%f,%f,%f,%f);", t.South, t.West, t.North, t.East)
	query.WriteString(");")
	query.WriteString("out body;")
	query.WriteString(">;")
	query.WriteString("out skel qt;")
	return query.String()
}

// CalculateDistance between 2 points.
func CalculateDistance(pt1, pt2 Coordinate) float64 {
	coordinate1 := haversine.Coord{Lat: pt1.lat, Lon: pt1.long}
	coordinate2 := haversine.Coord{Lat: pt2.lat, Lon: pt2.long}
	_, km := haversine.Distance(coordinate1, coordinate2)
	return km
}
