package main

import (
	"fmt"
	"strings"
)

func main() {

	var confernceName = "Go Confernce"
	const totalTickets = 100
	var remainingTickets uint = 100
	var bookings []string

	fmt.Printf("Welcome to %v\n", confernceName)
	fmt.Printf("We have a total of %v tickets\n", totalTickets)
	for {

		var userTickets uint
		var firstName string
		var lastName string
		var email string

		fmt.Printf("Enter your first name\n")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name\n")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email address\n")
		fmt.Scan(&email)
		fmt.Printf("Enter the tickets you require\n")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a conformation regarding this on your email address %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We now have %v remaining tickets from a total of %v tickets\n", remainingTickets, totalTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are the first names of all the clients: %v\n", firstNames)

			if userTickets == 0 {
				fmt.Printf("The %v is house full\n", confernceName)
				break
			}
		} else {
			fmt.Printf("We only have %v tickets so you cannot book %v tickets\n", remainingTickets, userTickets)
		}
	}
}
