package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"unicode/utf8"
)

const RATING_MAX = 5
const TEXT_MIN_LEN = 15

type Review struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProductID int64              `bson:"product_id" json:"product_id" valid:"required"`
	UserID    string             `bson:"user_id" json:"user_id" valid:"required"`
	Rating    *int               `bson:"rating" json:"rating"`
	Text      string             `bson:"text" json:"text" valid:"required"`
}

type ContentReview struct {
	Rating *int   `bson:"rating" json:"rating"`
	Text   string `bson:"text" json:"text"`
}

func (r Review) Validate() error {
	errorsStr := strings.Builder{}
	valid, err := govalidator.ValidateStruct(r)
	if !valid {
		var validationErrors govalidator.Errors
		if errors.As(err, &validationErrors) {
			for _, validationError := range validationErrors {
				errorsStr.WriteString(validationError.Error())
				errorsStr.WriteString("\n")
			}
		}
	}

	err = ContentReview{r.Rating, r.Text}.Validate()
	if err != nil {
		errorsStr.WriteString(err.Error())
	}

	if errorsStr.Len() != 0 {
		return fmt.Errorf(errorsStr.String())
	}

	return nil
}

func (c ContentReview) Validate() error {
	errorsStr := strings.Builder{}
	if c.Rating != nil {
		if *c.Rating > RATING_MAX {
			errorsStr.WriteString(fmt.Sprintf("Rating can`t be greater than %d\n", RATING_MAX))
		}
		if *c.Rating < 0 {
			errorsStr.WriteString("Rating can`t be negative\n")
		}
	}

	if utf8.RuneCountInString(c.Text) < TEXT_MIN_LEN {
		errorsStr.WriteString(fmt.Sprintf("Text can`t be less %d symbols\n", TEXT_MIN_LEN))
	}

	if errorsStr.Len() != 0 {
		return fmt.Errorf(errorsStr.String())
	}

	return nil
}
