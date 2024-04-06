package main

import (
	"fmt"
)

func Print() {
	// Using fmt.Println
	name := "Jane"
	age := 28
	fmt.Println("Using fmt.Println:")
	fmt.Println("Name:", name, "Age:", age)

	// Adding a separator line for clarity
	fmt.Println("--------------------------")

	// Using fmt.Printf
	fmt.Println("Using fmt.Printf:")
	fmt.Printf("Name: %s Age: %d\n", name, age)
}
