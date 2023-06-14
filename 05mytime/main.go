package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to mytime study")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("09-02-2021   15:04:23 Monday"))
	createDate := time.Date(2020, time.November, 23, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)

}
