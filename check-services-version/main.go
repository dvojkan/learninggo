package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServiceVersion struct {
	ServiceName string `json:"serviceName"`
	Version     string `json:"version"`
}

func main() {
	// URLs to fetch data from
	dev_urls := []string{
		//here add URLs
	}

	// Loop through each URL
	for _, url := range dev_urls {
		// Make HTTP GET request
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching URL %s: %v\n", url, err)
			continue
		}
		defer response.Body.Close()

		// Read response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading response body from %s: %v\n", url, err)
			continue
		}

		// Parse JSON response
		var ServiceVersion ServiceVersion
		err = json.Unmarshal(body, &ServiceVersion)
		if err != nil {
			fmt.Printf("Error parsing JSON from %s: %v\n", url, err)
			continue
		}

		// Print parsed data
		fmt.Printf("Servic name: %s\n", ServiceVersion.ServiceName)
		fmt.Printf("Version: %s\n", ServiceVersion.Version)
		fmt.Println("--------------------------")
	}
}
