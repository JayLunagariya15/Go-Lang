package main

import (
	// "crypto/rand"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Maths of GO lang")

	// var num1 int = 10
	// var num2 float64 = 5.5
	// fmt.Println("Total is", float64(num1)+num2)

	// Crypto Random Method
	// randomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(randomNumber)

	//Math Random Method
	fmt.Println(rand.Intn(5))
}
