package main

import "fmt"

//arays in go are of fixed lenght

func main() {
	fmt.Println("welcome to arrays")

	var fruitList [3]string
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[2] = "cherry"
	//fruitList[4] = "cherry"
	fmt.Println("the frutList is ", fruitList)
	fmt.Println("the frutList  lenght is ", len(fruitList))

	var vegeList = [2]string{
		"potato", "tomato",
	}
	fmt.Println(vegeList)
}
