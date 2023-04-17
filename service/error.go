package service

import "net/http"

type ServiceError struct {
	Message string
	Code    int // HTTP status code
}

func (se ServiceError) Error() string {
	return se.Message
}

func NewErrBadRequest(msg string) ServiceError {
	return ServiceError{Message: msg, Code: http.StatusBadRequest}
}

func NewErrNotFound(msg string) ServiceError {
	return ServiceError{Message: msg, Code: http.StatusNotFound}
}
