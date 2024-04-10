package main

import (
	"fmt"
)

func main() {
    // Declaring and initializing an array
    var arr [5]int
    fmt.Println("Empty array:", arr)

    // Initializing array with values
    arr = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array with values:", arr)

    // Accessing elements
    fmt.Println("First element:", arr[0])
    fmt.Println("Third element:", arr[2])

    // Length of the array
    fmt.Println("Length of the array:", len(arr))

    // Iterating over the array
    fmt.Println("Iterating over the array:")
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // Using range for iteration
    fmt.Println("Iterating over the array using range:")
    for index, value := range arr {
        fmt.Printf("arr[%d] = %d\n", index, value)
    }

    // Copying array
    arrCopy := arr
    fmt.Println("Copied array:", arrCopy)

    // Modifying copied array
    arrCopy[0] = 10
    fmt.Println("Original array:", arr)
    fmt.Println("Modified copied array:", arrCopy)

    // Comparing arrays
    arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    isEqual := true
    for i := range arr1 {
        if arr1[i] != arr2[i] {
            isEqual = false
            break
        }
    }
    fmt.Println("Are arr1 and arr2 equal?", isEqual)
}
