package services

import (
	"encoding/json"
	"errors"

	"raffiner.com.br/pkg/validators"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ErrorInvalidUserID      = "invalid user ID"
	ErrorInvalidUserData    = "invalid user data"
	ErrorInvalidEmail       = "invalid email"
	ErrorCouldNotDeleteItem = "could not delete item"
)

type Address struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipcode"`
}

type Client struct {
	ID      int     `json:"id,omitempty"`
	Email   string  `json:"email"`
	Name    string  `json:"Name"`
	Address Address `json:"address"`
	Phone   string  `json:"phone"`
}

func FetchClient(id int) (*Client, error) {
	if id == 0 {
		return nil, errors.New(ErrorInvalidUserID)
	}

	a := new(Address)
	a.Street = "Rua Um"
	a.Number = "123"
	a.Complement = "Bloco B"
	a.Neighborhood = "Bairro Dois"
	a.City = "Itajaí"
	a.State = "Santa Catarina"
	a.ZipCode = "88300-000"

	c := new(Client)
	c.ID = id
	c.Email = "client@email.com"
	c.Name = "Client One"
	c.Phone = "47 99999-9999"
	c.Address = *a

	return c, nil
}

func FetchClients() ([]*Client, error) {
	a := new(Address)
	a.Street = "Rua Um"
	a.Number = "123"
	a.Complement = "Bloco B"
	a.Neighborhood = "Bairro Dois"
	a.City = "Itajaí"
	a.State = "Santa Catarina"
	a.ZipCode = "88300-000"

	c := new(Client)
	c.ID = 1
	c.Email = "client@email.com"
	c.Name = "Client One"
	c.Phone = "47 99999-9999"
	c.Address = *a

	d := new(Client)
	d.ID = 2
	d.Email = "client2@email.com"
	d.Name = "Client Two"
	d.Phone = "47 98888-8888"
	d.Address = *a

	clients := []*Client{c, d}

	return clients, nil
}

func CreateClient(req events.APIGatewayProxyRequest) (
	*Client,
	error,
) {
	c := new(Client)
	err := json.Unmarshal([]byte(req.Body), c)
	if err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(c.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}
	c.ID = 1
	return c, nil
}

func UpdateClient(req events.APIGatewayProxyRequest) (
	*Client,
	error,
) {
	c := new(Client)
	err := json.Unmarshal([]byte(req.Body), c)
	if err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(c.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}
	return c, nil
}

func DeleteClient(req events.APIGatewayProxyRequest) error {
	id := req.QueryStringParameters["id"]
	if id == "" || !validators.IsIdInt(id) {
		return errors.New(ErrorCouldNotDeleteItem)
	}
	return nil
}
