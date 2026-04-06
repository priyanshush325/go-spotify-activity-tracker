package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var accessToken string

type Artist struct {
	Name string `json:"name"`
}

type AlbumImage struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Album struct {
	Name   string       `json:"name"`
	Images []AlbumImage `json:"images"`
}

type Track struct {
	Name    string   `json:"name"`
	Artists []Artist `json:"artists"`
	Album   Album    `json:"album"`
}

type RecentlyPlayedItem struct {
	Track    Track  `json:"track"`
	PlayedAt string `json:"played_at"`
}

type RecentlyPlayedResponse struct {
	Items []RecentlyPlayedItem `json:"items"`
}


var spotifyLogger *log.Logger

func init() {
	spotifyLogger = log.New(os.Stdout, "[SPOTIFY] ", log.LstdFlags)
}

//Gets the user's most recently listened to song from Spotify, using the access token and refreshing it if necessary. 
func GetLastListened () (Track, error){
	spotifyLogger.Println("Calling GetLastListened func")

	// Check if the access token is currently initialized. If not, refresh it first
	if accessToken == "" {
		refreshedAccessToken, err := refreshAccessToken()

		if err != nil {
			spotifyLogger.Println("Error refreshing Access Token")
			return Track{}, err
		}

		accessToken = refreshedAccessToken
	}

	// Use the Access token to call the Spotify endpoint for retrieving the user's most recently listened song.
	headers := map[string]string {
	"Authorization": "Bearer " + accessToken,
	}

	resp, err := RequestGet("https://api.spotify.com/v1/me/player/recently-played?limit=1", headers)

	if err != nil {
		spotifyLogger.Println("Error requesting last played song: %v", err)
		return Track{}, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		spotifyLogger.Println("Error reading response body: %v", err)
		return Track{}, err
	}

	var responseResult RecentlyPlayedResponse
	json.Unmarshal(responseBody, &responseResult)

	if len(responseResult.Items) == 0 {
		spotifyLogger.Println("No recently played tracks found")
		return Track{}, nil
	}

	return responseResult.Items[0].Track, nil
}
