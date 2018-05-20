package main

import (
	"context"
	"fmt"
	"github.com/dylanpinn/FIT3036-backend/area"

	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n",
		request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	// Initialise rectangle
	rect := &area.PointRect{}

	// Parse request body
	json.Unmarshal([]byte(request.Body), rect)
	area := area.CalculateArea(*rect)
	body, _ := json.Marshal(area)

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
