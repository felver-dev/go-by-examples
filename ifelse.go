package main

import "fmt"

func ifelse() {
	var n int = 101
	if n%2 == 0 {
		fmt.Println("This is a odd number")
	} else if n%2 != 0 {
		fmt.Println("This is an even number")
	}

	name := "JulesF"
	if name == "Jules" {
		fmt.Println("Go man !!!")
	} else {
		fmt.Println("Oops man, gotta stop here ")
	}

}
