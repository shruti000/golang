package main

import "fmt"

//bascially used to cnenct the channel where output ofone channel is given to input of other channe
//this fucntion takes number and store it in the channel named ot.It is the frst channel
//here ... is the varidic fucntion means which can recieve ay nuber of argument an dwe can lopp jsing slice
func generate(nums ...int) <-chan int {
	out := make(chan int, 4)
	go func() {
		for _, value := range nums {
			out <- value
		}
		close(out)
	}()
	return out
}

//this fuction takes values from the above channel as input and sqaure it and produce it pwn oupptu
func square(val <-chan int) <-chan int {
	out := make(chan int, 4)
	go func() {
		for n := range val {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	c := generate(2, 3, 5, 6)
	out := square(c)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}
