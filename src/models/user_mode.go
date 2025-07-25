package models

import "github.com/golang-jwt/jwt/v5"

type UserMode struct {
	ID       uint   `gorm:"primaryKey",json:"id"`
	Name     string `gorm:"not null",json:"name"`
	Username string `gorm:"not null",json:"username"`
	Password string `gorm:"not null",json:"password"`
	Image    string `gorm:"not null",json:"image"`
	Email    string `gorm:"unique;not null",json:"email"`
	LevelId  string `gorm:"unique;not null",json:"level_id"`
}

// TableName returns the name of the table in the database
func (UserMode) TableName() string {
	return "users"
}

type UserClaims struct {
	UserID  uint   `json:"user_id"`
	LevelID string `json:"level_id"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}
