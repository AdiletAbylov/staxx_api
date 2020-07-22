package staxxapi

import (
	"fmt"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/adiletabylov/staxxapi/model"
	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {
	Init("http://localhost", "4000")
	resp, _ := ListUsers()

	assert.Equal(t, true, resp.IsResponseStatusOK(), "Reponse status should OK")
}

func TestCreateUser(t *testing.T) {
	Init("http://localhost", "4000")
	user := model.User{
		Email:  randomdata.Email(),
		Name:   randomdata.FullName(1),
		Active: true,
		Admin:  false,
	}
	resp, _ := CreateUser(&user)
	assert.Equal(t, true, resp.IsResponseStatusOK(), "Reponse status should be OK")

	userID := fmt.Sprintf("%g", resp.Data.(map[string]interface{})["id"])

	// test GetUserByID also
	getresp, _ := GetUserByID(userID)
	assert.Equal(t, true, getresp.IsResponseStatusOK(), "Reponse status should OK")

	assert.Equal(t, user.Name, getresp.Data.(map[string]interface{})["name"], "Created and returned users should be the same")
}

func TestUpdateUser(t *testing.T) {
	Init("http://localhost", "4000")
	user := model.User{
		Email:  randomdata.Email(),
		Name:   randomdata.FullName(1),
		Active: true,
		Admin:  false,
	}
	resp, _ := CreateUser(&user)
	assert.Equal(t, true, resp.IsResponseStatusOK(), "Reponse status should be OK")
	userID := fmt.Sprintf("%g", resp.Data.(map[string]interface{})["id"])

	user.Name = randomdata.FullName(1)
	user.Active = false

	updresp, _ := UpdateUser(userID, &user)

	assert.Equal(t, true, updresp.IsResponseStatusOK(), "Reponse status should be OK")
	assert.Equal(t, user.Name, updresp.Data.(map[string]interface{})["name"], "Created and updated names should be the same")
	assert.Equal(t, user.Active, updresp.Data.(map[string]interface{})["active"], "Created and updated 'active' field should be the same")
}
