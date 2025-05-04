package main

import (
	"fmt"
	"go-learn/helper1" // why we use go-learn because we using our own packages, go-mod is import functions
	"strconv"
	"strings"
)

// struct holds the different data types in the key values pairs
type UserData struct { // UserData is name of struct, type will be provide new custom data type
	firstName string
	lastName  string
	email     string
	phone     uint
}

// package level scope we need to declare here only actual declaration, not the syentatic sugar syntax
var sharedVariableBTWFunc = "string values" // same like this we can share

func main() {
	// arrays its fixed length array
	var bookings = [3]string{"santhosh kumar", "kum ar", "santh osh"}

	var bookingsWithMap = make([]map[string]string, 1)

	var bookingsWithStruct = make([]UserData, 0)

	// create a map for a user
	var userData = make(map[string]string) // we can't mix the data type

	var userDataWithStruct = UserData{
		firstName: "santhoshkumar",
		lastName:  "Viswanathan",
		email:     "santhoshkumar@ownBrand.com",
		phone:     9889898899,
	}

	bookingsWithStruct = append(bookingsWithStruct, userDataWithStruct)

	userData["firstName"] = "Santhosh"
	userData["lastName"] = "Kumar"
	userData["email"] = "santhosh@mai;.com"
	userData["tickets"] = strconv.FormatUint(uint64(23), 10)

	fmt.Println("bookingsWithMap", bookingsWithMap)
	fmt.Println("bookingsWithMap", bookingsWithMap)
	fmt.Println("userData", userData)
	fmt.Println("userData firstName", userData["firstName"]) // for map we need access the property with array notations

	fmt.Println("userDataWithStruct", userDataWithStruct)
	fmt.Println("userDataWithStruct firstName", userDataWithStruct.firstName) // for struct we need access the property with dot notations
	fmt.Println("bookingsWithStruct", bookingsWithStruct)

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
	fmt.Println("Println", dummy2, helper1.Tickes)

	helper1.FunctionFromAnotherPackage()
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

// func functionFromAnotherPackage() {
// 	fmt.Println("Dummy package")
// } // same file declarations
