package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("switch case in go lang")
	// rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can move 1 cell")
	case 2:
		fmt.Println("Dice value is 2 and you can move 2 cell")
	case 3:
		fmt.Println("Dice value is 3 and you can move 3 cell")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 and you can move 4 cell")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 and you can move 5 cell")
	case 6:
		fmt.Println("Dice value is 6 and you can move 6 cell and roll the dice again")
	default:
		fmt.Println("What was that !!")
	}
}
