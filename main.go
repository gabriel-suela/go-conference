package main

import (
	"fmt"
	"strconv"
	"strings"
)

const confereceName = "Go conference"
const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

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
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confereceName)
}
