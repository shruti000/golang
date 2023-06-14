package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("welcome to web request part of golang")

	//this is used to get the reposne f the url
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	//this is used to read the content of the url
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//to print the content
	fmt.Println("the contents are : ", string(content))

	//always remeber to close the response body
	defer response.Body.Close()

}
