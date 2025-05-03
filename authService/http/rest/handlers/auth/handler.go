package auth

import (
	"authService/internal/auth/model"
	"authService/internal/auth/service"
	"authService/internal/servicesStorage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service service.Service
}

func NewHandler(services servicesStorage.ServicesStorage) Handler {
	return Handler{
		service: services.AuthService,
	}
}

func (h Handler) Register(c *gin.Context) {
	var regForm model.Register
	if err := c.ShouldBindJSON(&regForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Register(regForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully register %s", regForm.Username)})
}

func (h Handler) Login(c *gin.Context) {
	var loginForm model.Login

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Login(c, loginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h Handler) Auth(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	claims, err := h.service.Auth(authToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, claims.Permissions)
}

func (h Handler) RefreshTokens(c *gin.Context) {
	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	tokens, err := h.service.RefreshTokens(c, refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}
