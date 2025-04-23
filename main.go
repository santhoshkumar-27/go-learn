package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket uint = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// fixed size array length values
	// var bookings = [50]string{"Santhoshkumar viswanathan"}
	// var bolling [50]string
	bookings := [50]string{"Santhoshkumar viswanathan"}

	// dynamic array length
	// var bookings1 = []string{}
	// var bookings1 []string
	bookings1 := []string{}

	var fullName string
	var age uint

	fmt.Println("Please enter full name")
	fmt.Scan(&fullName)
	fmt.Println("Please enter age")
	fmt.Scan(&age)

	bookings[0] = fullName

	bookings1 = append(bookings1, fullName)

	fmt.Printf("Hi %v this your name, and this is your age %v \n", fullName, age)

	// fixed length
	fmt.Printf("The whole array %v \n", bookings)
	fmt.Printf("The first array %v \n", bookings[0])
	fmt.Printf("The type of  array %T \n", bookings)
	fmt.Printf("The size of  array %v \n", len(bookings))

	// dynamic length
	fmt.Printf("The whole array %v \n", bookings1)
	fmt.Printf("The first array %v \n", bookings1[0])
	fmt.Printf("The type of  array %T \n", bookings1)
	fmt.Printf("The size of  array %v \n", len(bookings1))

}
