package main

import "fmt"

type User struct {
	name   string
	email  string
	status bool
	age    int
}

func main() {
	fmt.Println("structs in golang")

	//no inhertitance in golang;no super or parent
	u := User{
		"shruti",
		"s@gmail.com",
		true,
		14,
	}
	fmt.Println(u)
	//wih name
	fmt.Printf("%+v\n", u)
	fmt.Println(u.name)

}
