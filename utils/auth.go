package utils 

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"io"
	"encoding/json"
)

var authLogger *log.Logger

func init() {
	err := godotenv.Load()

	authLogger = log.New(os.Stdout, "[AUTH] ", log.LstdFlags)

	if err != nil {
		authLogger.Fatalf("Error loading .env: %v", err)
	}
}

// Helper function that makes a request to Spotify's token endpoint to exchange a refresh token for a new access token.
func refreshAccessToken() (string, error) {
	requestBody := map[string]string {
		"grant_type": "refresh_token",
		"refresh_token": os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		"client_id": os.Getenv("CLIENT_ID"),
		"client_secret": os.Getenv("CLIENT_SECRET"),
	}

	authLogger.Println("Refreshing Spotify Access Token")

	resp, err := RequestPost("https://accounts.spotify.com/api/token", requestBody)

	if err != nil {
		authLogger.Println("Error Refreshing Access Token: %v", err)
		return "", err
	}

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		authLogger.Println("Error reading refresh response body: %v", err)
		return "", err
	}

	var responseResult map[string]interface{}
	json.Unmarshal(responseBody, &responseResult)

	return responseResult["access_token"].(string), nil



}
