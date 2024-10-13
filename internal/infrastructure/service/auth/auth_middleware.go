package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		token := strings.Split(authHeader, "Bearer ")[1]
		claims, err := ValidateJWT(token)
		if err != nil {
			// audit.LogRequest(c.Request, "Unauthorized") TODO implement audit
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the claims in the context for use in subsequent handlers
		c.Set("claims", claims)

		// Proceed to the next handler
		c.Next()
	}
}
