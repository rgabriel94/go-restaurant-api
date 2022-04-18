package service

import (
	"github.com/go-playground/validator/v10"
	"go-restaurant-api/api/model/exception"
)

var validate = validator.New()

func Validate(value any) {
	if err := validate.Struct(value); err != nil {
		if err != nil {
			exception.PanicBadRequest(err.Error())
		}
	}
}
