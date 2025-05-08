package handlers

import (
	"github.com/gin-gonic/gin"
	"orderService/http/rest/handlers/order"
	orderItems "orderService/http/rest/handlers/orderItems"
	"orderService/http/rest/middleware"
	services "orderService/internal/servicesStorage"
)

func Register(gin *gin.Engine, services services.ServicesStorage) {
	orderHandler := order.NewHandler(services.OrderService)
	orderItemsHandler := orderItems.NewHandler(services.OrderItemsService)
	m := middleware.CreateJWTMiddleware(services)

	gin.GET("/orders", m.PermissionCheckMiddleware("orders.getAll"), orderHandler.GetAllOrders)
	gin.GET("/orders/:id", m.UserPermissionCheckMiddleware("orders.getById"), orderHandler.GetOrderById)
	gin.POST("/orders", m.PermissionCheckMiddleware("orders.createOrder"), orderHandler.CreateOrder)
	gin.PUT("/orders/:id/items", m.UserPermissionCheckMiddleware("orders.putOrder"), orderItemsHandler.UpdateItems)
	gin.PATCH("/orders/:id", m.PermissionCheckMiddleware("orders.patchOrder"), orderHandler.UpdateOrder)
	gin.DELETE("/orders/:id", m.PermissionCheckMiddleware("orders.deleteOrder"), orderHandler.DeleteOrder)
}
