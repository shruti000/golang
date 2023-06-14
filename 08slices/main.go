package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	//define empty slice
	var test = []string{}
	fmt.Println(test)

	var fruitList = []string{
		"mango",
		"banana",
		"apple",
		"peach",
		"watermelon",
	}
	fmt.Println(fruitList)

	//to add in slice we use append fucntion
	fruitList = append(fruitList, "guava", "pineapple")
	fmt.Println(fruitList)

	//now we will slice eg [1:3] where 3 is inclusive
	fruitList = fruitList[:3]
	fmt.Println(fruitList)

	veg := make([]int, 5)
	veg[0] = 23
	veg[1] = 13
	veg[4] = 30
	//here if we try tomadd more vlue it will give error
	fmt.Println(veg)

	//here if we try to add more value using append no error
	veg = append(veg, 55, 88, 77, 99, 00)
	fmt.Println(veg)

	//lets see the sorting part
	//we are usig sort fucntion
	//ints will sort the integers
	sort.Ints(veg)
	fmt.Println(veg)

	//sprt the strings
	sort.Strings(fruitList)
	fmt.Println(fruitList)

	//remove from a slice i.e delete
	//... basically reprsents there will be more varibale
	var courses = []string{
		"java",
		"python",
		"go",
		"ruby",
		"angular",
	}
	fmt.Println(courses)
	fmt.Println("delete operation are :")
	courses = append(courses[:1], courses[2:]...)
	fmt.Println(courses)
}
