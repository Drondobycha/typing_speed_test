package main

import (
	"log"
	"os"
	"typing-speed-test/api"
	"typing-speed-test/database"

	"github.com/gin-gonic/gin"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	store, err := database.Newstore(connStr)
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL store: %v", err)
	}
	router := gin.Default()

	api.SetupRoutes(router, store)
	router.Run(":3000")
}
