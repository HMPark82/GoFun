package ch01

import "fmt"

func A() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))
}
