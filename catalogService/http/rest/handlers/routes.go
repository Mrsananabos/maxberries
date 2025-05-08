package handlers

import (
	"catalogService/http/rest/handlers/category"
	product "catalogService/http/rest/handlers/product"
	"catalogService/http/rest/middleware"
	"catalogService/internal/servicesStorage"
	"github.com/gin-gonic/gin"
)

func Register(gin *gin.Engine, services servicesStorage.ServicesStorage) {
	categoryHandler := category.NewHandler(services)
	productHandler := product.NewHandler(services)
	m := middleware.CreateJWTMiddleware(services)

	gin.GET("/categories", categoryHandler.GetAllCategories)
	gin.GET("/categories/:id", categoryHandler.GetCategoryById)
	gin.POST("/categories", m.PermissionCheckMiddleware("category.create"), categoryHandler.CreateCategory)
	gin.PUT("/categories/:id", m.PermissionCheckMiddleware("category.edit"), categoryHandler.UpdateCategory)
	gin.DELETE("/categories/:id", m.PermissionCheckMiddleware("category.delete"), categoryHandler.DeleteCategory)

	gin.GET("/products", productHandler.GetAllProducts)
	gin.GET("/products/:id", productHandler.GetProductById)
	gin.POST("/products", m.PermissionCheckMiddleware("product.create"), productHandler.SaveProduct)
	gin.PUT("/products/:id", m.PermissionCheckMiddleware("product.edit"), productHandler.UpdateProduct)
	gin.DELETE("/products/:id", m.PermissionCheckMiddleware("product.delete"), productHandler.DeleteProduct)
}
