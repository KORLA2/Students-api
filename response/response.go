package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) Response {

	return Response{
		"Error",
		err.Error(),
	}

}

func ValidationError(errs validator.ValidationErrors) Response {
	var errmsg string
	for _, err := range errs {

		switch err.ActualTag() {
		case "required":
			errmsg = fmt.Sprintf("%s Filed is reqired", err.Field())

		default:
			errmsg = fmt.Sprintf("%s file dis invalid", err.Field())

		}
	}

	return Response{
		StatusError,
		errmsg,
	}

}
