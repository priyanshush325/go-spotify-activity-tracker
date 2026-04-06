package utils

import (
	 "bytes"
		"encoding/json"
		"fmt"
		"net/http"
)

// Utility function to make Post Requests a lot easier. 
func RequestPost(requestUrl string, requestBody map[string]string) (*http.Response, error) {
	jsonData, _ := json.Marshal(requestBody)

	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return resp, nil
}
