package main

import (
	"log"

	// "github.com/Jay-SCM/alchemy/api/go/routes"
	"github.com/Jay-SCM/alchemy-backend/api/go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	log.Println("Starting the server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
