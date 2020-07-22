package staxxapi

import (
	"github.com/adiletabylov/staxxapi/client"
	"github.com/adiletabylov/staxxapi/helpers"
	"github.com/adiletabylov/staxxapi/model"
)

// ListUsers returns list of users
func ListUsers() (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "user")
	return client.Get(url)
}

// GetUserByID returns user details by given id
func GetUserByID(userID string) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "user", userID)
	return client.Get(url)
}

// CreateUser makes request to create user with given data
func CreateUser(user *model.User) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "user")
	bytes, err := user.ToReader()
	if err != nil {
		return nil, err
	}
	return client.Post(url, bytes)
}

// UpdateUser makes request to update user with given ID
func UpdateUser(userID string, user *model.User) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "user", userID)
	bytes, err := user.ToReader()
	if err != nil {
		return nil, err
	}
	return client.Post(url, bytes)
}
