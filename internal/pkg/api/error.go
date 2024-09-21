package api

import (
	"net/http"
)

var ErrInternalServerError = Error{
	httpStatus: http.StatusInternalServerError,
	msg:        "internal server error",
	code:       "500",
}

type Error struct {
	httpStatus int
	code       string
	err        error
	msg        string
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) Unwrap() error {
	return e.err
}

func InternalServerError(err error) error {
	e := ErrInternalServerError
	e.err = err
	return e
}

func Unauthorized() error {
	return Error{
		httpStatus: http.StatusUnauthorized,
		msg:        "unauthorized",
		code:       "401",
	}
}

func NotFound() error {
	return Error{
		httpStatus: http.StatusNotFound,
		msg:        "not found",
		code:       "404",
	}
}
