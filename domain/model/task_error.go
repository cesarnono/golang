package model

import (
	"fmt"
	"net/http"
)


type ValidationError struct {
	error string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", v.error)
}

func (v ValidationError) StatusCode() int {
	return http.StatusBadRequest
}

type NotFoundError struct {
	error string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("validation error: %s", n.error)
}

func (n NotFoundError) StatusCode() int {
	return http.StatusNotFound
}