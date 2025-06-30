package main

import "fmt"

func variables() {
	var name string = "jules"

	var b, c int = 1, 3

	var boolean bool = true

	var typedIntVar int
	var typedStringVar string
	var typedFloatVar float64
	var typedBoolVar bool
	inferredIntVar := 42

	fmt.Println("Name:", name)

	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Println("Boolean:", boolean)

	fmt.Println("Typed Int Variable:", typedIntVar)
	fmt.Println("Typed String Variable:", typedStringVar)

	fmt.Println("Typed Float Variable:", typedFloatVar)
	fmt.Println("Typed Bool Variable:", typedBoolVar)

	fmt.Println("Inferred Int Variable:", inferredIntVar)
}
