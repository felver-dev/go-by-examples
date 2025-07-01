package main

import "fmt"

func arrays() {
	var arr [5]int

	fmt.Println(arr)

	// Affecting values
	arr[2] = 100
	arr[0] = 5

	fmt.Println(arr)

	// Check the length of the array
	fmt.Println(len(arr))

	// Declare and initialize an array
	arr2 := [4]int{1, 2, 3, 4}
	var arr3 = [4]int{1, 2, 3, 4}

	fmt.Println(arr2)
	fmt.Println(arr3)

	// If you don't know the size of
	// the array then you can the compiler count for you
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr4)

	// Multi dimension arrays
	twoDimArr := [][]int{
		{10, 40},
		{12, 18},
	}

	fmt.Println(twoDimArr)

	// Reading array elements
	for j := 0; j < len(arr4); j++ {
		fmt.Println("Elt: ", arr4[j])
	}

	for k := range arr {
		fmt.Println("Arr :", arr[k])
	}

	fmt.Println("Reading two dim array")
	for m := range 2 {
		for n := range 2 {
			twoDimArr[m][n] = m + n
			fmt.Println("Second dim", n)
		}
		fmt.Println("First dim", m)

	}

	// fmt.Println("2d:", twoDimArr)

}
