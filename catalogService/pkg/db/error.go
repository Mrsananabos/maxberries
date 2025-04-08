package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ErrObjectNotFound struct{}

func (ErrObjectNotFound) Error() string {
	return "object not found"
}

func (ErrObjectNotFound) Unwrap() error {
	return fmt.Errorf("object not found")
}

func HandleError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrObjectNotFound{}
	}
	return err
}
