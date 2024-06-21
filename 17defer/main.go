package main

import (
	"fmt"
)

// defer work on LIFO order  if we use more than one defer
func main() {
	defer fmt.Println("this will be printed first")
	defer fmt.Println("this will be printed second")

	fmt.Println("Hello from defer in go lang")

	myDefered()
}

func myDefered() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
