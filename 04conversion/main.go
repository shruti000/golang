package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Pizza Ratings")
	//all the input from the console is a string
	input, _ := reader.ReadString('\n')
	fmt.Println("The rating of pizza is ", input)
	//here we are usig trim because hwile entering aslong with user input we get \n
	numrating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("new raitng is ", numrating+1)

}
