package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Json in golang")
	// EncodeJson()
	DecodeJson()
}

type course struct {
	Name      string `json:"coursename`
	Price     int
	Plateform string
	Password  string   `json:"-"`
	Tags      []string `json:"tags, omitempty"`
}

func EncodeJson() {
	myCourse := []course{
		{"ReactJS", 299, "Youtube", "123456", []string{"web-dev", "js"}},
		{"Angular", 299, "Youtube", "123456", []string{"web-dev", "js"}},
		{"VueJS", 299, "Youtube", "123456", []string{"web-dev", "js"}},
		{"EmberJS", 299, "Youtube", "123456", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(myCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS",
		"Price": 299,
		"Plateform": "Youtube",
		"Password": "123456",
		"tags": ["web-dev", "js"]	
	}`)

	var myCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON was not valid")
	}
}
