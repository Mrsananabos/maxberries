package model

import (
	"authService/internal/permission/model"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type Role struct {
	ID          uint64             `gorm:"primary_key" json:"id"`
	Name        string             `gorm:"size:50;not null;unique" json:"name" valid:"required"`
	Description string             `gorm:"size:255;not null" json:"description"`
	Permissions []model.Permission `gorm:"many2many:roles_permissions"  json:"permissions" valid:"required"`
}

func (r Role) Validate() error {
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
