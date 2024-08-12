package main

import (
	"clean-storemap-api/src/driver/db"
	router "clean-storemap-api/src/driver/router"
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	routerI, err := router.InitializeRouter(ctx)
	if err != nil {
		fmt.Printf("failed to create Router: %s\n", err)
		os.Exit(2)
	}
	db.InitDB()
	routerI.Serve(ctx)
}
