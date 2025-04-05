package routes

import (
	"catalogService/routes/handlers/category"
	product "catalogService/routes/handlers/product"
	"github.com/gin-gonic/gin"
)

func LoadRoutes() {
	r := gin.Default()

	r.GET("/categories", category.GetAllCategories)
	r.GET("/categories/:id", category.GetCategoryById)
	r.POST("/categories", category.SaveCategory)
	r.PUT("/categories/:id", category.UpdateCategory)
	r.DELETE("/categories/:id", category.DeleteCategory)

	r.GET("/products", product.GetAllProducts)
	r.GET("/products/:id", product.GetProductById)
	r.POST("/products", product.SaveProduct)
	r.PUT("/products/:id", product.UpdateProduct)
	r.DELETE("/products/:id", product.DeleteProduct)

	r.Run(":8080")
}
