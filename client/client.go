package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create an HTTP client
	client := &http.Client{}

	address := "http://localhost:8080/sse/testid"

	log.Println("SSE client connecting on address:", address)

	// Create a GET request to the SSE API endpoint
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Create a new scanner to read the response body
	scanner := bufio.NewScanner(resp.Body)

	// Read events from the response body
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Empty line indicates the end of an event
			fmt.Println()
		} else {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading response: %v\n", err)
	}
}
