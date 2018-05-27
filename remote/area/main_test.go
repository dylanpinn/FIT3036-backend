package main

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  int
		err     error
	}{
		{
			request: events.APIGatewayProxyRequest{Body: `{"north": -37.9072244235794, "south": -37.9162075764206, "east":  145.13289004553383, "west":  145.12150395446622}`},
			expect:  200,
			err:     nil,
		},
	}

	for _, test := range tests {
		response, _ := HandleRequest(nil, test.request)

		if status := response.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}
