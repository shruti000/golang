package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang-  LearnCodeOnline.in")
	//channels are basically used to communicate between the go routines
	// 	//here 2 is the bfferd channel
	// 	//buffered channel eman even if we are passing e values and printing 1 it wn'nt give error
	// because of buffered channels
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	// R ONLY
	//anainymous fucntion which has no name identifier
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)
		//reding channel
		//fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		//writing channel
		myCh <- 4
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)
	fmt.Println(myCh)
	wg.Wait()
}

// 	// fmt.Println(<-myCh)   //here remmebr chanhel going pout of the box
// 	// myCh <- 5              //here channel going in the box
// 	wg.Add(2)
// 	// R ONLY
// 	go func(ch <-chan int, wg *sync.WaitGroup) {

// 		val, isChnelOpen := <-myCh

// 		fmt.Println(isChanelOpen)
// 		fmt.Println(val)

// 		//fmt.Println(<-myCh)
