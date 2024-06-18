package main

import "fmt"

// var jwtToken int = 300000.0
// := is not allowed to use outside the class or method
const LoginToken = "oiugahdvsDYGHFDBFHBbcsabcIHCiAKN85112516AvVF56d4v5d1v56d4165vFDFd51ca" // public because L is capital in LoginToken

func main() {
	var username string = "Jay Lunagariya"
	fmt.Println(username)
	fmt.Printf("Variables is type of string: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is type of string: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is type of string: %T \n", smallVal)

	var smallFloat float32 = 255.84651324653286532
	fmt.Println(smallFloat)
	fmt.Printf("Variables is type of string: %T \n", smallFloat)

	// default values and some aliases
	var anotherValues string
	fmt.Println(anotherValues)
	fmt.Printf("Variables is type of string: %T \n", anotherValues)

	//implicit type
	var website = "google.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variables is type of: %T \n", LoginToken)
}
