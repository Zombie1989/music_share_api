package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserData struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	ID          string `json:"id"`
	Country     string `json:"country"`
}

// GetCurrentUser fetches the current user details from Spotify API
func GetSpotifyUserData(token string) (*UserData, error) {
	// Create the HTTP request to the Spotify API
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return nil, err
	}

	// Set the Authorization header with the Bearer token
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a new HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Spotify API returned status: %d", resp.StatusCode)
	}

	// Decode the JSON response into the UserData struct
	var userData UserData
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}
