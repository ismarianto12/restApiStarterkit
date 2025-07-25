package utils

import (
	"errors"
	"fmt"
	"golangRest/src/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(username string, level_id string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("failed to load env: %w", err)
	}

	secretKey := os.Getenv("APP_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("APP_SECRET_KEY environment variable not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"level_id": level_id,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	// Sign the token with the secret key (converted to []byte)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func DecodeToken(tokenString string) (*models.UserClaims, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load env: %w", err)
	}

	secretKey := os.Getenv("APP_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("APP_SECRET_KEY environment variable not set")
	}

	// Parse the token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract and return the claims
	if claims, ok := token.Claims.(*models.UserClaims); ok {
		return claims, nil
	}

	return nil, errors.New("failed to extract claims from token")
}
