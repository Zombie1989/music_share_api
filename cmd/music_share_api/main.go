package main

import (
	"github.com/Zombie1989/music_share_api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"}, // Allow the frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, // Allow credentials (cookies, tokens, etc.)
	}))

	// Set up all routes
	routes.SetupRoutes(router)

	// Start the Gin server
	router.Run(":8080")
}
