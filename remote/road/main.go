package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dylanpinn/FIT3036-backend/area"
)

func handleRequest(ctx context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Initialise rectangle
	rect := &area.PointRect{}

	// Parse request body
	json.Unmarshal([]byte(request.Body), rect)
	calcArea := area.CalculateRoadArea(*rect)
	body, _ := json.Marshal(calcArea)

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
