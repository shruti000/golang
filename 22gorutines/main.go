package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //basically this is a pointer
//wait group consiste of 3 methods add,done,and wait
//add bascially acts there are not how ro ruitne like add(1)
//done()  its responsibility is to basiclaly send the signal that the go routine is done
//wait() function is mostly at the end of main method or fucntion which will tell the main method not to stop

// slice of strings
var signals = []string{"test"}

// using for locking purpose
var mut sync.Mutex //it is a pointer

func main() {
	//go words basically resembles that it runs but it forgets to come back and the nextfucnt get exceuted and also ended
	go greeter("shruti")
	go greeter("pranav")

	website := []string{
		"https://lco.dev",
		"https://google.com",
		"https://youtube.com",
	}
	for _, value := range website {
		go getStatusCode(value)
		//to keep the track of go routine likem how many go routinrs is working
		wg.Add(2)

	}
	//mpsty at the end to stop the main fucntion from ending
	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("welcome to Go Routine ", s)
	}
}

func getStatusCode(endpoint string) {
	//done besically helps to pass the signal once the go rotuine is complete
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("There was error in readin the endpoints")
	} else {
		//while writing in signal lock it
		mut.Lock()
		signals = append(signals, endpoint)
		//after writing to the signal
		mut.Unlock()
		fmt.Printf("%d status code for %s\n ", res.StatusCode, endpoint)
	}

}
