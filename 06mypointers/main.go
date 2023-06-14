package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	//& means refernce
	mynumber := 23
	var ptr = &mynumber
	fmt.Println(ptr, *ptr)

	*ptr = *ptr * 2
	fmt.Println("value after changed is ", mynumber, ptr, *ptr)
}
