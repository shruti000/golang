package main

import (
	"fmt"
	"net/http"
)

func handleGoServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello welocme to web development in go")
}

func main() {

	http.HandleFunc("/goserver", handleGoServer)
	http.ListenAndServe("127.0.0.1:5000", nil)
}
