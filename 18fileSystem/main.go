package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("What is file system in go lang")
	content := "This needs go to in file - US Visa Interview"

	file, err := os.Create("./newTextFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	fmt.Println("New file is created", length)

	readFile("./newTextFile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("data in the file is : ", string(databyte))
}

// for checking error

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
