package common

import (
	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validate *validator.Validate
}

func NewValidator() *StructValidator {
	return &StructValidator{validate: validator.New()}
}

func (v *StructValidator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}
