package handlers

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/internal/usdRates/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"strings"
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
	currency := strings.ToUpper(c.Param("id"))
	rate, err := h.service.GetUSDRatesCache(c, currency)

	if err != nil {
		log.Println(err.Error())

		usdRates, err := h.service.GetUSDRates(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		rate, ok := usdRates.Rates[currency]

		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Not found range for %s", currency)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"rate": rate})

	} else {
		c.JSON(http.StatusOK, gin.H{"rate": rate})
	}
}
