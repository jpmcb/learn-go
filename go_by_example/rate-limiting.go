package main

import (
	"fmt"
	"time"
)

func main() {

	// Fill channel with reqs
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Ticks every 200 mls
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Allow for additional bursts of receiving reqs
	burstLimit := make(chan time.Time, 3)

	// Fill the burst limiter with times
	for i := 0; i < 3; i++ {
		burstLimit <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimit <- t
		}
	}()

	burstReq := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstReq <- i
	}
	close(burstReq)
	for req := range burstReq {
		<-burstLimit
		fmt.Println("request", req, time.Now())
	}

}
