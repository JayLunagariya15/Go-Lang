package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices are used more in Go Lang")

	var fruitsList = []string{"Apple", "Peach", "Pineapple"}
	fmt.Printf("Type of fruitsList is: %T\n ", fruitsList)

	fruitsList = append(fruitsList, "Mango", "Banana", "Orange")
	fmt.Println(fruitsList)

	fruitsList = append(fruitsList[:4])
	fmt.Println(fruitsList)

	//make() and New()  for memory allocation

	highscores := make([]int, 4)

	highscores[0] = 978
	highscores[1] = 254
	highscores[2] = 627
	highscores[3] = 39
	fmt.Println(highscores)

	highscores = append(highscores, 555, 666, 777)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	var courses = []string{"ReactJS", "Angular", "VueJS", "EmberJS", "Polymer", "Python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
