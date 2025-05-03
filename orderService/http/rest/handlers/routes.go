package handlers

import (
	"github.com/gin-gonic/gin"
	"orderService/http/rest/handlers/order"
	orderItems "orderService/http/rest/handlers/orderItems"
	services "orderService/internal/servicesStorage"
)

func Register(gin *gin.Engine, services services.ServicesStorage) {
	orderHandler := order.NewHandler(services.OrderService)
	orderItemsHandler := orderItems.NewHandler(services.OrderItemsService)

	gin.GET("/orders", orderHandler.GetAllOrders)
	gin.GET("/orders/:id", orderHandler.GetOrderById)
	gin.POST("/orders", orderHandler.CreateOrder)
	gin.PUT("/orders/:id/items", orderItemsHandler.UpdateItems)
	gin.PATCH("/orders/:id", orderHandler.UpdateOrder)
	gin.DELETE("/orders/:id", orderHandler.DeleteOrder)
}
