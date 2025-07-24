package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func generateToke(username string) (string, error) {
	errenv := godotenv.Load()
	if errenv != nil {
		fmt.Errorf("failed load env: %v", errenv)
	}
	secretKey := os.Getenv("APP_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
