package main

func closures() func() int {
	counter := 0

	return func() int {
		counter++
		return counter
	}
}
