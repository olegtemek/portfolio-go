package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return Response{
		Status: 400,
		Error:  strings.Join(errMsgs, ", "),
	}
}

func New(w *http.ResponseWriter, r *http.Request, err error, statusCode int) {
	(*w).WriteHeader(statusCode)
	render.JSON(*w, r, &Response{
		Status: statusCode,
		Error:  err.Error(),
	})
}
