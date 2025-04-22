package review

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"reviewsService/configs"
	"reviewsService/http/rest/client"
	"reviewsService/internal/review/model"
	"reviewsService/internal/review/repository"
	review "reviewsService/internal/review/service"
)

type Handler struct {
	service    review.Service
	httpClient client.HttpClient
}

func NewHandler(mCollection *mongo.Collection, cnf configs.Services) Handler {
	return Handler{
		service:    review.NewService(repository.NewRepository(mCollection)),
		httpClient: client.NewHttpClient(cnf),
	}
}

func (h Handler) GetByProductId(c *gin.Context) {
	productId := c.Param("id")

	reviews, err := h.service.GetByProductId(c, productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h Handler) CreateReview(c *gin.Context) {
	var reviewRequest model.Review

	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := reviewRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.httpClient.GetProductById(reviewRequest.ProductID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Not found product with id = %s", reviewRequest.ProductID)})
		return
	}

	err := h.service.Create(c, reviewRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review created successfully"})
}

func (h Handler) DeleteByProductId(c *gin.Context) {
	productID := c.Query("product_id")
	err := h.service.DeleteByProductId(c, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reviews deleted successfully"})
}

func (h Handler) DeleteById(c *gin.Context) {
	idStr := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error converting %s to ObjectID", idStr)})
		return
	}

	err = h.service.DeleteById(c, objectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

func (h Handler) UpdateReview(c *gin.Context) {
	idStr := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error converting %s to ObjectID", idStr)})
		return
	}

	var updatedReview model.ContentReview
	if err = c.ShouldBindJSON(&updatedReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Update(c, objectID, updatedReview)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}
