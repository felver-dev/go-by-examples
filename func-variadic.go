package main

import "fmt"

func variadic_func(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}
