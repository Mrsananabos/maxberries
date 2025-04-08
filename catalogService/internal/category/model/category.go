package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
	"time"
)

type Category struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:text" json:"name" valid:"required"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (c Category) Validate() error {
	valid, err := govalidator.ValidateStruct(c)
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
