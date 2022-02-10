package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"http/http_error"
)

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewValidatorError(err error) http_error.JsonError {
	e := http_error.JsonError{}
	e.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.Errors[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}
	return e
}
