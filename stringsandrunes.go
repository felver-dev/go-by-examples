package main

import "fmt"

func stringsandrunes() {
	const s = "abcdefghijklmnopqrstuvwxyz"

	fmt.Println("Len: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	for _, val := range s {
		fmt.Printf("Character %d: %c\n", val)
	}
}
