package main

import (
	"raffiner.com.br/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetSupplier(req)
	case "POST":
		return handlers.CreateSupplier(req)
	case "PUT":
		return handlers.UpdateSupplier(req)
	case "DELETE":
		return handlers.DeleteSupplier(req)
	default:
		return handlers.UnhandledMethod()
	}
}

func main() {
	lambda.Start(handler)
}
