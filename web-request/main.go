package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Specify the URL
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Send a GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer response.Body.Close() // Don't forget to close the response body

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response body
	fmt.Println("Response Body:", string(body))
}
