package handlers

import (
	"catalogService/http/rest/handlers/category"
	product "catalogService/http/rest/handlers/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(gin *gin.Engine, db *gorm.DB) {
	categoryHandler := category.NewHandler(db)
	productHandler := product.NewHandler(db)

	gin.GET("/categories", categoryHandler.GetAllCategories)
	gin.GET("/categories/:id", categoryHandler.GetCategoryById)
	gin.POST("/categories", categoryHandler.SaveCategory)
	gin.PUT("/categories/:id", categoryHandler.UpdateCategory)
	gin.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	gin.GET("/products", productHandler.GetAllProducts)
	gin.GET("/products/:id", productHandler.GetProductById)
	gin.POST("/products", productHandler.SaveProduct)
	gin.PUT("/products/:id", productHandler.UpdateProduct)
	gin.DELETE("/products/:id", productHandler.DeleteProduct)
}
