package main

import "fmt"

func main() {

	loginCount := 9
	var result string

	if loginCount < 10 {
		result = "regualr user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "excatly 10 count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("odd")
	}

	//declae and condition at same time
	if num := 3; num < 10 {
		fmt.Println("number less than 10")
	} else {
		fmt.Println("number greater than 10")
	}

}
