package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFileName = "balance.txt"

func readBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFileName)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountBalanceFileName, []byte(balanceText), 0644)
}

func main() {
	fmt.Println("Welcome to ATM machine!")
	var balance float64 = readBalanceFromFile()

	for {
		fmt.Println()
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println()

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
			writeBalanceToFile(balance)

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
			writeBalanceToFile(balance)

			fmt.Println("Your new balance is: ", balance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using ATM machine!")
			return
		}
	}
}
