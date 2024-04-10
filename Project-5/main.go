package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket represents a token bucket used for rate limiting.
type TokenBucket struct {
	rate       float64     // Tokens per second
	capacity   float64     // Maximum tokens the bucket can hold
	tokens     float64     // Current number of tokens in the bucket
	lastUpdate time.Time   // Last time the bucket was updated
	mu         sync.Mutex  // Mutex for thread safety
}

// NewTokenBucket creates a new TokenBucket with the specified rate and capacity.
func NewTokenBucket(rate, capacity float64) *TokenBucket {
	return &TokenBucket{
		rate:     rate,
		capacity: capacity,
		tokens:   capacity,
	}
}

// Take takes a specified number of tokens from the bucket if available.
func (tb *TokenBucket) Take(tokens float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	tb.tokens += tb.rate * now.Sub(tb.lastUpdate).Seconds()
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastUpdate = now

	if tb.tokens >= tokens {
		tb.tokens -= tokens
		return true // Enough tokens available, request can be processed
	}
	return false // Not enough tokens available, request should be delayed
}

func main() {
	// Create a new token bucket with rate 2 tokens per second and capacity 5 tokens.
	bucket := NewTokenBucket(2, 5)

	// Simulate incoming requests.
	for i := 0; i < 10; i++ {
		if bucket.Take(1) {
			fmt.Println("Processing request", i+1)
		} else {
			fmt.Println("Rate limit exceeded, request delayed")
		}
		time.Sleep(time.Second) // Simulate delay between requests
	}
}
