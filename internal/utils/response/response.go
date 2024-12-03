package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type ErrorResponse struct {
	Status string   `json:"status"`
	Errors []string `json:"errors"`
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) ErrorResponse {
	return ErrorResponse{
		Status: StatusError,
		Errors: []string{
			err.Error(),
		},
	}
}

func ValidationError(errs validator.ValidationErrors) ErrorResponse {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}
	return ErrorResponse{
		Status: StatusError,
		Errors: errMsgs,
	}
}
