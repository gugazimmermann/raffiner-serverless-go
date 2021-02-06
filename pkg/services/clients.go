package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"raffiner.com.br/pkg/constants"
	"raffiner.com.br/pkg/types"
	"raffiner.com.br/pkg/validators"

	"github.com/aws/aws-lambda-go/events"
	//
	_ "github.com/go-sql-driver/mysql"
)

var connectionString = fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true",
	constants.DbUser,
	constants.DbPassword,
	constants.DbHost,
	constants.DbPort,
	constants.DbDatabase,
)

// FetchClient - Get a Client by ID
func FetchClient(id int) (*types.Client, error) {
	if id == 0 {
		return nil, errors.New(constants.ErrorInvalidID)
	}

	db, err := sql.Open("mysql", connectionString)
	validators.CheckError(err)
	defer db.Close()
	err = db.Ping()
	validators.CheckError(err)

	client := new(types.Client)
	row := db.QueryRow(`
		SELECT 
			c.id,
			c.email,
			c.name,
			c.phone,
			a.street,
			a.number,
			a.complement,
			a.neighborhood,
			a.city,
			a.state,
			a.zipcode
		FROM client c
		JOIN address a ON c.address_id = a.id 
		WHERE c.id= ?`, id)
	err = row.Scan(
		&client.ID,
		&client.Email,
		&client.Name,
		&client.Phone,
		&client.Address.Street,
		&client.Address.Number,
		&client.Address.Complement,
		&client.Address.Neighborhood,
		&client.Address.City,
		&client.Address.State,
		&client.Address.ZipCode,
	)
	if err != nil && err != sql.ErrNoRows {
		validators.CheckError(err)
	}
	return client, nil
}

// FetchClients - Get all Clients
func FetchClients() ([]*types.Client, error) {
	db, err := sql.Open("mysql", connectionString)
	validators.CheckError(err)
	defer db.Close()
	err = db.Ping()
	validators.CheckError(err)

	results, err := db.Query(`
		SELECT 
			c.id,
			c.email,
			c.name,
			c.phone,
			a.street,
			a.number,
			a.complement,
			a.neighborhood,
			a.city,
			a.state,
			a.zipcode
		FROM client c
		JOIN address a ON c.address_id = a.id`)
	if err != nil {
		validators.CheckError(err)
	}
	clients := []*types.Client{}
	for results.Next() {
		var client = new(types.Client)
		err = results.Scan(
			&client.ID,
			&client.Email,
			&client.Name,
			&client.Phone,
			&client.Address.Street,
			&client.Address.Number,
			&client.Address.Complement,
			&client.Address.Neighborhood,
			&client.Address.City,
			&client.Address.State,
			&client.Address.ZipCode,
		)
		if err != nil {
			validators.CheckError(err)
		}
		clients = append(clients, client)
	}
	return clients, nil
}

// CreateClient - Create a new Client with Address
func CreateClient(req events.APIGatewayProxyRequest) (*types.Client, error) {
	c := new(types.Client)

	err := json.Unmarshal([]byte(req.Body), &c)
	if err != nil {
		return nil, errors.New(constants.ErrorInvalidData)
	}
	if !validators.IsEmailValid(c.Email) {
		return nil, errors.New(constants.ErrorInvalidEmail)
	}

	db, err := sql.Open("mysql", connectionString)
	validators.CheckError(err)
	defer db.Close()
	err = db.Ping()
	validators.CheckError(err)

	sqlStatement, err := db.Prepare(`
		INSERT INTO address (
			street,
			number,
			complement,
			neighborhood,
			city,
			state,
			zipcode
		) VALUES (?, ?, ?, ?, ?, ?, ?);`)
	res, err := sqlStatement.Exec(
		c.Address.Street,
		c.Address.Number,
		c.Address.Complement,
		c.Address.Neighborhood,
		c.Address.City,
		c.Address.State,
		c.Address.ZipCode,
	)
	validators.CheckError(err)
	addressID, err := res.LastInsertId()
	validators.CheckError(err)

	sqlStatement, err = db.Prepare(`
		INSERT INTO client (
			email,
			name,
			phone,
			address_id
		) VALUES (?, ?, ?, ?);`)
	res, err = sqlStatement.Exec(
		c.Email,
		c.Name,
		c.Phone,
		int(addressID),
	)
	validators.CheckError(err)
	clientID, err := res.LastInsertId()
	validators.CheckError(err)

	client, err := FetchClient(int(clientID))
	return client, nil
}

func UpdateClient(req events.APIGatewayProxyRequest) (*types.Client, error) {
	c := new(types.Client)
	err := json.Unmarshal([]byte(req.Body), c)
	if err != nil {
		return nil, errors.New(constants.ErrorInvalidData)
	}
	if !validators.IsEmailValid(c.Email) {
		return nil, errors.New(constants.ErrorInvalidEmail)
	}
	return c, nil
}

func DeleteClient(req events.APIGatewayProxyRequest) error {
	id := req.QueryStringParameters["id"]
	if id == "" || !validators.IsIdInt(id) {
		return errors.New(constants.ErrorCouldNotDelete)
	}
	return nil
}
