package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dylanpinn/FIT3036-backend/area"
)

// our main function
func main() {
	fmt.Println("Listening on port 8080")
	http.HandleFunc("/area", AreaHandler)
	http.HandleFunc("/roadArea", RoadAreaHandler)
	http.ListenAndServe(":8080", nil)
}

// AreaHandler sets up the API to handle Area requests
func AreaHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t area.PointRect
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	area := area.CalculateArea(t)
	js, err := json.Marshal(area)
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

	area := area.CalculateRoadArea(t)
	js, err := json.Marshal(area)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
