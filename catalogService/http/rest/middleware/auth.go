package middleware

import (
	"catalogService/http/rest/client/auth"
	"catalogService/internal/servicesStorage"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
)

type Middleware struct {
	authClient auth.HttpClient
}

func CreateJWTMiddleware(services servicesStorage.ServicesStorage) Middleware {
	return Middleware{
		authClient: services.AuthHttpClient,
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if hasPermission := slices.Contains(claims.Permissions, requiredPermission); hasPermission {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}
