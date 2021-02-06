package services

import (
	"encoding/json"
	"errors"

	"raffiner.com.br/pkg/types"
	"raffiner.com.br/pkg/validators"

	"github.com/aws/aws-lambda-go/events"
)

func FetchProduct(id int) (*types.Product, error) {
	if id == 0 {
		return nil, errors.New(types.ErrorInvalidID)
	}

	p := new(types.Product)
	p.ID = id
	p.Name = "Dish"
	p.Description = "Hand painted porcelain plate"
	p.Supplier = "Supplier A"
	p.Tags = []string{"Hand painted", "porcelain", "plate"}
	p.Category = "Plates"
	p.Subcategory = "Porcelain"
	p.CostPrice = 59.99
	p.RentalPrice = 19.99
	p.Quantity = 12

	return p, nil
}

func FetchProducts() ([]*types.Product, error) {
	p := new(types.Product)
	p.ID = 1
	p.Name = "Dish"
	p.Description = "Hand painted porcelain plate"
	p.Supplier = "Supplier A"
	p.Tags = []string{"Hand painted", "Porcelain", "Plate"}
	p.Category = "Plates"
	p.Subcategory = "Porcelain"
	p.CostPrice = 59.99
	p.RentalPrice = 19.99
	p.Quantity = 12

	p1 := new(types.Product)
	p1.ID = 2
	p1.Name = "Platter"
	p1.Description = "Silver Platter - Big"
	p1.Supplier = "Supplier B"
	p1.Tags = []string{"Silver", "Platter", "Big"}
	p1.Category = "Platters"
	p1.Subcategory = "Silver"
	p1.CostPrice = 210
	p1.RentalPrice = 70
	p1.Quantity = 3

	products := []*types.Product{p, p1}

	return products, nil
}

func CreateProduct(req events.APIGatewayProxyRequest) (*types.Product, error) {
	p := new(types.Product)
	err := json.Unmarshal([]byte(req.Body), p)
	if err != nil {
		return nil, errors.New(types.ErrorInvalidData)
	}
	p.ID = 1
	return p, nil
}

func UpdateProduct(req events.APIGatewayProxyRequest) (*types.Product, error) {
	p := new(types.Product)
	err := json.Unmarshal([]byte(req.Body), p)
	if err != nil {
		return nil, errors.New(types.ErrorInvalidData)
	}
	return p, nil
}

func DeleteProduct(req events.APIGatewayProxyRequest) error {
	id := req.QueryStringParameters["id"]
	if id == "" || !validators.IsIdInt(id) {
		return errors.New(types.ErrorCouldNotDelete)
	}
	return nil
}
