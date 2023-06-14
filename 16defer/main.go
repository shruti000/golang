package main

import "fmt"

func main() {
	//since defer is placed it is read line by line but exceuted at the last just like FILO

	//here defer statck looksl ike hello, world ,mumbai when come sput mumbai,world ,hello
	fmt.Println("first")
	defer fmt.Println("hello")
	defer fmt.Println("world")
	defer fmt.Println("mumbai")
	fmt.Println("last")
	mydefer()
	//output -  first.last,mumbai,world,hello,54321
}

//one,two,three,four,five
//while comming out five ,four ,three ,two,one
func mydefer() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
