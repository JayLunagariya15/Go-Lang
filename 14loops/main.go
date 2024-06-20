package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops in go lang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("days of week: ", days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days { // at the place of _ we can use index short hand for index is i
	// 	fmt.Printf("Index is v and value is: %v \n", day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto label
		}

		if rougueValue == 5 {
			rougueValue++
			continue // to skip the current iteration if we use continue
			// if we use break then it will terminate the loop
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

label:
	fmt.Println("We have used goto label")
}
