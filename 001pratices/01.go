// take a number from cosole and check wether it is between 1 -10
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	num := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number :")
	value, _ := num.ReadString('\n')
	num1, _ := strconv.Atoi(value)
	if num1 >= 1 && num1 < 10 {
		fmt.Println("number is between 1-10")
	} else {
		fmt.Println("NUmber is grater than 10")
	}

}
