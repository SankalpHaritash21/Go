// Declare the main package. In Go, every executable program starts with the main package.
package main

// Import necessary packages. Here, we're importing "errors" for creating error messages and "fmt" for printing and reading input.
import (
	"errors"
	"fmt"
)

// divide is a function that takes two float64 numbers as input and returns a float64 and an error.
// It tries to divide the first number by the second and handles division by zero as an error.
func divide(a, b float64) (float64, error) {
	// Check if the divisor is zero. Division by zero is not allowed and returns an error.
	if b == 0 {
		// If b is 0, we return 0 as the result and an error message saying "cannot divide by zero".
		return 0, errors.New("cannot divide by zero")
	}
	// If b is not 0, perform the division, return the result, and nil for the error (indicating no error occurred).
	return a / b, nil
}

// main is the entry point of the program.
func main() {
	var a, b float64

	// Ask the user for the first number.
	fmt.Print("Enter the first number: ")
	// Read the first number from the user input.
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Error reading the first number:", err)
		return
	}

	// Ask the user for the second number.
	fmt.Print("Enter the second number: ")
	// Read the second number from the user input.
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Error reading the second number:", err)
		return
	}

	// Attempt to divide the two numbers.
	result, err := divide(a, b)
	// Check if an error occurred.
	if err != nil {
		// If err is not nil, print the error.
		fmt.Println("Error:", err)
	} else {
		// If there is no error, print the result of the division.
		fmt.Println("Result:", result)
	}
}
