package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(username string) (string, error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("failed to load env: %w", err)
	}

	// Get secret key from environment
	secretKey := os.Getenv("APP_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("APP_SECRET_KEY environment variable not set")
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	// Sign the token with the secret key (converted to []byte)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}
