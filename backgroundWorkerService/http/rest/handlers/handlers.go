package handlers

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/internal/usdRates/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type Handler struct {
	service service.Service
}

func NewHandler(config configs.Config, r *redis.Client) Handler {
	return Handler{
		service: service.NewService(config, r),
	}
}

func (h Handler) GetUsdRates(c *gin.Context) {
	usdRates, err := h.service.GetUSDRates(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, usdRates.Rates)
}
