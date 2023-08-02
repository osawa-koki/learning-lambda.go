package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var routes = map[string]map[Method]func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	"/hello": {
		GET:    handler_hello_get,
		POST:   handler_hello_post,
		PUT:    handler_hello_put,
		DELETE: handler_hello_delete,
	},
	"/env": {
		GET: handler_env_get,
	},
}

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

var DOTENV_PARAM_1 string
var DOTENV_PARAM_2 string
var DOTENV_PARAM_3 string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
		os.Exit(1)
	}
	DOTENV_PARAM_1 = os.Getenv("DOTENV_PARAM_1")
	DOTENV_PARAM_2 = os.Getenv("DOTENV_PARAM_2")
	DOTENV_PARAM_3 = os.Getenv("DOTENV_PARAM_3")
}

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

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.Path
	method := Method(request.HTTPMethod)
	if _, ok := routes[path]; !ok {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Not Found"),
			StatusCode: 404,
		}, nil
	}
	if _, ok := routes[path][method]; !ok {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Method Not Allowed"),
			StatusCode: 405,
		}, nil
	}
	return routes[path][method](request)
}

func main() {
	lambda.Start(handler)
}
