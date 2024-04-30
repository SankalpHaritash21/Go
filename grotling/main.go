package main

import (
	"fmt"
)

// Function to sum numbers in a slice and send the result to a channel
func sum(numbers []int, resultChan chan int) {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    resultChan <- sum // send sum to channel
}

func main() {
    // Slices of numbers
    slice1 := []int{1, 2, 3, 4, 5}
    slice2 := []int{6, 7, 8, 9, 10}

    // Channels to receive the results
    resultChan1 := make(chan int)
    resultChan2 := make(chan int)

    // Start goroutines to sum each slice
    go sum(slice1, resultChan1)
    go sum(slice2, resultChan2)

    // Receive results from channels
    sum1 := <-resultChan1
    sum2 := <-resultChan2

    // Calculate total sum
    total := sum1 + sum2

    // Print the results
    fmt.Printf("Sum of %v is %d\n", slice1, sum1)
    fmt.Printf("Sum of %v is %d\n", slice2, sum2)
    fmt.Printf("Total sum is %d\n", total)
}
