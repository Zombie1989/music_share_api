package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Config struct holds all environment variables
type Config struct {
	SpotifyClientID     string
	SpotifyClientSecret string
	BaseURL             string
}

// Singleton pattern to load the config once
var (
	cfg  *Config
	once sync.Once
)

// LoadConfig loads the environment variables once
func LoadConfig() *Config {
	once.Do(func() {
		// Load .env file (ignore error if file is missing)
		_ = godotenv.Load()

		cfg = &Config{
			SpotifyClientID:     getEnv("SPOTIFY_CLIENT_ID", ""),
			SpotifyClientSecret: getEnv("SPOTIFY_CLIENT_SECRET", ""),
			BaseURL:             getEnv("BASE_URL", "http://localhost:8080"),
		}
		println(cfg)
		log.Println("Config loaded successfully")
	})

	return cfg
}

// getEnv retrieves environment variables or returns a fallback value
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
