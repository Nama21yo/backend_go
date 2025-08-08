package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Context keys
const (
	ContextUserIDKey = "current_user_id"
	ContextUserRole   = "current_user_role"
)

// AuthMiddleware returns a Gin middleware that validates JWT tokens via provided JWTService
func AuthMiddleware(jwtSvc JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
			return
		}
		token := parts[1]
		claims, err := jwtSvc.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token: " + err.Error()})
			return
		}
		// inject into context
		c.Set(ContextUserIDKey, claims.UserID)
		c.Set(ContextUserRole, claims.Role)
		c.Next()
	}
}
