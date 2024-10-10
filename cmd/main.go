package main

import (
	"api-mongo-go/internal/core/providers"
	"api-mongo-go/internal/routes"
	"api-mongo-go/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup database connection
	db, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup router
	router := gin.Default()

	// Create API provider
	apiProvider := providers.NewApiProvider(db, router)

	// Setup routes
	routes.SetupRoutes(apiProvider)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
