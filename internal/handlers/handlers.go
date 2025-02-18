package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Zombie1989/music_share_api/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var cfg = config.LoadConfig()

var oauthConf = &oauth2.Config{
	ClientID:     cfg.SpotifyClientID,
	ClientSecret: cfg.SpotifyClientSecret,
	RedirectURL:  fmt.Sprintf("%s/callback", cfg.BaseURL),
	Scopes:       []string{"user-read-private", "user-read-email"},
	Endpoint:     spotify.Endpoint,
}

// Login redirects users to the Spotify authentication page
func Login(c *gin.Context) {
	authURL := oauthConf.AuthCodeURL("state")
	c.Redirect(http.StatusFound, authURL)
}

// Callback handles the OAuth2 response from Spotify
func Callback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not found"})
		return
	}

	// Exchange the authorization code for an access token
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Error exchanging token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token.AccessToken,
		"token_type":   token.TokenType,
		"expires_in":   token.Expiry,
	})
}
