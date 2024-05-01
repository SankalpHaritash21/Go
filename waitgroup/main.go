package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Notify the WaitGroup that this worker is done

    // Simulate some work by sleeping
    sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
    fmt.Printf("Worker %d is sleeping for %v\n", id, sleepDuration)
    time.Sleep(sleepDuration)
    fmt.Printf("Worker %d is done\n", id)
}

func main() {
    var wg sync.WaitGroup
    numWorkers := 5

    // Seed random number generator (not necessary for security purposes)
    rand.Seed(time.Now().UnixNano())

    // Add the number of goroutines to wait for
    wg.Add(numWorkers)

    // Launch several goroutines and pass the WaitGroup pointer
    for i := 1; i <= numWorkers; i++ {
        go worker(i, &wg)
    }

    // Wait for all goroutines to finish
    wg.Wait()
    fmt.Println("All workers completed")
}
