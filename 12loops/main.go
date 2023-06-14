package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	days := []string{
		"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday",
	}
	for d := 0; d <= len(days)-1; d++ {
		fmt.Println(days[d])
	}
	fmt.Println()
	for _, value := range days {
		fmt.Println(value)
	}

	fmt.Println()
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("******************************")
	roughValue := 7
	for roughValue < 5 {
		fmt.Println("value less than 5 ", roughValue)
		roughValue++

		if roughValue == 7 {
			goto loc
		}
	}

loc:
	fmt.Println("testing the goto part")
}
