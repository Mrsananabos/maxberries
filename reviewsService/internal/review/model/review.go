package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProductID string             `bson:"product_id"`
	UserID    string             `bson:"user_id"`
	Rating    int                `bson:"rating"`
	Text      string             `bson:"text"`
}

//func (r Review) Validate() error {
//	valid, err := govalidator.ValidateStruct(c)
//	if !valid {
//		var validationErrors govalidator.Errors
//		if errors.As(err, &validationErrors) {
//			errorsStr := strings.Builder{}
//			for _, validationError := range validationErrors {
//				errorsStr.WriteString(validationError.Error())
//				errorsStr.WriteString("\n")
//			}
//			return fmt.Errorf(errorsStr.String())
//		}
//	}
//	return nil
//}
