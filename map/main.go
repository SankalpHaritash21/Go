package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define the input text
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	// Split the text into words
	words := strings.Fields(text)

	// Create a map to store word frequencies
	wordFreq := make(map[string]int)

	// Iterate over the words and count their frequencies
	for _, word := range words {
		// Convert word to lowercase to ensure case-insensitive counting
		word = strings.ToLower(word)
		wordFreq[word]++
	}

	// Print the word frequencies
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
