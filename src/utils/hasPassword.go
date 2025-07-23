package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Haspassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
