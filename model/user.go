package model

import (
	"bytes"
	"encoding/json"
	"io"
)

// User type describes data of service's user
type User struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Email       string                 `json:"email"`
	Admin       bool                   `json:"admin"`
	Active      bool                   `json:"active"`
	Preferences map[string]interface{} `json:"preferences,omitempty"`
}

// ToReader returns user's data as io.Reader
func (u *User) ToReader() (io.Reader, error) {
	jsonReq, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonReq), nil
}
