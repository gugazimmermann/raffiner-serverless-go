package services

import (
	"encoding/json"
	"errors"

	"raffiner.com.br/pkg/types"
	"raffiner.com.br/pkg/validators"

	"github.com/aws/aws-lambda-go/events"
)

func FetchClient(id int) (*types.Client, error) {
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

	c := new(types.Client)
	c.ID = id
	c.Email = "client@email.com"
	c.Name = "Client One"
	c.Phone = "47 99999-9999"
	c.Address = *a

	return c, nil
}

func FetchClients() ([]*types.Client, error) {
	a := new(types.Address)
	a.Street = "Rua Um"
	a.Number = "123"
	a.Complement = "Bloco B"
	a.Neighborhood = "Bairro Dois"
	a.City = "Itajaí"
	a.State = "Santa Catarina"
	a.ZipCode = "88300-000"

	c := new(types.Client)
	c.ID = 1
	c.Email = "client@email.com"
	c.Name = "Client One"
	c.Phone = "47 99999-9999"
	c.Address = *a

	d := new(types.Client)
	d.ID = 2
	d.Email = "client2@email.com"
	d.Name = "Client Two"
	d.Phone = "47 98888-8888"
	d.Address = *a

	clients := []*types.Client{c, d}

	return clients, nil
}

func CreateClient(req events.APIGatewayProxyRequest) (*types.Client, error) {
	c := new(types.Client)
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

func UpdateClient(req events.APIGatewayProxyRequest) (*types.Client, error) {
	c := new(types.Client)
	err := json.Unmarshal([]byte(req.Body), c)
	if err != nil {
		return nil, errors.New(types.ErrorInvalidData)
	}
	if !validators.IsEmailValid(c.Email) {
		return nil, errors.New(types.ErrorInvalidEmail)
	}
	return c, nil
}

func DeleteClient(req events.APIGatewayProxyRequest) error {
	id := req.QueryStringParameters["id"]
	if id == "" || !validators.IsIdInt(id) {
		return errors.New(types.ErrorCouldNotDelete)
	}
	return nil
}
