package main

import "fmt"
import "time"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend yo")
	default:
		fmt.Println("Not the weekened ... ")
	}

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("It's before noon!")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Don't know %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
