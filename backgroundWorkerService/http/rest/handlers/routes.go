package handlers

import (
	services "backgroundWorkerService/internal/servicesStorage"
	"github.com/gin-gonic/gin"
)

func Register(gin *gin.Engine, services services.ServicesStorage) {
	usdRatesHandler := NewHandler(services)

	gin.GET("/rates/:id", usdRatesHandler.GetUsdRates)
}
