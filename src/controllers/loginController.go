package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginController handles login-related requests
type LoginController struct{}

// NewLoginController creates a new LoginController instance

func NewLoginController() *LoginController {
	return &LoginController{}
}

// Login handles POST requests for user login
func (lc *LoginController) Login(c *gin.Context) {
	var loginData map[string]interface{}

	// Bind the incoming JSON to loginData
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Simulate login logic (e.g., checking credentials)
	if loginData["username"] == "user" && loginData["password"] == "pass" {
		c.JSON(http.StatusOK, gin.H{"status": "logged in"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	}
}
