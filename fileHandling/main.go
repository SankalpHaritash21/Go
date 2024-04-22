package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Step 1: Create and write to a file
	fileName := "example.txt"
	content := []byte("Hello, i Sam I am learning about Golang! This demonstrates file handling.\n")

	// Creating and writing to the file
	err := ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println("Error creating and writing to file:", err)
		return
	}
	fmt.Println("File created and data written.")

	// Step 2: Read from the file
	readContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Read from file: %s", string(readContent))

	// Step 3: Delete the file
	err = os.Remove(fileName)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully.")
}
