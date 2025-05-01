package middleware

import (
	"authService/internal/auth/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Middleware struct {
	authService service.Service
}

func CreateJWTMiddleware(authService service.Service) Middleware {
	return Middleware{
		authService: authService,
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

		claims, err := m.authService.Auth(authToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		for _, perm := range claims.Permissions {
			if perm == requiredPermission {
				c.Next()
				return
			}
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

		claims, err := m.authService.Auth(authToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		for _, perm := range claims.Permissions {
			if perm == requiredPermission {
				c.Next()
				return
			}
		}

		userUUID := c.Param("id")
		sub, err := claims.GetSubject()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}

		if userUUID == sub {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}
