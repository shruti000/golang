package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string = "marry"
	fmt.Println(s)

	//program to take input from console
	name := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name ")
	value, _ := name.ReadString('\n')
	fmt.Println(value)

}
