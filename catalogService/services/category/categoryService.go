package category

import (
	"catalogService/models"
	"catalogService/services"
	"catalogService/storage/postgres/config"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"strings"
)

type ServiceInterface interface {
	GetAll() ([]models.Category, error)
	GetByIdGetById(id int64) (models.Category, error)
	Save(category models.Category)
	Delete(id int64)
	Update(id int64, category models.Category)
}

type Service struct {
}

func (service Service) GetAll() (categories []models.Category, err error) {
	rsl := config.DB.Find(&categories)

	if rsl.Error != nil {
		err = rsl.Error
		return
	}

	return
}

func (service Service) GetById(id int64) (category models.Category, err error) {
	rsl := config.DB.Take(&category, id)

	if errors.Is(rsl.Error, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("Not found category with id = %d", id)
		return
	}

	if rsl.Error != nil {
		err = rsl.Error
		return
	}

	return
}

func (service Service) Save(category models.Category) error {
	if err := validateCategory(category); err != nil {
		return err
	}

	rsl := config.DB.Create(&category)

	if rsl.Error != nil {
		return rsl.Error
	}

	return nil
}

func (service Service) Delete(id int64) error {
	rsl := config.DB.Delete(&models.Category{}, id)

	if rsl.Error != nil {
		return rsl.Error
	}

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("Not found category with id = %d", id)
	}

	return nil
}

func (service Service) Update(category models.Category) (err error) {
	if err = validateCategory(category); err != nil {
		return
	}

	rsl := config.DB.Model(&models.Category{}).Where("id = ?", category.Id).Updates(category)

	if rsl.Error != nil {
		err = rsl.Error
		return
	}

	if rsl.RowsAffected != 1 {
		err = fmt.Errorf("Not found category with id = %d", category.Id)
		return
	}

	return
}

func validateCategory(category models.Category) error {
	err := services.Validate.Struct(category)
	if err != nil {
		strBuilder := strings.Builder{}

		for _, err := range err.(validator.ValidationErrors) {
			strBuilder.WriteString(fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		return fmt.Errorf(strBuilder.String())
	}

	return nil
}
