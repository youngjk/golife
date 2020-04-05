package main

import (
	"fmt"
	"time"
)

func main() {
	// Requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Limited Channel will receive value every 200 milliseconds
	limiter := time.Tick(200 * time.Millisecond)

	// Blocking on receive from limiter channel before serving each request => Limit request per every 200 milliseconds
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Allow short bursts of requests while preserving overall rate limit
	burstyLimiter := make(chan time.Time, 3)

	// Fill up channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 millisecond we'll try to add new value to burstyLimiter
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Simulate 5 more incoming requests. First 3 will benefit from burstyLimiter
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Request", req, time.Now())
	}
}