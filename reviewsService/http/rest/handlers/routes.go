package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"reviewsService/configs"
	"reviewsService/http/rest/handlers/review"
)

func Register(gin *gin.Engine, mCollection *mongo.Collection, cnf configs.Services) {
	reviewHandler := review.NewHandler(mCollection, cnf)

	gin.POST("/reviews/", reviewHandler.CreateReview)
	gin.GET("/reviews/:id", reviewHandler.GetByProductId)
	gin.PATCH("/reviews/:id", reviewHandler.UpdateReview)
	gin.DELETE("/reviews/:id", reviewHandler.DeleteById)
	gin.DELETE("/reviews/", reviewHandler.DeleteByProductId)
}
