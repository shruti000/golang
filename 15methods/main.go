package main

import "fmt"

type User struct {
	name   string
	email  string
	status bool
	age    int
}

//method
//basicllay think like class sruct has varibale and ethods
//methoda r like fucntion only but they have a reciever
func (us User) getStatus() {
	fmt.Println("The status of user is ", us.status)
}

func (us User) updateEmail() {
	us.email = "p@gmail.com"
	fmt.Println(us)
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
	fmt.Println("***********method Testing **************")
	u.getStatus()
	u.updateEmail()
	//here only the copy is passes because the email is not chnanged
	fmt.Println("***original user is ")
	fmt.Println(u)

}
