package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Cannot read: %v", err)
	}
	dsn := os.Getenv("TESTDB_CONNECTION")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Store{}) // ここで使用するすべての構造体をMigrate
}
