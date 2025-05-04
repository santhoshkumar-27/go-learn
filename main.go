package main

import (
	"fmt"
	"strings"
)

// global level scope we need to declare here only actual declaration, not the syentatic sugar syntax
var sharedVariableBTWFunc = "string values" // same like this we can share

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

	// slice its dynamic
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

	fmt.Println("Dummy data", firstNames, len(firstNames), strings.Contains(firstNames[0], "@"), true && true, false || false, 1 >= 11, 1 != 0)

	switch i {
	case 12:
		fmt.Print("Switch statme single value")
	case 5, 6, 8:
		fmt.Println("switch statement printed")
		break
	default:
		fmt.Println("Default statment")
	}

	greetUsers()

	greetUsersValues("thnaku ")

	singleReturnValue()

	dummy, _, _ := multipleReturnValues() // ignoring first and second values

	_, dummy1, _ := multipleReturnValues() // ignoring first and second values

	_, _, dummy2 := multipleReturnValues() // ignoring first and second values

	fmt.Println("Println", dummy)
	fmt.Println("Println", dummy1)
	fmt.Println("Println", dummy2)

}

// if we want to run mulitple files
// go run main.go helper.go
// or
// go run . -> folder name to run entire file

// if condition {
// } else if condition {
// } else {
// }

func greetUsers() { // normal functions
	fmt.Println("Welcome to the main area of function, can call the encapsulation code")
}

func greetUsersValues(paramName string) { // normal functions
	fmt.Println("function with parameters", paramName)
}

func singleReturnValue() string {
	return sharedVariableBTWFunc
}
