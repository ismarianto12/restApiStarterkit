package utils

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func AuthMiddleware(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to load environment variables",
			"details": err.Error(),
		})
		return
	}

	secretKey := os.Getenv("APP_SECRET_KEY")
	if secretKey == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "APP_SECRET_KEY not configured",
		})
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header required",
			"code":  http.StatusUnauthorized,
		})
		return
	}

	// bearer prefix
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// aprsing  validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		// Convert secret key to []byte
		return []byte(secretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid token",
			"details": err.Error(),
		})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not valid",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("username", claims["username"])
	}

	c.Next()
}
