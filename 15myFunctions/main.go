package main

import "fmt"

func main() {
	greeter()
	fmt.Println("functions in go lang")

	result := adder(2, 5)
	fmt.Println("Result is: ", result)

	proResult, message := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println("Pro extra result is: ", message)
}

func adder(x int, y int) int {
	return x + y
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, anyValue := range values {
		total += anyValue
	}

	return total, "Extra Line"
}

func greeter() {
	fmt.Println("Namstey from India")
}
