package handlers

import (
	"net/http"

	"github.com/Zombie1989/music_share_api/internal/services"
	"github.com/gin-gonic/gin"
)

// GetMe fetches the current user from the Spotify API using the access token
func GetMe(c *gin.Context) {
	// Retrieve the access token from the Authorization header
	token := c.GetHeader("Authorization")
	// Remove the "Bearer " prefix from the token
	accessToken := token[len("Bearer "):]
	userData, err := services.GetSpotifyUserData(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	// Send the JSON response to the client
	c.JSON(http.StatusOK, userData)

}
