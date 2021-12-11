package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello world2!")
// }

import (
	"fmt"
	"mp/glhf/ch01"
	"mp/glhf/ch02"
)

func main() {
	duh()
	ch01.Hello1()
	//ch01.A()
	//ch01.B()
	//ch01.C()
	//ch01.D()
	//ch01.E()
	//ch01.F()
	ch02.B()
}

func duh() {
	fmt.Print("Hello world!")
	fmt.Println("Hello world!")
	fmt.Println('H')
}
