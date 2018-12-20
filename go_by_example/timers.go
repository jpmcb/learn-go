package main

import "time"
import "fmt"

func main() {

	// Indicates how long you want to wait to do something
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 expired!")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired!")
	}()

	// Stop the above go routine before it can execute
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
