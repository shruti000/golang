package main

import "fmt"

func server1(ch chan string) {
	ch <- "From server 1"
}

func server2(ch chan string) {
	ch <- "From server 2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case message := <-output1:
		fmt.Println(message)
	case message := <-output2:
		fmt.Println(message)

	}
}
