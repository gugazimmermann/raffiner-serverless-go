package main

import (
	"raffiner.com.br/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetClient(req)
	case "POST":
		return handlers.CreateClient(req)
	case "PUT":
		return handlers.UpdateClient(req)
	case "DELETE":
		return handlers.DeleteClient(req)
	default:
		return handlers.UnhandledMethod()
	}
}

func main() {
	lambda.Start(handler)
}
