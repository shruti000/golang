package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")

	//writing content to a file
	content := "this content needs to go inside a file with some edition"

	//to create file we use os

	file, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}

	//to write file we use io package
	//it return the lenght of the file
	lenght, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("sucessfull ", lenght)
	readFile("sample.txt")
	file.Close()

}

// method to read the file
func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	//it reads the file in bytes so return in bytes only.we have to chnage it into string
	fmt.Println("text data inside file is \n:", string(databytes))
}
