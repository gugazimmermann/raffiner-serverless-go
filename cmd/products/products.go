package main

import (
	"raffiner.com.br/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetProduct(req)
	case "POST":
		return handlers.CreateProduct(req)
	case "PUT":
		return handlers.UpdateProduct(req)
	case "DELETE":
		return handlers.DeleteProduct(req)
	default:
		return handlers.UnhandledMethod()
	}
}

func main() {
	lambda.Start(handler)
}
