package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"orderService/configs"
	"orderService/http/rest/handlers/order"
)

func Register(gin *gin.Engine, db *gorm.DB, cnf configs.Services) {
	orderHandler := order.NewHandler(db, cnf)

	gin.GET("/orders", orderHandler.GetAllOrders)
	gin.GET("/orders/:id", orderHandler.GetOrderById)
	gin.POST("/orders", orderHandler.CreateOrder)
	gin.PATCH("/orders/:id", orderHandler.UpdateOrderStatus)
	gin.DELETE("/orders/:id", orderHandler.DeleteOrder)
}
