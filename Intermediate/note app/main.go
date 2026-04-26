package main

import (
	"example/note-app/note"
	"fmt"
)

// This is a simple note-taking application in Go. It allows users to create, view, and delete notes.
func main() {
	title, content := getNoteData()

	_, err := note.New(title, content)
	if err != nil {
		fmt.Println("Failed to create note:", err)
		return
	}
	fmt.Println("Note created successfully!")
}

// getUserInput prompts the user for input and returns the entered string.
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of your note: ")
	content := getUserInput("Enter the content of your note: ")
	return title, content
}
