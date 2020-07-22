package model

import (
	"encoding/json"
	"io"
)

const responseStatusOK int64 = 0

// Response type describes response struct from service
type Response struct {
	Data    interface{}
	Errors  []ResponseError "json:omitempty"
	Status  int64
	Message string
}

// ResponseError type describes error returned from service
type ResponseError struct {
	Detail string
}

// IsResponseStatusOK returns if response status is OK or error
func (r *Response) IsResponseStatusOK() bool {
	return r.Status == responseStatusOK
}

// NewResponseFromBody decodes and returns model.Response struct
func NewResponseFromBody(body io.ReadCloser) (*Response, error) {
	var response Response
	err := json.NewDecoder(body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
