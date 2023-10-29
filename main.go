package main

import (
	"fmt"
	"strings"
)

func main() {
	confereceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", confereceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your tickets here to attend")

	for {
		var fristName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your frist name: ")
		fmt.Scan(&fristName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want? ")
		fmt.Scan(&userTickets)

		if userTickets <= int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, fristName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", fristName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaning for %v.\n", remainingTickets, confereceName)

			fristNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				fristNames = append(fristNames, names[0])
			}
			fmt.Printf("These are all our bookings %v\n", fristNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next Yeark")
				break
			}

		} else {

			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			continue

		}

	}

}
