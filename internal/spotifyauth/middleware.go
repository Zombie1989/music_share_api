package spotifyauth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireAuth is a middleware that requires an access token for all routes except login and callback
func RequireSpotifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Ensure it's a Bearer token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		accessToken := tokenParts[1]

		if accessToken == "" || !ValidateSpotifyToken(accessToken) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired Spotify token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
