package main

import (
	"context"
	"github.com/dylanpinn/FIT3036-backend/area"

	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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
