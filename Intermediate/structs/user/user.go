package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (user User) OutputUserData() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

// clearUserData clears the user data, need to pass a pointer to the User struct otherwise copy of the struct will be passed
func (user *User) ClearUserData() {
	user.firstName = ""
	user.lastName = ""
	user.birthDate = ""
}

func New(firstName string, lastName string, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		fmt.Println("Invalid user data!")
		return nil, errors.New("Invalid user data!")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}
