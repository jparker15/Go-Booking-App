package main

import "fmt"

func main() {
	// var conferenceName  string = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 //specifiing the type of data prevents this variable from becoming negative

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	// fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are stil available")
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets) //prints the memory address of the variable remainingTickets
	fmt.Println("- Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("- Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("- Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("- Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
