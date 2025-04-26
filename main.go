package main

import (
	"fmt"
	"strings"
)

func main() {

	var bookings = [3]string{"santhosh kumar", "kum ar", "santh osh"}

	// infinite loop
	// for {
	// 	fmt.Println("for loop for ")

	// }

	var incrementer uint

	// normal loop with break
	for {
		fmt.Println("for loop for ")

		incrementer += 1

		if incrementer == 5 {
			break // break will terminate the loop
			// continue // will skip the current iteration and skip to next iteratio
		}
	}

	firstNames := []string{}

	// for-each
	// basically it has for index, booking := range bookings
	// if i use underscore it represents as blank operator
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	var i = 0

	for i < 5 {
		fmt.Println("condtions", i)

		i++
	}

	fmt.Println("Dummy data", firstNames, len(firstNames))

}

// if condition {
// } else if condition {
// } else {
// }
