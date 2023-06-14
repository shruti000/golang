package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Testing the Race condition i.e mutex and wait Groups")

	wg := &sync.WaitGroup{}

	//as they are simple program it is happening smootly but we need to implemtn mutex for larger programs
	mut := &sync.Mutex{}

	//mut:=&sync.RWMutex    this is a read write mutex

	var score = []int{}

	//as we know there are 3 go routine so there will be 3 add
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		//ust lock it before operation adn unlock it after operation
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
