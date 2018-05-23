package area

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/serjvanilla/go-overpass"
	"github.com/umahmood/haversine"
)

// LaneWidth is the Default Lane width in Australia in km's.
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
	sumArea := sumArea(result, t)
	return sumArea
}

func sumArea(osmData overpass.Result, rect PointRect) float64 {
	total := 0.0
	for _, v := range osmData.Ways {
		total += calculateAreaOfWay(v, rect)
	}

	return total
}

func calculateAreaOfWay(way *overpass.Way, rect PointRect) float64 {
	lanes := 2
	noOfLanes := way.Meta.Tags["lanes"]
	if l, err := strconv.Atoi(noOfLanes); err == nil {
		lanes = l * 2
	}
	wayNodes := way.Nodes
	distance := 0.0
	noOfWays := len(wayNodes)
	for k, v := range wayNodes {
		if k != noOfWays-1 {
			nextNode := wayNodes[k+1]
			pt1 := Coordinate{v.Lat, v.Lon}
			fmt.Println(IsPointInsideBounds(pt1, rect), pt1, rect)
			if !IsPointInsideBounds(pt1, rect) {
				fmt.Println("Outside", pt1)
				continue
			}
			pt2 := Coordinate{nextNode.Lat, nextNode.Lon}
			fmt.Println(IsPointInsideBounds(pt2, rect), pt2, rect)
			if !IsPointInsideBounds(pt2, rect) {
				fmt.Println("Outside", pt2)
				continue
			}
			d := CalculateDistance(pt1, pt2)
			// distance += CalculateDistance(pt1, pt2)
			fmt.Println(d)
			distance += d
		}
	}

	fmt.Println(distance, LaneWidth, lanes)
	area := distance * LaneWidth * float64(lanes)
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

// CalculateDistance returns km's between 2 points.
func CalculateDistance(pt1, pt2 Coordinate) float64 {
	coordinate1 := haversine.Coord{Lat: pt1.lat, Lon: pt1.long}
	coordinate2 := haversine.Coord{Lat: pt2.lat, Lon: pt2.long}
	_, km := haversine.Distance(coordinate1, coordinate2)
	return km
}

// check if lat is less than east and bigger than west
// check if long is less than north and bigger than south
func IsPointInsideBounds(coor Coordinate, rect PointRect) bool {
	longBetween := InBetween(coor.long, rect.West, rect.East)
	latBetween := InBetween(coor.lat, rect.South, rect.North)
	if longBetween && latBetween {
		return true
	}
	return false
}

// InBetween checks if float64 is between 2 others.
func InBetween(i, min, max float64) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
