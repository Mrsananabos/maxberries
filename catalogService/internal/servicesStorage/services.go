package servicesStorage

import (
	"catalogService/configs"
	"catalogService/http/rest/client/auth"
	categoryRepo "catalogService/internal/category/repository"
	category "catalogService/internal/category/service"
	productRepo "catalogService/internal/product/repository"
	product "catalogService/internal/product/service"
	redisCl "catalogService/pkg/db/redisClient"
	"gorm.io/gorm"
	"net/http"
)

type ServicesStorage struct {
	CategoryService category.Service
	ProductService  product.Service
	AuthHttpClient  auth.HttpClient
}

func NewServicesStorage(cnf configs.Config, db *gorm.DB) (ServicesStorage, error) {
	var storage ServicesStorage
	redisClient, err := redisCl.Connect(cnf.Redis)
	if err != nil {
		return storage, err
	}
	categoryService := category.NewService(categoryRepo.NewRepository(db), redisClient)

	storage = ServicesStorage{
		CategoryService: categoryService,
		ProductService:  product.NewService(productRepo.NewRepository(db), categoryService),
		AuthHttpClient:  auth.NewHttpClient(cnf.Services, http.Client{}),
	}

	return storage, nil
}
