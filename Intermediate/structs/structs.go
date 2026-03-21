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

	var appUser *User

	appUser = newUser(userFirstName, userLastName, userBirthDate)

	appUser.outputUserData()
	appUser.clearUserData()
	appUser.outputUserData()
}

func (user User) outputUserData() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

// clearUserData clears the user data, need to pass a pointer to the User struct otherwise copy of the struct will be passed
func (user *User) clearUserData() {
	user.firstName = ""
	user.lastName = ""
	user.birthDate = ""
}

func newUser(firstName string, lastName string, birthDate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}
