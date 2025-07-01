package main

import "fmt"

func greetPeople() {
	fmt.Println("Hello World")
}

func greetByName(name string) {
	fmt.Println("Hello", name)
}

func checkUser(name string) string {
	if name == "Jules" {
		return "Admin"
	}

	return name
}
