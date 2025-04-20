package handlers

import (
	"backgroundWorkerService/configs"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Register(gin *gin.Engine, config configs.Config, redis *redis.Client) {
	usdRatesHandler := NewHandler(config, redis)

	gin.GET("/rates/:id", usdRatesHandler.GetUsdRates)
}
