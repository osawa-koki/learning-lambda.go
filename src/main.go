package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var invokeCount = 0

func init() {
	invokeCount = 100
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	invokeCount++
	path := request.Path
	method := request.HTTPMethod
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("path: %s, method: %s, invokeCount: %d", path, method, invokeCount),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
