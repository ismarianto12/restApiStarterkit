package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "root:root@tcp(localhost:3308)/golang_rest?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// ⛳️ ini penting: JANGAN pakai :=
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	fmt.Println("Database connected successfully")
	return nil
}
