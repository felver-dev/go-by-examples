package main

import "fmt"

func maps() {

	students := make(map[string]string)

	students["jules"] = "Jules"
	fmt.Println(students["jules"])

}
