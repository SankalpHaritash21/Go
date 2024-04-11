package main

import (
	"fmt"
)

// filterAndDouble takes a slice of integers, filters out even numbers,
// and doubles the odd numbers.
func filterAndDouble(nums []int) []int {
	var result []int // Creating an empty slice
	for _, num := range nums {
		if num%2 != 0 {
			// Doubling the number if it's odd and appending to the result slice
			result = append(result, num*2)
		}
	}
	return result
}

func main() {
	// Initial slice of numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filtering and doubling the numbers
	modifiedNumbers := filterAndDouble(numbers)

	// Printing the original and modified slices
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Modified numbers:", modifiedNumbers)


		// Creating a slice with make
	slice := make([]int, 5) // Creating a slice of int with length 5
	fmt.Println("Slice:", slice)

	// Modifying the slice
	for i := range slice {
		slice[i] = i * 2 // Assigning values to each element of the slice
	}

	fmt.Println("Modified Slice:", slice)
	
}

