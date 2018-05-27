package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dylanpinn/FIT3036-backend/area"
)

// our main function which sets up the route handlers.
func main() {
	fmt.Println("Listening on port 8080")
	http.HandleFunc("/area", AreaHandler)
	http.HandleFunc("/roadArea", RoadAreaHandler)
	http.ListenAndServe(":8080", nil)
}

// Check that the API request contains all of the values.
func isValid(rect area.PointRect) bool {
	if rect.North == 0 {
		return false
	}
	if rect.South == 0 {
		return false
	}
	if rect.East == 0 {
		return false
	}
	if rect.West == 0 {
		return false
	}
	return true
}

// AreaHandler sets up the API to handle Area requests
func AreaHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t area.PointRect
	err := decoder.Decode(&t)

	valid := isValid(t)
	if valid == false {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}

	if err != nil {
		panic(err)
	}

	calcArea := area.CalculateArea(t)
	js, err := json.Marshal(calcArea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// RoadAreaHandler sets up the API to handle Road Area requests.
func RoadAreaHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t area.PointRect
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	valid := isValid(t)
	if valid == false {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}

	calcArea := area.CalculateRoadArea(t)
	js, err := json.Marshal(calcArea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
