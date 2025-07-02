package main

import (
	"fmt"
	"time"
)

func fg(from string) {
	for i := range 3 {
		fmt.Println(from, ": ", i)
	}
}

func Goroutines() {
	fg("direct")

	go fg("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
