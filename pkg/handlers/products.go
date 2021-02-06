package handlers

import (
	"net/http"
	"strconv"

	"raffiner.com.br/pkg/services"
	"raffiner.com.br/pkg/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

func GetProduct(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	queryid := req.QueryStringParameters["id"]
	if len(queryid) > 0 {
		id, err := strconv.Atoi(queryid)
		if err != nil {
			return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
		}
		result, err := services.FetchProduct(id)
		if err != nil {
			return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}
	result, err := services.FetchProducts()
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}

func CreateProduct(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	result, err := services.CreateProduct(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusCreated, result)
}

func UpdateProduct(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	result, err := services.UpdateProduct(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}

func DeleteProduct(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	err := services.DeleteProduct(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, nil)
}
