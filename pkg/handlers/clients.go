package handlers

import (
	"net/http"
	"strconv"

	"raffiner.com.br/pkg/services"
	"raffiner.com.br/pkg/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

// GetClient -
func GetClient(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	queryid := req.QueryStringParameters["id"]
	if len(queryid) > 0 {
		id, err := strconv.Atoi(queryid)
		if err != nil {
			return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
		}
		result, err := services.FetchClient(id)
		if err != nil {
			return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}
	result, err := services.FetchClients()
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}

// CreateClient -
func CreateClient(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	result, err := services.CreateClient(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusCreated, result)
}

// UpdateClient -
func UpdateClient(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	result, err := services.UpdateClient(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}

// DeleteClient -
func DeleteClient(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	err := services.DeleteClient(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, nil)
}
