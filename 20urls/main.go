package main

import (
	"fmt"
	"net/url"
)

const myurls string = "http://localhost:5500/19webRequest/test.html?coursename=react&payment=sa645c1d65sv"

func main() {
	fmt.Println("Handeling urls in go ")
	fmt.Println("urls are: ", myurls)

	//parsing ( extract data )
	result, _ := url.Parse(myurls)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("params is: ", val)
	}

	partsofUrls := &url.URL{
		Scheme:   "http",
		Host:     "localhost:5500",
		Path:     "/19webRequest/test.html",
		RawQuery: "coursename=react&payment=sa645c1d65sv",
	}

	anotherUrl := partsofUrls.String()
	fmt.Println(anotherUrl)
}
