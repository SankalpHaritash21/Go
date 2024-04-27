package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define a struct that matches the expected JSON payload structure
type Payload struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Country string `json:"country"`
}

func main() {
    // Create a new instance of Payload
    data := Payload{
        Name:    "John Doe",
        Age:     30,
        Country: "USA",
    }

    // Marshal the data into a JSON byte slice
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println(err)
        return
    }

    // The URL of the API endpoint
    url := "http://example.com/api/data"

    // Create a new HTTP POST request
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Set the Content-Type header to indicate JSON data
    req.Header.Set("Content-Type", "application/json")

    // Create a new HTTP client and execute the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    // Read and print the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Response Status: %s\n", resp.Status)
    fmt.Printf("Response Body: %s\n", string(body))
}

