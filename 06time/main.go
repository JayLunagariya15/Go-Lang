package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time to learn time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// standard formate have to use all the time for perfect date and time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.January, 25, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
