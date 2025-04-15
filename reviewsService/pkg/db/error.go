package db

//
//import (
//	"errors"
//	"fmt"
//	"gorm.io/gorm"
//)
//
//type ErrObjectNotFound struct {
//	msg string
//}
//
//func (e ErrObjectNotFound) Error() string {
//	return e.msg
//}
//
//func (e ErrObjectNotFound) Unwrap() error {
//	return fmt.Errorf(e.Error())
//}
//
//func HandleError(err error) error {
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return ErrObjectNotFound{
//			msg: "object not found",
//		}
//	}
//	return err
//}
