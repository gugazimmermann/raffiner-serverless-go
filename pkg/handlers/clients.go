package handlers

import (
	"net/http"
	"strconv"

	"raffiner.com.br/pkg/services"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

var ErrorMethodNotAllowed = "method Not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetClient(req events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse,
	error,
) {
	queryid := req.QueryStringParameters["id"]
	if len(queryid) > 0 {
		id, err := strconv.Atoi(queryid)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		result, err := services.FetchClient(id)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}

		return apiResponse(http.StatusOK, result)
	}

	result, err := services.FetchClients()
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

func CreateClient(req events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse,
	error,
) {
	result, err := services.CreateClient(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)
}

func UpdateClient(req events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse,
	error,
) {
	result, err := services.UpdateClient(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

func DeleteClient(req events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse,
	error,
) {
	err := services.DeleteClient(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
