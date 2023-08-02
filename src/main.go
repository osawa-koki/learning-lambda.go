package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"log"
	"os"
)

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

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.Path
	method := Method(request.HTTPMethod)
	if _, ok := Routes[path]; !ok {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Not Found"),
			StatusCode: 404,
		}, nil
	}
	if _, ok := Routes[path][method]; !ok {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Method Not Allowed"),
			StatusCode: 405,
		}, nil
	}
	return Routes[path][method](request)
}

func main() {
	lambda.Start(handler)
}
