package models

import (
	"gorm.io/gorm"
)

// UserMode represents the user model
type UserMode struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// TableName returns the name of the table in the database
func (UserMode) TableName() string {
	return "users"
}
