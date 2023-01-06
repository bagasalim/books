package helper

import "strings"

func ValidateInputUser(firstName string, lastName string, email string, userTickets uint, RemainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets <= RemainingTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTickets
}