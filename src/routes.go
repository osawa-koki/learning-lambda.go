package main

import (
	"github.com/aws/aws-lambda-go/events"
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
