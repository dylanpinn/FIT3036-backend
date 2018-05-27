package main

import (
	"context"

	"github.com/dylanpinn/FIT3036-backend/area"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var headers map[string]string

func HandleRequest(ctx context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Initialise rectangle
	rect := &area.PointRect{}

	// Parse request body
	json.Unmarshal([]byte(request.Body), rect)
	calcArea := area.CalculateArea(*rect)
	body, _ := json.Marshal(calcArea)

	headers = make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200,
		Headers: headers}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
