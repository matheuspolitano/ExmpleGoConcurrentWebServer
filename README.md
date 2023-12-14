## Developer
- **Name:** Matheus Politano
- **LinkedIn:** [Matheus Politano](https://www.linkedin.com/in/matheus-politano-08b762123/)
 
## Overview 

My goal was created a simplified version of the web server scenario using Go involves implementing a basic HTTP server, goroutines for handling database fetching and data processing, and channels for communication between these components. For this example, let's assume the data processing is a simple operation, and we'll simulate database fetching with a dummy function.



# Explanation

## HTTP Server Setup
- `http.HandleFunc` sets up a basic HTTP server that listens to requests on the root path.
- Each incoming request increments `requestID` for unique identification.

## Goroutine for Request Handling
- `handleRequest` function is called as a goroutine for each request.
- It simulates fetching data from a database and then processing that data.

## Channels for Data Transfer
- A channel `dataChannel` is created for each request to transfer data between the goroutine handling the request and the main goroutine that sends the HTTP response.

## Data Fetching and Processing
- `fetchDataFromDatabase` and `processData` are placeholder functions representing database operations and data processing.

## Response to Client
- The server waits for the processed data via the channel and sends it back as an HTTP response.

## Server Execution
- The server listens on port 8080.

# Running the Server

To run this server, save the code in a file (e.g., `main.go`), and then run it using the Go command:

```sh
go run main.go
```

Then, you can visit `http://localhost:8080` in a browser or use a tool like `curl` to make requests to the server. Each request will be processed concurrently, demonstrating the use of goroutines and channels in a web server context.

