package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type Permission struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	Code        string `gorm:"size:50;not null;unique" json:"code" valid:"required"`
	Description string `gorm:"size:255;not null" json:"description"`
}

func (p Permission) Validate() error {
	valid, err := govalidator.ValidateStruct(p)
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
