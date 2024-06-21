package main

import "fmt"

func main() {
	fmt.Println("Struct in Go Lang")
	//no inheritance in go lang; no super or parent or child
	Jay := User{"Jay Lunagariya", "jay@gamil.com", true, 25}
	fmt.Println(Jay)

	fmt.Printf("All Details are : %+v\n", Jay)

	fmt.Printf("Name is %v and Email is %v\n", Jay.Name, Jay.Email)

	Jay.GetStatus()

	Jay.NewMail()

	fmt.Println("Email of this user is: ", Jay.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
