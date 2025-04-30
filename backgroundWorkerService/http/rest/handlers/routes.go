package handlers

import (
	"backgroundWorkerService/internal/servicesStorage"
	"github.com/gin-gonic/gin"
)

func Register(gin *gin.Engine, services servicesStorage.InternalServices) {
	usdRatesHandler := NewHandler(services)

	gin.GET("/rates/:id", usdRatesHandler.GetUsdRates)
}
