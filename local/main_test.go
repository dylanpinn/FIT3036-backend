package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAreaHandler(t *testing.T) {
	var jsonStr = []byte(`{"north": -37.9072244235794, "south": -37.9162075764206, "east":  145.13289004553383, "west":  145.12150395446622}`)

	req, err := http.NewRequest("POST", "/area", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AreaHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestRoadAreaHandler(t *testing.T) {
	var jsonStr = []byte(`{"north": -37.9072244235794, "south": -37.9162075764206, "east":  145.13289004553383, "west":  145.12150395446622}`)

	req, err := http.NewRequest("POST", "/area", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RoadAreaHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
