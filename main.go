package main

import (
	"booking-apps/helper"
	"fmt"
	"sync"
	"time"
)

const conferanceTickets int = 50
var RemainingTickets uint = 50
var conferanceName = "Go Conference"
var bookings =make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()


		
		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTickets := helper.ValidateInputUser(firstName, lastName, email, userTickets, RemainingTickets)
		
		if isValidName && isValidEmail && isValidTickets {
			
			bookTicket( userTickets,  firstName, lastName, email)

			wg.Add(1)
			go sendTicket( userTickets,  firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Printf("These all are our attendances %v\n", firstNames)

			if RemainingTickets == 0 {
				fmt.Println("Fully Booked, So Sorry")
			}
			
		} else {
			if !isValidName {
				fmt.Println("Name was too short")
			} else if !isValidEmail {
				fmt.Println("Input the right email")
			}else 	if !isValidTickets {
				fmt.Println("Input valid tickets")
			}
		}
		wg.Wait()
	
}

func greetUser() {
	fmt.Printf("Welcome to %v booking applications\n", conferanceName)
	fmt.Printf("We have total %v tickets and %v still available\n", conferanceTickets, RemainingTickets)
	fmt.Println("Get your ticket to attend the show.")
}

func getFirstName() []string{
	firstNames := []string{}
	for _, booking := range bookings {
	firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Fill your identity to book the ticket")

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket( userTickets uint,  firstName string, lastName string, email string) {
	RemainingTickets = RemainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tikets, You will receive a details book on %v\n", firstName, lastName, userTickets, email)
	if RemainingTickets != 0 {
		fmt.Printf("%v tickets remaining for %v\n", RemainingTickets, conferanceName)
	}
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
	fmt.Printf("#############\n")
	fmt.Printf("Sending ticket:\n %v \nto email address %v", ticket, email)
	fmt.Printf("#############\n")
	wg.Done()
}