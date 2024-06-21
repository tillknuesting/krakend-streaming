package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var messageChannel = make(chan string, 100)

func main() {
	go generateMessages()
	http.HandleFunc("/events-stream/", handleEvents)
	address := ":9081"
	fmt.Printf("Server Listening on %s\n", address)
	http.ListenAndServe(address, nil)
}

func generateMessages() {
	for i := 1; i <= 100; i++ {
		message := fmt.Sprintf("data: Hello, SSE! (Message %d)\n\n", i)
		messageChannel <- message
		sleepDuration := time.Duration(rand.Intn(250)+950) * time.Millisecond
		time.Sleep(sleepDuration)
	}
	messageChannel <- "event: done\ndata: Server is done sending events.\n\n"
	close(messageChannel)
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Extract UUID from URL
	uuid := strings.TrimPrefix(r.URL.Path, "/events-stream/")
	fmt.Printf("Received UUID: %s\n", uuid)

	for {
		select {
		case message, ok := <-messageChannel:
			if !ok {
				return
			}
			fmt.Fprintf(w, "event: message\n")
			fmt.Fprintf(w, message)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
