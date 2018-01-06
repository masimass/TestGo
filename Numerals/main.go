package main

import "fmt"

func main() {
	for i := 0; i < 16; i++ {

		fmt.Printf("%b \t %d \t %X \n", i, i, i)
	}

}
