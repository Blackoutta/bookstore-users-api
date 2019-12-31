package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Error string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "Bad request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "Not found",
	}
}