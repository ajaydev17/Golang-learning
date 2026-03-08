package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthDate := getUserData("Enter your birth date (YYYY-MM-DD): ")

	var appUser User

	appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	appUser.outputUserData()
}

func (user User) outputUserData() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}
