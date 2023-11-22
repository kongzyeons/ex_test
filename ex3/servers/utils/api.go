package utils

import (
	"fmt"
	"go_ex3/models"
	"io"
	"net/http"
	"time"
)

func CallApiGetBody(url string) (result string, err error) {
	// url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	// Make an HTTP GET request to the URL
	maxRetries := 3
	for retry := 0; retry < maxRetries; retry++ {
		// Make an HTTP GET request to the URL
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making HTTP request: %v\n", err)
			return result, models.ErrHttpRequest
		}
		defer response.Body.Close()

		// Check if the response status code is OK (200)
		if response.StatusCode == http.StatusOK {
			// Read the response result
			body, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("Error reading response result: %v\n", err)
				return result, models.ErrReadResult
			}
			return string(body), nil
		}

		fmt.Printf("Error: Status code %d\n", response.StatusCode)

		// Exponential backoff before retrying
		sleepTime := time.Duration(retry) * time.Second
		time.Sleep(sleepTime)
	}

	return result, models.ErrMaxRetriesExceeded
}
