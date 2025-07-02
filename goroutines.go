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

func direBonjour(nom string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Bonjour %s! (%d)\n", nom, i)
		time.Sleep(time.Second)
	}
}
func Goroutines() {
	// fg("direct")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")
	// go fg("goroutine")

	// time.Sleep(time.Second)
	// fmt.Println("done")

	go direBonjour("Alice")
	go direBonjour("Bob")

	time.Sleep(4 * time.Second)
	fmt.Println("Terminé")
}
