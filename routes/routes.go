package routes

import (
	"github.com/Zombie1989/music_share_api/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers API endpoints
func SetupRoutes(router *gin.Engine) {
	router.GET("/login", handlers.Login)
	router.GET("/callback", handlers.Callback)
}
