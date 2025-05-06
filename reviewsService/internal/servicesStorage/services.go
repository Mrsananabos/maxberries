package servicesStorage

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"reviewsService/configs"
	authClient "reviewsService/http/rest/client/auth"
	productClient "reviewsService/http/rest/client/product"
	reviewRepo "reviewsService/internal/review/repository"
	reviewServ "reviewsService/internal/review/service"
)

type ServicesStorage struct {
	ReviewService     reviewServ.Service
	ProductHttpClient productClient.HttpClient
	AuthHttpClient    authClient.HttpClient
}

func NewServicesStorage(config configs.Config, mongo *mongo.Collection) ServicesStorage {
	return ServicesStorage{
		ReviewService:     reviewServ.NewService(reviewRepo.NewRepository(mongo)),
		ProductHttpClient: productClient.NewHttpClient(config.Services),
		AuthHttpClient:    authClient.NewHttpClient(config.Services, http.Client{}),
	}

}
