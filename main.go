package main

import "fmt"

var conferenceName = "Go Conference"

var conferenceName1 string = "Go Conference"

const conferenceTicket = 50

var userName string
var userAge uint

func main() {
	dummmssssy := "value" // only apply for var

	fmt.Printf("Conference is %T \n %v", userName, dummmssssy)
	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("Get you ticket from here to watch IPl")
	fmt.Println(conferenceName, conferenceTicket)

	fmt.Printf("replace value %v ", conferenceName1)

	fmt.Printf("adsfdsaf %v", 1+2)

	// fmt.Scan(userName)  // we are just passing values only here so they can't stop
	fmt.Scan(&userName) // we are passing as pointer of (memory address) but here they wait

	// default value will be empty string "",
	// userName = "Santhosh"
	// default value will be 0,
	userAge = 26

	fmt.Println(userAge)
	// fmt.Println(&userAge) // Pointer it will return the memory address

	fmt.Print(userName, userAge)
}
