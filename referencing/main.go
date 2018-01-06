package main

import "fmt"

func main() {

	var a = 30

	var b = &a

	fmt.Println(b, a)   // 0xc04204c058 30
	fmt.Println(*b, &a) // 30 0xc04204c058

	*b = 35

	fmt.Println(a)  // 35
	fmt.Println(*b) // 35

}
