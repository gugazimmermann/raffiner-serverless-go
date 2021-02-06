package main

import (
	"raffiner.com.br/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req)
	case "POST":
		return handlers.CreateUser(req)
	case "PUT":
		return handlers.UpdateUser(req)
	case "DELETE":
		return handlers.DeleteUser(req)
	default:
		return handlers.UnhandledMethod()
	}
}

func main() {
	lambda.Start(handler)
}
