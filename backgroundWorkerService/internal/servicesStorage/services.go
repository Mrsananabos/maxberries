package servicesStorage

import (
	"backgroundWorkerService/configs"
	httpClient "backgroundWorkerService/http/rest/client"
	deliveryTariffRepo "backgroundWorkerService/internal/deliveryTariffGrid/repository"
	deliveryTariff "backgroundWorkerService/internal/deliveryTariffGrid/service"
	orderPrice "backgroundWorkerService/internal/orderPrice/service"
	usdRates "backgroundWorkerService/internal/usdRates/service"
	db "backgroundWorkerService/pkg/db/gorm"
	redisCl "backgroundWorkerService/pkg/db/redisClient"
)

type ServicesStorage struct {
	OrderPriceService     orderPrice.Service
	USDRatesService       usdRates.Service
	DeliveryTariffService deliveryTariff.Service
}

func NewServicesStorage(cnf configs.Config) (ServicesStorage, error) {
	var storage ServicesStorage

	database, err := db.Connect(cnf.Database)
	if err != nil {
		return storage, err
	}

	redisClient, err := redisCl.Connect(cnf.Redis)
	if err != nil {
		return storage, err
	}

	deliveryTariffService := deliveryTariff.NewService(deliveryTariffRepo.NewRepository(database))
	usdRatesService := usdRates.NewService(cnf, redisClient)
	return ServicesStorage{
		OrderPriceService:     orderPrice.NewService(deliveryTariffService, usdRatesService, httpClient.NewHttpClient(cnf.Services)),
		USDRatesService:       usdRatesService,
		DeliveryTariffService: deliveryTariffService,
	}, nil
}
