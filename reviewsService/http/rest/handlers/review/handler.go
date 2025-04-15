package review

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"reviewsService/internal/review/repository"
	review "reviewsService/internal/review/service"
)

type Handler struct {
	service review.Service
}

func NewHandler(db *mongo.Client) Handler {
	return Handler{
		service: review.NewService(repository.NewRepository(db)),
	}
}

func (h Handler) GetByProductId(c *gin.Context) {
	productId := c.Param("id")

	reviews, err := h.service.GetByProductId(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, reviews)
}

//
//func (h Handler) CreateReview(c *gin.Context) {
//	var reviewRequest bson.M
//
//	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err := h.service.Create(categoryRequest)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
//}
//
//func (h Handler) DeleteCategory(c *gin.Context) {
//	idStr := c.Param("id")
//	id, err := strconv.ParseInt(idStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
//		return
//	}
//
//	err = h.service.Delete(id)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
//}
//
//func (h Handler) UpdateCategory(c *gin.Context) {
//	idStr := c.Param("id")
//	id, err := strconv.ParseInt(idStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
//		return
//	}
//
//	var updatedCategory model.Category
//	if err = c.ShouldBindJSON(&updatedCategory); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	updatedCategory.Id = id
//
//	err = h.service.Update(updatedCategory)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
//}
