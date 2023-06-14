package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arrfiles, err := os.Create("test1.txt")
	if err != nil {
		panic(err)
	}
	defer arrfiles.Close()

	var arr = [5]int{1, 2, 3, 4, 5}
	for _, value := range arr {
		arrfiles.WriteString(strconv.Itoa(value))
		arrfiles.WriteString("\n")

	}
	fmt.Println("written sucessfully")
	os.Rename("test1.txt", "text2.txt")

}
