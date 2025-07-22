package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InterCeptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("SecureToken")
		if token != "S" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "forbiden invalid secure",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
