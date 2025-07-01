package main

import (
	"fmt"
	"time"
)

func Switch() {
	i := 9
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Weird")
	}

	fmt.Println("Are we in the weekend?")
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("Yeahhh, This is weekend")
	default:
		fmt.Println("No, Hold on, you are in the week")
	}

}
