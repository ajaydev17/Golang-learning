package main

import (
	"bufio"
	"example/note-app/note"
	"fmt"
	"os"
	"strings"
)

// This is a simple note-taking application in Go. It allows users to create, view, and delete notes.
func main() {
	title, content := getNoteData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Failed to create note:", err)
		return
	}
	fmt.Println("Note created successfully!")
	fmt.Println(note.Display())

	err = note.Save()
	if err != nil {
		fmt.Println("Failed to save note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}

// getUserInput prompts the user for input and returns the entered string.
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r") // Handle Windows line endings
	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of your note: ")
	content := getUserInput("Enter the content of your note: ")
	return title, content
}
