package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func handler_hello_get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("hello, GET!"),
		StatusCode: 200,
	}, nil
}

func handler_hello_post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("hello, POST!"),
		StatusCode: 200,
	}, nil
}

func handler_hello_put(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("hello, PUT!"),
		StatusCode: 200,
	}, nil
}

func handler_hello_delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("hello, DELETE!"),
		StatusCode: 200,
	}, nil
}

func handler_env_get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("DOTENV_PARAM_1: %s\nDOTENV_PARAM_2: %s\nDOTENV_PARAM_3: %s\n", DOTENV_PARAM_1, DOTENV_PARAM_2, DOTENV_PARAM_3),
		StatusCode: 200,
	}, nil
}
