package main

import (
	"fmt"
)

// & means Reference and * means value

func main() {
	fmt.Println("points to pointer section in go lang")

	// var ptr *int
	// fmt.Println("ptr value is ", ptr)

	myNumber := 20
	var ptr = &myNumber
	fmt.Println("ptr value is ", ptr)  // this will return memory address of pointer variable because ptr = &myNumber which is reference
	fmt.Println("ptr value is ", *ptr) // this will return value of pointer variable because *ptr = myNumber which is value

	*ptr = *ptr * 2
	fmt.Println("new value is: ", myNumber)
	fmt.Println("Pointer new value is: ", *ptr)
}
