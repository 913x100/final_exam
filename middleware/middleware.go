package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "mytokenbase" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access denied"})
		c.Abort()
		return
	}

	c.Next()
}
