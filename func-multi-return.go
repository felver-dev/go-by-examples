package main

func multiple_return(name string) (bool, int) {
	isAdmin := false
	age := 0
	if name == "jules" {
		isAdmin = true
		age = 28
		return isAdmin, age
	}
	return isAdmin, age
}
