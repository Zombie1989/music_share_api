package main

import (
	"github.com/Zombie1989/music_share_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	router := gin.Default()

	// Set up all routes
	routes.SetupRoutes(router)

	// Start the Gin server
	router.Run(":8080")
}
