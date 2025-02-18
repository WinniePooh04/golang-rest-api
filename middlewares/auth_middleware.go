package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PSKMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		psk := c.GetHeader("Authorization")

		// Check if PSK matches the expected key
		if psk != os.Getenv("PRE_SHARED_KEY") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing Pre Shared key"})
			c.Abort()
			return
		}

		// Continue processing request
		c.Next()
	}
}
