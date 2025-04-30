package handlers

import (
	"backgroundWorkerService/internal/servicesStorage"
	"backgroundWorkerService/internal/usdRates/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Handler struct {
	service service.Service
}

func NewHandler(services servicesStorage.InternalServices) Handler {
	return Handler{
		service: services.USDRatesService,
	}
}

func (h Handler) GetUsdRates(c *gin.Context) {
	currency := strings.ToUpper(c.Param("id"))
	rate, err := h.service.GetUSDRate(c, currency)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"rate": rate})
		return

	}

	c.JSON(http.StatusOK, gin.H{"rate": rate})
}
