package handlers

import (
	"github.com/gin-gonic/gin"
	"reviewsService/http/rest/handlers/review"
	"reviewsService/http/rest/middleware"
	"reviewsService/internal/servicesStorage"
	"reviewsService/pkg/kafka"
)

func Register(gin *gin.Engine, services servicesStorage.ServicesStorage, kafkaProducer kafka.Producer) {
	reviewHandler := review.NewHandler(kafkaProducer, services)
	m := middleware.CreateJWTMiddleware(services)

	gin.POST("/reviews/", m.PermissionCheckMiddleware("review.create"), reviewHandler.CreateReview)
	gin.GET("/reviews/:id", reviewHandler.GetByProductId)
	gin.PATCH("/reviews/:id", m.UserPermissionCheckMiddleware("review.updateAny", "review.update"), reviewHandler.UpdateReview)
	gin.DELETE("/reviews/:id", m.UserPermissionCheckMiddleware("review.deleteAny", "review.delete"), reviewHandler.DeleteById)
	gin.DELETE("/reviews/", m.PermissionCheckMiddleware("review.deleteProductReviews"), reviewHandler.DeleteByProductId)
}
