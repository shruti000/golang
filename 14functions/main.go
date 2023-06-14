package main

import "fmt"

func greeter() {
	fmt.Println("This is a greeter fucntion")
}
func adder(a, b int) int {
	return a + b
}

//IMPORTANT
//important to add multiple values
//if we return two value we hav to change the declaration
func proadder(values ...int) (int, string) {
	toltal := 0
	//here values will be a slice of values so loopimg through for loop
	for _, v := range values {
		toltal += v
	}
	return toltal, "shruti"
}

func main() {
	fmt.Println("welcome to fucntions in golang")
	greeter()
	m := 2
	n := 5
	result := adder(m, n)
	fmt.Println(result)
	fmt.Println("***PROADDER****")
	t1, message := proadder(1, 2, 3, 4, 5)
	fmt.Println(t1, message)
}
