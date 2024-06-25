package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:5500/19webRequest/test.html"

func main() {
	fmt.Println("In this section we are not going to see API realted web request")

	response, err := http.Get(url)
	checkNilError(err)
	fmt.Printf("Type of response is: %T\n", response)

	defer response.Body.Close() // callers responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	content := string(databytes)
	fmt.Println("the data is: ", content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
