package services

import (
	"encoding/json"
	"errors"

	"raffiner.com.br/pkg/types"
	"raffiner.com.br/pkg/validators"

	"github.com/aws/aws-lambda-go/events"
)

func FetchSupplier(id int) (*types.Supplier, error) {
	if id == 0 {
		return nil, errors.New(types.ErrorInvalidID)
	}

	a := new(types.Address)
	a.Street = "Rua Um"
	a.Number = "123"
	a.Complement = "Bloco B"
	a.Neighborhood = "Bairro Dois"
	a.City = "Itajaí"
	a.State = "Santa Catarina"
	a.ZipCode = "88300-000"

	c := new(types.Supplier)
	c.ID = id
	c.Email = "supplier@email.com"
	c.Name = "Supplier One"
	c.Phone = "47 99999-9999"
	c.Address = *a

	return c, nil
}

func FetchSuppliers() ([]*types.Supplier, error) {
	a := new(types.Address)
	a.Street = "Rua Um"
	a.Number = "123"
	a.Complement = "Bloco B"
	a.Neighborhood = "Bairro Dois"
	a.City = "Itajaí"
	a.State = "Santa Catarina"
	a.ZipCode = "88300-000"

	c := new(types.Supplier)
	c.ID = 1
	c.Email = "supplier@email.com"
	c.Name = "Supplier One"
	c.Phone = "47 99999-9999"
	c.Address = *a

	d := new(types.Supplier)
	d.ID = 2
	d.Email = "supplier2@email.com"
	d.Name = "Supplier Two"
	d.Phone = "47 98888-8888"
	d.Address = *a

	clients := []*types.Supplier{c, d}

	return clients, nil
}

func CreateSupplier(req events.APIGatewayProxyRequest) (*types.Supplier, error) {
	c := new(types.Supplier)
	err := json.Unmarshal([]byte(req.Body), c)
	if err != nil {
		return nil, errors.New(types.ErrorInvalidData)
	}
	if !validators.IsEmailValid(c.Email) {
		return nil, errors.New(types.ErrorInvalidEmail)
	}
	c.ID = 1
	return c, nil
}

func UpdateSupplier(req events.APIGatewayProxyRequest) (*types.Supplier, error) {
	c := new(types.Supplier)
	err := json.Unmarshal([]byte(req.Body), c)
	if err != nil {
		return nil, errors.New(types.ErrorInvalidData)
	}
	if !validators.IsEmailValid(c.Email) {
		return nil, errors.New(types.ErrorInvalidEmail)
	}
	return c, nil
}

func DeleteSupplier(req events.APIGatewayProxyRequest) error {
	id := req.QueryStringParameters["id"]
	if id == "" || !validators.IsIdInt(id) {
		return errors.New(types.ErrorCouldNotDelete)
	}
	return nil
}
