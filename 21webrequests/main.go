package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Handeling web requests in go ")

	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:7007/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("content length is: ", response.ContentLength)

	var responseString strings.Builder // used to build string which is library
	content, _ := io.ReadAll(response.Body)
	byteCode, _ := responseString.Write(content)

	fmt.Println("Byte code is: ", byteCode)
	fmt.Println(responseString.String())
	// fmt.Println("content is: ", string(content))
}

func PerformPostRequest() {
	const myUrl = "http://localhost:7007/post"

	// fake payload data
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("response is: ", string(content))
}

func PerformPostJsonRequest() {
	const myurls = "http://localhost:7007/postform"

	data := url.Values{}
	data.Add("firstname", "jay")
	data.Add("lastname", "lunagariya")
	data.Add("email", "jay@gmail.com")

	response, err := http.PostForm(myurls, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("response is: ", string(content))
}
