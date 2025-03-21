package main

import (
	"log"
	"os"
	"typing-speed-test/api"
	"typing-speed-test/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
		"@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"
	store, err := database.Newstore(connStr)
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL store: %v", err)
	}
	router := gin.Default()
	port := os.Getenv("PORT")
	api.SetupRoutes(router, store)
	log.Printf("Server runnig on port: %s", port)
	router.Run(":" + port)
}
