package spotifyauth

import (
	"net/http"
)

// ValidateSpotifyToken checks if a token is valid by making a request to Spotify API
func ValidateSpotifyToken(token string) bool {
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// If Spotify API returns 200, the token is valid
	println(resp.StatusCode)
	return resp.StatusCode == http.StatusOK
}
