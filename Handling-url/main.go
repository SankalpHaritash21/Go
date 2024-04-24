package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Example data structure to respond to requests or receive data
type Message struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/api/message", messageHandler)

	// Set the server to listen on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handles requests to the /api/message route
func messageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetMessage(w, r)
	case "POST":
		handlePostMessage(w, r)
	default:
		// Handle other HTTP methods
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
	}
}

// Handle GET requests to /api/message
func handleGetMessage(w http.ResponseWriter, r *http.Request) {
	// Create an example message to send as a response
	responseMessage := Message{Text: "Hello from the server!"}
	responseData, err := json.Marshal(responseMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

// Handle POST requests to /api/message
func handlePostMessage(w http.ResponseWriter, r *http.Request) {
	// Read the body of the request
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var message Message
	if err := json.Unmarshal(requestBody, &message); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Log the received message to the console
	fmt.Printf("Received message: %s\n", message.Text)

	// Send a response back to the client
	responseMessage := Message{Text: "Message received!"}
	responseData, err := json.Marshal(responseMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
