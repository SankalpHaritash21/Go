package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Define a sample string
	sampleText := "Hello, World! 🌍"

	// Strings in Go are UTF-8 encoded and can be ranged over to handle each Unicode code point
	// This prints each Unicode code point of the string
	for i, runeValue := range sampleText {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
	}

	// Splitting a string into parts
	parts := strings.Split(sampleText, ", ")
	fmt.Println("Parts:", parts)

	// Converting string to uppercase
	upper := strings.ToUpper(sampleText)
	fmt.Println("Uppercase:", upper)

	// Reversing a string by treating it as a slice of runes
	runes := []rune(sampleText)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println("Reversed String:", string(runes))

	// Finding and counting substrings
	count := strings.Count(sampleText, "l")
	fmt.Println("Count of 'l':", count)

	// Trim function using custom cutset
	trimmed := strings.Trim("!Hello, World! !!!", " !")
	fmt.Println("Trimmed:", trimmed)

	// Replace all occurrences of a substring with another substring
	replaced := strings.ReplaceAll(sampleText, "World", "Go")
	fmt.Println("Replaced:", replaced)

	// Advanced: Using FieldsFunc to split a string based on a function
	advancedSplit := strings.FieldsFunc(sampleText, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	fmt.Println("Advanced Split:", advancedSplit)
}

