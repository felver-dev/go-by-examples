package main

import "fmt"

type Rect struct {
	width  int
	height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r Rect) perimeter() int {
	return 2 * (r.width + r.height)
}

func methods() {
	rect := Rect{width: 10, height: 5}

	fmt.Println(rect.area())
	fmt.Println(rect.perimeter())

	rp := &rect
	fmt.Println("Area: ", rp.area())
	fmt.Println("Perimeter: ", rp.perimeter())
}
