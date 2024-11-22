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

	// Userテーブルを先に作成する必要がある
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate User: %v", err)
	}

	// FavoriteStoreテーブルを作成
	if err := DB.AutoMigrate(&FavoriteStore{}); err != nil {
		log.Fatalf("failed to migrate FavoriteStore: %v", err)
	}
}
