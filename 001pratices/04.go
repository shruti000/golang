package main

import "fmt"

func students(name ...string) {
	for _, value := range name {
		fmt.Println(value)
	}

}

func main() {
	s := "hello world"
	fmt.Println(s[0:4])
	fmt.Println(s[6:])
	students("shruti", "pranav", "chinna", "ammu", "junnu")

	//vardiaci fucntion

}
