package product

import (
	"catalogService/models"
	"catalogService/services"
	"catalogService/services/category"
	"catalogService/storage/postgres/config"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"strings"
)

type ServiceInterface interface {
	GetAll() []models.Product
	Save(Product models.Product)
	Delete(id int64)
	UpdateFieldsById(id int64, fields map[string]interface{})
}

type Service struct {
}

var categoryService = category.Service{}

func (service Service) GetAll() []models.Product {
	var products []models.Product
	rsl := config.DB.Find(&products)

	if rsl.Error != nil {
		fmt.Println(rsl.Error)
	}

	return products
}

func (service Service) GetById(id int64) (product models.Product, err error) {
	rsl := config.DB.Take(&product, id)

	if errors.Is(rsl.Error, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("Not found product with id = %d", id)
		return
	}

	if rsl.Error != nil {
		err = rsl.Error
		return
	}

	return
}

func (service Service) Save(product models.Product) error {
	if err := validateProduct(product); err != nil {
		return err
	}

	_, err := categoryService.GetById(product.CategoryId)
	if err != nil {
		return err
	}

	rsl := config.DB.Create(&product)

	if rsl.Error != nil {
		return rsl.Error
	}

	return nil
}

func (service Service) Update(product models.Product) (err error) {
	if err = validateProduct(product); err != nil {
		return
	}

	rsl := config.DB.Model(&models.Product{}).Where("id = ?", product.Id).Updates(product)

	if rsl.Error != nil {
		err = rsl.Error
		return
	}

	if rsl.RowsAffected != 1 {
		err = fmt.Errorf("Not found product with id = %d", product.Id)
		return
	}

	return
}

func (service Service) Delete(id int64) error {
	rsl := config.DB.Delete(&models.Product{}, id)

	if rsl.Error != nil {
		return rsl.Error
	}

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("Not found product with id = %d", id)
	}

	return nil
}

func (service Service) UpdateFieldsById(id int64, fields map[string]interface{}) {
	rsl := config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(fields)

	if rsl.Error != nil {
		fmt.Println(rsl.Error)
	}

	if rsl.RowsAffected != 1 {
		fmt.Printf("Can`t update product with с id = %d\n", id)
	}
}

func validateProduct(product models.Product) error {
	err := services.Validate.Struct(product)
	if err != nil {
		strBuilder := strings.Builder{}
		validationErrors := err.(validator.ValidationErrors)

		for _, errVal := range validationErrors {
			strBuilder.WriteString(fmt.Sprintf("%s %s", errVal.Field(), errVal.Tag()))
		}
		return fmt.Errorf(strBuilder.String())
	}

	return nil
}
