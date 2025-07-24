package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func validateToken(token string) (*jwt.Token, error) {
	errenv := godotenv.Load()
	if errenv != nil {
		fmt.Errorf("failed load env: %v", errenv)
	}
	secretKey := os.Getenv("APP_SECRET_KEY")
	inputtoken, err := jwt.Parse(token, func(inputtoken *jwt.Token) (interface{}, error) {
		if _, ok := inputtoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Kontol bapak kau pecah pusing pala gw %v")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	return inputtoken, nil

}
