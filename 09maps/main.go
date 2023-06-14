package main

import "fmt"

func main() {
	fmt.Println("welcome to maps")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["py"] = "python"
	languages["RB"] = "ruby"
	fmt.Println(languages)
	//retrieve
	fmt.Println(languages["JS"])

	//add to mao
	languages["go"] = "golang"
	fmt.Println(languages)

	//delete
	delete(languages, "go")
	fmt.Println(languages)

	//forloop in map
	for index, value := range languages {
		fmt.Println("key is ", index, "value is ", value)
	}

}
