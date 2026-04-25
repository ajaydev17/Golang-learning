package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := user.GetUserData("Enter your first name: ")
	userLastName := user.GetUserData("Enter your last name: ")
	userBirthDate := user.GetUserData("Enter your birth date (YYYY-MM-DD): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println("Error creating user: ", err)
		return
	}

	appUser.OutputUserData()
	appUser.ClearUserData()
	appUser.OutputUserData()

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OutputUserData()
	admin.ClearUserData()
	admin.OutputUserData()
}
