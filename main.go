package main

import "fmt"
import "math"

func main() {
	north := -37.93412680920308
	south := -37.83417314312876
	east := 145.1707244873047
	west := 145.0707244873047
	area := CalculateArea(north, south, east, west)
	fmt.Printf("Calculated Area: %f", area)
}

// CalculateArea the area of a Spherical Rectangle
// Taken from http://mathforum.org/library/drmath/view/63767.html
// A = (pi/180)R^2 |sin(lat1)-sin(lat2)| |lon1-lon2|
// R = radius of the Earth https://www.space.com/17638-how-big-is-earth.html
func CalculateArea(lat1, lat2, lon1, lon2 float64) float64 {
	radius := 6372.0
	area := (math.Pi / 180) * math.Pow(radius, 2) *
		math.Abs(math.Sin(lat1)-math.Sin(lat2)) * math.Abs(lon1-lon2)
	return area
}
