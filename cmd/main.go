package main

import (
	"clean-storemap-api/src/driver/db"
	router "clean-storemap-api/src/driver/router"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envPath := ".env"
	// debug時の.envファイルのパス指定
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		envPath = "../.env"
	}
	if err := godotenv.Load(envPath); err != nil {
		fmt.Printf("Cannot read: %v", err)
	}

	ctx := context.Background()
	routerI, err := router.InitializeRouter(ctx)
	if err != nil {
		fmt.Printf("failed to create Router: %s\n", err)
		os.Exit(2)
	}
	db.InitDB()
	routerI.Serve(ctx)
}
