package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the rating off the pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating  ", input)
}
