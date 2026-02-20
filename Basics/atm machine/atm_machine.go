package main

import (
	"example.com/atm-machine/fileops"
	"fmt"
)

const accountBalanceFileName = "balance.txt"

func main() {
	fmt.Println("Welcome to ATM machine!")
	var balance, err = fileops.ReadFloatFromFile(accountBalanceFileName, 1000.00)

	if err != nil {
		fmt.Println("Error reading balance from file: ", err)
		panic("Cannot continue, Sorry!")
	}

	for {
		presentMenu()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		fmt.Println("You have chosen: ", choice)
		fmt.Println()

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", balance)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount!, please enter a positive amount.")
				continue
			}

			balance += amount
			fileops.WriteFloatToFile(accountBalanceFileName, balance)

			fmt.Println("Your new balance is: ", balance)
		case 3:
			fmt.Print("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount!, please enter a positive amount.")
				continue
			}

			if amount > balance {
				fmt.Println("Invalid amount!, account balance is: ", balance)
				continue
			}

			balance -= amount
			fileops.WriteFloatToFile(accountBalanceFileName, balance)

			fmt.Println("Your new balance is: ", balance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using ATM machine!")
			return
		}
	}
}
