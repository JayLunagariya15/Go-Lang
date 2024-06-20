package main

import "fmt"

func main() {
	fmt.Println("Struct in Go Lang")
	//no inheritance in go lang; no super or parent or child
	Jay := User{"Jay Lunagariya", "jay@gamil.com", true, 25}
	fmt.Println(Jay)

	fmt.Printf("All Details are : %+v\n", Jay)

	fmt.Printf("Name is %v and Email is %v\n", Jay.Name, Jay.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
