package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"orderService/http/rest/client/auth"
	"orderService/internal/order/service"
	"orderService/internal/servicesStorage"
	"strconv"
)

type Middleware struct {
	orderService service.Service
	authClient   auth.HttpClient
}

func CreateJWTMiddleware(services servicesStorage.ServicesStorage) Middleware {
	return Middleware{
		orderService: services.OrderService,
		authClient:   services.AuthHttpClient,
	}
}

func (m Middleware) PermissionCheckMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		claims, err := m.authClient.GetAuthClaims(authToken)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if hasPermission := slices.Contains(claims.Permissions, requiredPermission); hasPermission {
			c.Request.Header.Set("userId", claims.Sub)
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}

func (m Middleware) UserPermissionCheckMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		claims, err := m.authClient.GetAuthClaims(authToken)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		hasPermission := slices.Contains(claims.Permissions, requiredPermission)
		if hasPermission {
			if claims.Role == "admin" {
				c.Next()
				return
			}

			orderIdStr := c.Param("id")
			orderId, err := strconv.ParseInt(orderIdStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
				return
			}

			order, err := m.orderService.GetById(orderId)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			if claims.Sub == order.UserId.String() {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}
