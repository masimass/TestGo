package main

import "fmt"

const (
	_ = iota //this is an iota test
	// Y = 1*5
	Y = iota * 5
	// Z = 2*10
	Z = iota * 10 //Z = 2*10
)

func main() {

	fmt.Println("test\tthis")
	fmt.Printf("%b\t", Y)
	fmt.Printf("%d\n", Y)
	fmt.Printf("%b\t", Z)
	fmt.Printf("%d\n", Z)
}
