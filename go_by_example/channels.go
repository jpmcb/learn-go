package main

import "fmt"

func buff() {
	// channel with a buffer of 2
	messages := make(chan string, 2)

	// Won't block even though there is no ready receive
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func main() {
	// Creates a new string typed channel
	messages := make(chan string)

	// Send a string into the channel in this concurrent go routine
	go func() { messages <- "ping" }()

	// sending and receiving are blocking
	msg := <-messages
	fmt.Println(msg)

	buff()
}
