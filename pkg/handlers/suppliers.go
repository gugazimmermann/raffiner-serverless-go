package handlers

import (
	"net/http"
	"strconv"

	"raffiner.com.br/pkg/services"
	"raffiner.com.br/pkg/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

func GetSupplier(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	queryid := req.QueryStringParameters["id"]
	if len(queryid) > 0 {
		id, err := strconv.Atoi(queryid)
		if err != nil {
			return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
		}
		result, err := services.FetchSupplier(id)
		if err != nil {
			return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}
	result, err := services.FetchSuppliers()
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}

func CreateSupplier(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	result, err := services.CreateSupplier(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusCreated, result)
}

func UpdateSupplier(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	result, err := services.UpdateSupplier(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}

func DeleteSupplier(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	err := services.DeleteSupplier(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, types.ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, nil)
}
