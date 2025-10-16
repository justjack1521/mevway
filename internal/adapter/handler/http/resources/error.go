package resources

import (
	"errors"
	"net/http"
)

var (
	errInvalidRequest = errors.New("invalid request")
)

type ErrorResponse struct {
	Code    int    `json:"Code"`
	Slug    string `json:"Slug"`
	Message string `json:"Message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

var (
	ErrInvalidRequestResponse = ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: errInvalidRequest.Error(),
	}
)
