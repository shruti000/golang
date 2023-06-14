package main

import (
	"fmt"
	"net/http"
)


func handleGoServer(w http.ResponseWriter, r *http.Request) {
       //here http.Reqest will get the resposne header
       //print will format the string and also wrte in the writer
	fmt.Fprintf(w, "hello welocme to web development in go")
}

func AnotherGoServer(w http.ResponseWriter, r *http.Request) {
       //here http.Reqest will get the resposne header
       //print will format the string and also wrte in the writer
	fmt.Fprintf(w, "welcone tanother path of web developemt")
}

func main() {

        //regsters the handler fucntionwith the router
	http.HandleFunc("/goserver", handleGoServer)


      http.HandleFunc("/anothergoserver", AnotherGoServer)


      //here nil is the default multiplexer
       //it will start at default port
	http.ListenAndServe("127.0.0.1:5000", nil)
}
