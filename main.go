package main

import (
	"fmt"
	"net/http"
	"time"
)

// Simulate database fetching
func fetchDataFromDatabase(requestID int) string {
	// Simulating a database fetch with a sleep
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Data for request #%d", requestID)
}

// Simulate data processing
func processData(data string) string {
	// Simulate some processing
	return fmt.Sprintf("Processed: %s", data)
}

// Handle each request
func handleRequest(requestID int, dataChannel chan string) {
	rawData := fetchDataFromDatabase(requestID)
	processedData := processData(rawData)
	dataChannel <- processedData
}

func main() {
	requestID := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestID++
		dataChannel := make(chan string)

		// Start the goroutine for handling this request
		go handleRequest(requestID, dataChannel)

		// Get the processed data from the channel
		processedData := <-dataChannel
		fmt.Fprintf(w, "Response: %s", processedData)
	})

	fmt.Println("Server is starting...")
	http.ListenAndServe(":8080", nil)
}
