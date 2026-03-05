package main

import (
	"fmt"
	"io"
	"net/http"
	// "time"
)

const url = "http://localhost:3000/"

// Handler function for the server
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the local server!")
}

// Function to start the server
func startServer() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Start the server in a separate goroutine
	go startServer()

	// Give the server a moment to start
	// time.Sleep(1 * time.Second)

	// Client request
	fmt.Println("Welcome to web requests")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // Close the response body

	fmt.Printf("The response is of type %T\n", response)

	// Read and print the response
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
