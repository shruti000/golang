package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //here we can chnage the name just like @name in spring
	Price    int
	Platform string   `json:"website"` //we give values in bactop so the new name in the utput eill ne website
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` //here omitempty mean if the array is empty then it will not print it in the console lonly
}

func main() {
	fmt.Println("Welcome to JSON video")
	//EncodeJson()
	DecodeJson()
}

// function basiclaly to cnvert any obejct to json
func EncodeJson() {

	//creating an array of stucts to convert to json
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	//package this data as JSON data
	//to convert to json we can use json packge
	//thejson packge consit of marshal keyword to convert any oject to json here we are usig stuct
	//here we have used marshalUtent to print it is a mpre formatted way
	// (name of array of stuct,in the spcae there shuld be spce,\t to guve tab space)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(finalJson))
	//fmt.Printf("type is %T ", finalJson)
	//the putput is in the for  of byte
	fmt.Printf("%s\n", finalJson)

}

// this method is basically used to convert json to other objets
func DecodeJson() {

	//as json data are of byet form
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	//now we have to convert to either structor any other obejts
	//1]first check whether it is valid or notconst
	isValid := json.Valid(jsonDataFromWeb)
	var jsonToStruct course

	if isValid {
		fmt.Println("json is valid")
		//unmarshl basiclly convert from json to other object
		//it takes two variable the json to comvert and the strcut and for struct we store the copy also so &
		json.Unmarshal(jsonDataFromWeb, &jsonToStruct)
		//to print struct we use %#v
		fmt.Printf("%#v\n", jsonToStruct)
	} else {
		fmt.Println("Json is not valid")
	}

	fmt.Println("********************JSON TO MAP****************")

	//now to print data in maps form i.e key value form
	//here he key is string but value can be anything s we used interface which acts as any
	var jsonToMap map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &jsonToMap)
	fmt.Printf("%#v\n", jsonToMap)

	for k, v := range jsonToMap {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}
}
