package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
			burstyLimiter <- t
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 20)
	for i := 1; i <= 20; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("burstyRequests", req, time.Now())
	}
}
