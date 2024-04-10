package main

import (
	"context"
	"fmt"

	"golang.org/x/time/rate"
)

func rateLimit() {
	// Create a rate limiter that allows 2 requests per second with a burst limit of 5.
	limiter := rate.NewLimiter(2, 5)

	// Simulate incoming requests.
	for i := 0; i < 10; i++ {
		// Wait for the limiter to allow the next request.
		err := limiter.Wait(context.Background())
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Process the request.
		fmt.Println("Processing request", i+1)
	}
}
