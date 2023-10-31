package main

import (
	"fmt"
	"strings"
)

const confereceName = "Go conference"
const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidTickets, isValidEmail := validateUserInput(firstName, lastName, email, uint(userTickets))

		if isValidTickets && isValidEmail && isValidName {

			bookTicket(userTickets, bookings, firstName, lastName, email)
			firstNames := firstNames()
			fmt.Printf("These are all our bookings %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next Yeark")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("Frist name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}

			if !isValidTickets {
				fmt.Println("Number of ticketes you entered is invalid")
			}
			continue

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", confereceName)
	fmt.Printf("%v tickets remaning for %v.\n", remainingTickets, confereceName)
	fmt.Println("Get Your tickets here to attend")
}

func firstNames() []string {
	fristNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		fristNames = append(fristNames, names[0])
	}
	return fristNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidName, isValidTickets
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your frist name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, bookings []string, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confereceName)
}
