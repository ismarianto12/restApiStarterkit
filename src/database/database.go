package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:root@tcp(localhost:3308)/golang_rest?charset=utf8mb4&parseTime=True&loc=Local"

	// 1. Buka koneksi SQL biasa
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}

	// 2. Cek apakah koneksi bisa digunakan
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping the database: %v", err)
	}

	fmt.Println("✅ Database connection established successfully")

	// 3. Inisialisasi GORM dengan koneksi yang sudah dicek
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to initialize GORM: %v", err)
	}

	// 4. Simpan DB ke variabel global
	DB = gormDB

	fmt.Println("✅ GORM initialized successfully")
}
