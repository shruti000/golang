package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

// this si the first stage.It is sued to store the charcters of the string in the channel name d out
func generateChar(out chan byte, s string) {
	fmt.Println("Inside first channel")
	for _, c := range s {
		out <- byte(c)
	}
}

// this isused to remove the digit frot he channel and stoere it in ather channel
func removeDigit(in chan byte, out chan byte) {
	fmt.Println("Inside second Channel")
	digits := "0123456789"
	//for of type value ok
	for {
		value, ok := <-in
		if ok {
			if !strings.Contains(digits, string(value)) {
				out <- value
			} else {
				break
			}
		}
	}
}

// this cahnnel is sued to convert to upper case
func toUpper(in chan byte) {
	fmt.Println("inside third channel")
	for {
		value, ok := <-in
		if ok {
			fmt.Printf("%v ", string(unicode.ToUpper(rune(value))))
		} else {
			break
		}
	}
}

func main() {

	chanA := make(chan byte)
	chanB := make(chan byte)
	go generateChar(chanA, os.Args[1])
	go removeDigit(chanA, chanB)
	go toUpper(chanB)
	time.Sleep(10 * time.Millisecond)

}
