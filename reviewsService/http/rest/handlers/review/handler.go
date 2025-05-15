package review

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"reviewsService/http/rest/client/product"
	"reviewsService/internal/review/model"
	review "reviewsService/internal/review/service"
	"reviewsService/internal/servicesStorage"
	"reviewsService/pkg/kafka"
	"reviewsService/pkg/kafka/message"
	"strconv"
)

type Handler struct {
	service       review.Service
	httpClient    product.HttpClient
	kafkaProducer kafka.Producer
}

func NewHandler(kafkaProducer kafka.Producer, services servicesStorage.ServicesStorage) Handler {
	return Handler{
		service:       services.ReviewService,
		httpClient:    services.ProductHttpClient,
		kafkaProducer: kafkaProducer,
	}
}

func (h Handler) GetByProductId(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	reviewRequest.UserID = c.Request.Header.Get("userId")

	if err := reviewRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.httpClient.GetProductById(reviewRequest.ProductID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Not found product with id = %d", reviewRequest.ProductID)})
		return
	}

	createdReview, err := h.service.Create(c, reviewRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kafkaMsg := message.ReviewCreatedMsg{Event: message.REVIEW_CREATED_EVENT, ID: createdReview.ID, ProductID: createdReview.ProductID,
		UserID: createdReview.UserID, Rating: createdReview.Rating, Text: createdReview.Text}
	_, err = h.kafkaProducer.SentMsg(kafka.REVIEW_EVENTS_TOPIC, kafkaMsg)
	if err != nil {
		log.Printf("kafka: failed send order created msg %s", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review created successfully"})
}

func (h Handler) DeleteByProductId(c *gin.Context) {
	productIdStr := c.Query("product_id")
	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.DeleteByProductId(c, productId)
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
