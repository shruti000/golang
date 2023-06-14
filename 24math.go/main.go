package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to math in golang")
	//basically it helps to give ranfom number evetime ie number t not repeat
	rand.Seed(time.Now().UnixNano())
	//random number
	//0-4 5 is exclusive
	a := rand.Intn(5)
	fmt.Println(a)

	fmt.Println(rand.Intn(10) + 1)

}
