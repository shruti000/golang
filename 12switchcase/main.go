package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("switch case in golang")
	dicenumber := rand.Intn(6) + 1
	fmt.Println(dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("dice value is one")
		fallthrough
	case 2:
		fmt.Println("dice value is two")
	case 3:
		fmt.Println("dice value is three")
	case 4:
		fmt.Println("dice value is four")
	case 5:
		fmt.Println("dice value is five")
		//falthrough basically like a continue
		fallthrough
	case 6:
		fmt.Println("dice value is six")
	default:
		fmt.Println("invalid")
	}

}
