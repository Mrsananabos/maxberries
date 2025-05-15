package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"reviewsService/http/rest/client/auth"
	"reviewsService/internal/review/model"
	"reviewsService/internal/review/service"
	"reviewsService/internal/servicesStorage"
)

type Middleware struct {
	reviewService service.Service
	authClient    auth.HttpClient
}

func CreateJWTMiddleware(services servicesStorage.ServicesStorage) Middleware {
	return Middleware{
		reviewService: services.ReviewService,
		authClient:    services.AuthHttpClient,
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

func (m Middleware) UserPermissionCheckMiddleware(fullPermission string, userPermission string) gin.HandlerFunc {
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

		hasFullPermission := slices.Contains(claims.Permissions, fullPermission)
		if hasFullPermission {
			c.Next()
			return
		}

		hasUserPermission := slices.Contains(claims.Permissions, userPermission)
		if hasUserPermission {
			reviewId := c.Param("id")
			userReviews, err := m.reviewService.GetByUserId(c, claims.Sub)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			if hasReviewById := slices.ContainsFunc(userReviews, func(review model.Review) bool {
				return review.ID.Hex() == reviewId
			}); hasReviewById {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}
