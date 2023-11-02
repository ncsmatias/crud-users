package resterr

import (
	"net/http"
)

type RestErr struct {
	Message     string   `json:"message"`
	Err         string   `json:"error"`
	Description string   `json:"description"`
	Code        int      `json:"code"`
	Causes      []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, err string, code int, causes []Causes) *RestErr {

	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func BadRequestError(message string) *RestErr {

	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func BadRequestValidationError(message string, causes []Causes) *RestErr {

	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func InternalServerError(message string, description string) *RestErr {

	return &RestErr{
		Message:     message,
		Err:         "internal server error",
		Description: description,
		Code:        http.StatusInternalServerError,
	}
}

func NotFoundError(message string) *RestErr {

	return &RestErr{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func Forbidden(message string) *RestErr {

	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
