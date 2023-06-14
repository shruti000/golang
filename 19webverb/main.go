package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("welcome to web verb video ")
	fmt.Println()
	fmt.Println("Get Request")
	PerformGetRequest()
	fmt.Println()
	fmt.Println("Post Request")
	PerformPostRequest()
	fmt.Println()
	fmt.Println("post Form Request")
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("the reposne status of the code is ", response.StatusCode)
	fmt.Println("the lenght of the reposne is ", response.ContentLength)

	//here we convert the byte data to string
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	//another method to print the byte data of the content to string
	//this method is by uisng string builder of strings class
	var byteToString strings.Builder
	byteCount, _ := byteToString.Write(content)

	fmt.Println("ByteCount is ", byteCount)
	fmt.Println("content in string is ", byteToString.String())

}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	//fake jason data for post body
	//for this we can the strings inary which has a method os NewReader which can take values of any dattype

	requestBody := strings.NewReader(`
	{
		"coursename":"golang",
		"price":300,
		"platform":"youtube.com"
	}
	`)

	response, _ := http.Post(myurl, "application/json", requestBody)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//create a fake form data
	//this can be done using url package then using .Values
	//values empty for now then we will add key-value pairs in it
	data := url.Values{}
	data.Add("firstname", "shruti")
	data.Add("lastname", "koli")
	data.Add("age", "18")

	response, _ := http.PostForm(myurl, data)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
