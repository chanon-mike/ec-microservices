package main

import (
	"context"
	"log"
	"os"

	"github.com/chanon-mike/ec-microservices/config"
	"github.com/chanon-mike/ec-microservices/pkg/database"
	"github.com/chanon-mike/ec-microservices/server"
)

func main() {
	ctx := context.Background()

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	// Database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, &cfg, db)
}
