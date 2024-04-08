package main

import (
	"fmt"
	"strings"
)

// isPalindrome checks if a string is a palindrome.
// It ignores case and spaces for simplicity.
func isPalindrome(s string) bool {
    // Convert to lowercase to make the comparison case-insensitive
    s = strings.ToLower(s)
    // Remove spaces to ignore them in the comparison
    s = strings.Replace(s, " ", "", -1)

    // Compare characters from the start and the end moving towards the center
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}

func main() {
    examples := []string{"Madam", "racecar", "hello", "A man a plan a canal Panama","Sankalp","Moon"}
    for _, example := range examples {
        fmt.Printf("Is '%s' a palindrome? %t\n", example, isPalindrome(example))
    }
}
