package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() {
	envPath := ".env"
	// debug時の.envファイルのパス指定
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		envPath = "../.env"
	}
	if err := godotenv.Load(envPath); err != nil {
		fmt.Printf("Cannot read: %v", err)
	}
	dsn := os.Getenv("TESTDB_CONNECTION")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// ここで使用するすべての構造体をMigrate
	DB.AutoMigrate(&Store{})
	DB.AutoMigrate(&User{})
}