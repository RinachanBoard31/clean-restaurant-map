package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("TESTDB_CONNECTION")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// ここで使用するすべての構造体をMigrate
	DB.AutoMigrate(&FavoriteStore{})
	DB.AutoMigrate(&User{})
}
