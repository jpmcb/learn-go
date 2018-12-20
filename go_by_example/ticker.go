package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		// Receive from the ticker channel
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
