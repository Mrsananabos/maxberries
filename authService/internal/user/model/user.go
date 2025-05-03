package model

import (
	"authService/internal/role/model"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid" json:"id"`
	RoleID    uint64     `gorm:"not null" json:"-"`
	Role      model.Role `gorm:"foreignKey:RoleID" json:"role"`
	Username  string     `gorm:"size:255;not null;unique" json:"username"`
	Email     string     `gorm:"size:255;not null;unique" json:"email"`
	Password  string     `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

type UserUpdateForm struct {
	Username string `gorm:"size:255;not null;unique" json:"username" valid:"optional"`
	Email    string `gorm:"size:255;not null;unique" json:"email" valid:"email,optional"`
	Password string `gorm:"size:255;not null" json:"password" valid:"optional"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func (c *User) Validate() error {
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

func (uForm UserUpdateForm) Validate() error {
	valid, err := govalidator.ValidateStruct(uForm)
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
