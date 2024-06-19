package main

import "fmt"

func main() {
	fmt.Println("Arrays are less Use in Go Lang")

	var fruitsList [4]string
	fruitsList[0] = "Apple"
	fruitsList[1] = "Tomato"
	fruitsList[3] = "Banana"

	fmt.Println("List of fruits: ", fruitsList)
	fmt.Println("Total number of elements: ", len(fruitsList))

	var vegList = [3]string{"Potato", "Beans", "Carrot"}
	fmt.Println("List of vegitables: ", vegList)
	fmt.Println("Total number of elements: ", len(vegList))
}
