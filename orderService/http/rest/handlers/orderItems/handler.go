package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orderService/internal/orderItem/model"
	"orderService/internal/orderItem/service"
	"strconv"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) UpdateItems(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedItems model.EditOrderItemsRequest
	if err = c.ShouldBindJSON(&updatedItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = updatedItems.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateItems(id, updatedItems)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order items updated successfully"})
}
