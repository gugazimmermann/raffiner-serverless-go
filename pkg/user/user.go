package user

import (
	"encoding/json"
	"errors"

	"raffiner.com.br/pkg/validators"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ErrorInvalidUserData    = "invalid user data"
	ErrorInvalidEmail       = "invalid email"
	ErrorCouldNotDeleteItem = "could not delete item"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FetchUser(email string) (User, error) {
	var item User
	item.Email = email
	item.FirstName = "User"
	item.LastName = "Test"

	return item, nil
}

func FetchUsers() ([]User, error) {
	var item User
	item.Email = "user@test.com"
	item.FirstName = "User"
	item.LastName = "Test"

	var item2 User
	item2.Email = "user2@test.com"
	item2.FirstName = "User"
	item2.LastName = "Test 2"

	users := []User{item, item2}

	return users, nil
}

func CreateUser(req events.APIGatewayProxyRequest) (
	*User,
	error,
) {
	item := new(User)
	err := json.Unmarshal([]byte(req.Body), item)
	if err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(item.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}
	return item, nil
}

func UpdateUser(req events.APIGatewayProxyRequest) (
	*User,
	error,
) {
	item := new(User)
	err := json.Unmarshal([]byte(req.Body), item)
	if err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(item.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}
	return item, nil
}

func DeleteUser(req events.APIGatewayProxyRequest) error {
	email := req.QueryStringParameters["email"]
	if len(email) == 0 {
		return errors.New(ErrorCouldNotDeleteItem)
	}
	if !validators.IsEmailValid(email) {
		return errors.New(ErrorInvalidEmail)
	}
	return nil
}
