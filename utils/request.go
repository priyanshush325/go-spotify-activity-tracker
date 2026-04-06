package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Utility function to make Post Requests a lot easier. 
func RequestPost(requestUrl string, requestBody map[string]string) (*http.Response, error) {
	formData := url.Values{}
	for key, value := range requestBody {
		formData.Set(key, value)
	}

	resp, err := http.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return resp, nil
}

// Utility function to make Get Requests a lot easier. 

func RequestGet(requestUrl string, requestHeaders map[string]string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req,_ := http.NewRequest("GET", requestUrl, nil)

	for key, value := range (requestHeaders) {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return resp, nil
	
}
