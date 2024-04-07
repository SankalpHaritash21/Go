package main

import (
	"fmt"
)

func fibonacci(limit int) []int {
	sequence := make([]int, 0)
	a, b := 0, 1
	for b < limit {
		sequence = append(sequence, b)
		a, b = b, a+b
	}
	return sequence
}

func main() {
	var limit int
	fmt.Print("Enter the limit for the Fibonacci sequence: ")
	_, err := fmt.Scanf("%d", &limit)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Generating Fibonacci sequence up to %d\n", limit)
	fibSequence := fibonacci(limit)
	fmt.Println("Fibonacci sequence:", fibSequence)
}
