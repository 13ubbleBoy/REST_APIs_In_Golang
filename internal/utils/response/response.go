package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Error  string `json:"error"`
	Status string `json:"status"` // esse "Status" ka "S" small letter me likha aayega
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error { // "interface{}" eske place pe "any" bhi likh sakte h.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data) // converting struct data into json for that we have to encode the data.
}

// error ko json form me return kar rhe
func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response { // accepting slice of errors
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required field", err.Field()))

		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
