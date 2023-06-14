package main

import "fmt"

func captian(in chan string, s string) {
	in <- s
}

func main() {

	ninja1, ninja2 := make(chan string), make(chan string)
	go captian(ninja1, "NINJA 1")
	go captian(ninja2, "NINJA 2")
	//fmt.Println(<-ninja1)
	//fmt.Println(<-ninja2)

	select {
	//here we are readin the ninja one and putitng it not message new parameter
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	}

}
