package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	dsn := os.Getenv("TESTDB_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Store{})

	// Create
	db.Create(&Store{Code: "D42", Price: 100})

	// Read
	var store Store
	db.First(&store, 1)                 // find store with integer primary key
	db.First(&store, "Code = ?", "D42") // find store with code D42

	// Update - update store's price to 200
	db.Model(&store).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&store).Updates(Store{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&store).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete store
	// db.Delete(&store, 1)
}
