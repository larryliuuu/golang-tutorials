package main

import (
	"fmt"
	"time"
)

func burstyLimiter() {
	burstyLimiter := make(chan time.Time, 3)

	// initialize to allow for 3 requests to be received in the first 200ms
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// allow for 3 requests to be sent every 200ms
			for i := 0; i < 3; i++ {
				burstyLimiter <- t
			}
		}
	}()

	burstyRequests := make(chan int, 50)
	for i := 1; i <= 50; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func main() {
	burstyLimiter()
}
