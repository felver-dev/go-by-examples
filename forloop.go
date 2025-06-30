package main

import (
	"fmt"
	"time"
)

func forLoop() {

	// inferring the i var for the loop needs
	i := 1
	for i <= 3 {
		fmt.Println("Value : ", i)
		if i == 3 {
			fmt.Println("Reached the end ", time.Now().Second(), " seconds")
			break
		}
		i++
	}

	for j := 1; j <= 3; j++ {
		fmt.Println("Value : ", j)
	}

	// The range loop
	for i := range 3 {
		fmt.Println(i, "time")
	}

	for {
		fmt.Println("Infinite loop")
		time.Sleep(1 * time.Second)
		if time.Now().Second()%5 == 0 {
			fmt.Println("Breaking infinite loop at ", time.Now().Second(), " seconds")
			break
		}
		fmt.Println("Still running at ", time.Now().Second(), " seconds")
	}
}
