package main

import "fmt"

type person struct {
	age  int
	name string
}

func structs() {

	person := person{age: 20, name: "Jules"}

	person.age = 40

	fmt.Println(person.age)
}
