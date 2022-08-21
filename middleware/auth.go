package middleware

import (
	"net/http"

	"github.com/Imprezzaa/linkshortener/auth"
	"github.com/gin-gonic/gin"
)

// Auth provides middleware to intercept connections to protected routes to check a user is verified
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "request does not contain an access token",
			})
			return
		}

		err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Next()
	}
}
