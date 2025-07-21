package models

import (
	"gorm.io/gorm"
)

type UserMode struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey",json:"id"`
	Name     string `gorm:"not null",json:"name"`
	Email    string `gorm:"unique;not null",json:"email"`
	Password string `gorm:"not null",json:"password"`
}

// TableName returns the name of the table in the database
func (UserMode) TableName() string {
	return "users"
}
