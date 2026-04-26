package main

import (
	"errors"
	"fmt"
)

// This is a simple note-taking application in Go. It allows users to create, view, and delete notes.
func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Failed to get note data:", err)
		return
	}
	fmt.Printf("Note created!\nTitle: %s\nContent: %s\n", title, content)
}

// getUserInput prompts the user for input and returns the entered string.
func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)

	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	return input, nil
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter the title of your note: ")
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	content, err := getUserInput("Enter the content of your note: ")
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	return title, content, nil
}
