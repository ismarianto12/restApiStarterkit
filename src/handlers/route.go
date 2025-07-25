package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

}

func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to API"})
}
