package mperrors

import (
	"encoding/json"
	"errors"
	"fmt"
)

type APIError struct {
	StatusCode int
	RawBody    []byte
	Response   APIErrorResponse
}

type APIErrorResponse struct {
	Message string       `json:"message,omitempty"`
	Error   error        `json:"error,omitempty"`
	Status  int          `json:"status,omitempty"`
	Cause   []ErrorCause `json:"cause,omitempty"`
}

type ErrorCause struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Field       string `json:"field,omitempty"`
}

func (e *APIError) Error() string {
	if e.Response.Error != nil {
		return fmt.Sprintf("mercadopago: error=%s status=%d", e.Response.Error, e.StatusCode)
	}
	return fmt.Sprintf("mercadopago: status=%d", e.StatusCode)
}

func Parse(status int, body []byte) *APIError {
	var resp APIErrorResponse
	_ = json.Unmarshal(body, &resp)

	return &APIError{
		StatusCode: status,
		RawBody:    body,
		Response:   resp,
	}
}

var (
	ErrInvalidConfig  = errors.New("invalid mercadopago config")
	ErrNotImplemented = errors.New("mercadopago: not implemented yet")
)
