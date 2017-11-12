package main

import "fmt"

var b int

func main() {

	a := 2

	fmt.Println(a)
	foo()
}

func foo() {
	fmt.Println(b)
}
