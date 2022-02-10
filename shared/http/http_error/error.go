package http_error

import (
	"github.com/labstack/echo/v4"
)

type JsonError struct {
	Errors map[string]interface{} `json:"error"`
}

func NewError(err error) JsonError {
	e := JsonError{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

func AccessForbidden() JsonError {
	e := JsonError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return e
}

func NotFound() JsonError {
	e := JsonError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "resource not found"
	return e
}
