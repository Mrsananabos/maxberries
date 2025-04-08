package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
	"time"
)

type Product struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:text" json:"name" valid:"required"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float32   `gorm:"type:decimal(10,2)" json:"price" valid:"required"`
	CategoryId  int64     `json:"category_id" valid:"required"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (p Product) Validate() error {
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
