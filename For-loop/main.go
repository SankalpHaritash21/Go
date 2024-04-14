package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    eventQueue := make([]string, 0)
    stopAfter := 10

    // Simulating adding events to the queue
    for i := 0; i < stopAfter; i++ {
        if rand.Intn(10) > 5 {
            eventQueue = append(eventQueue, "Event: User Clicked")
        } else {
            eventQueue = append(eventQueue, "Event: System Update")
        }
    }
    eventQueue = append(eventQueue, "Event: Stop")

    // Basic event loop
    for len(eventQueue) > 0 {
        // Retrieve the first event
        event := eventQueue[0]
        eventQueue = eventQueue[1:]

        // Process the event
        fmt.Println("Processing:", event)

        // Stop condition
        if event == "Event: Stop" {
            fmt.Println("Stopping the event loop.")
            break
        }

        // Simulate processing time
        time.Sleep(time.Second)
    }
}
