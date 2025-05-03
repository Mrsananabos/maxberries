package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type Register struct {
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
	RoleID   uint64 `json:"role_id" valid:"required"`
}

type Login struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
	Ip       string
}

func (r Register) Validate() error {
	valid, err := govalidator.ValidateStruct(r)
	if !valid {
		var validationErrors govalidator.Errors
		if errors.As(err, &validationErrors) {
			errorsStr := strings.Builder{}
			for _, validationError := range validationErrors {
				errorsStr.WriteString(validationError.Error())
				errorsStr.WriteString("\n")
			}
			return fmt.Errorf(errorsStr.String())
		}
	}
	return nil
}

func (l Login) Validate() error {
	valid, err := govalidator.ValidateStruct(l)
	if !valid {
		var validationErrors govalidator.Errors
		if errors.As(err, &validationErrors) {
			errorsStr := strings.Builder{}
			for _, validationError := range validationErrors {
				errorsStr.WriteString(validationError.Error())
				errorsStr.WriteString("\n")
			}
			return fmt.Errorf(errorsStr.String())
		}
	}
	return nil
}
