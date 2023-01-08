package validator

import (
	"github.com/go-playground/validator/v10"
)

var GenderValidate validator.Func = func(fl validator.FieldLevel) bool {
	f, ok := fl.Field().Interface().(string)
	if ok {
		if f == "L" || f == "P" {
			return true
		}
	}
	return false
}
