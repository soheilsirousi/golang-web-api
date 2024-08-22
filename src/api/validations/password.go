package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/soheilsirousi/golang-web-api/src/commons"
)

func PasswordValidation(v validator.FieldLevel) bool {
	password, ok := v.Field().Interface().(string)

	if !ok {
		v.Param()
		return false
	}

	return commons.CheckPassword(password)
}
