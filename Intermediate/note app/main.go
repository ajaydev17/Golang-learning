package main

import (
	"bufio"
	"example/note-app/note"
	"example/note-app/todo"
	"fmt"
	"os"
	"strings"
)

// interface for note management
type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputable interface {
	saver
	displayer
}

// This is a simple note-taking application in Go. It allows users to create, view, and delete notes.
func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todoItem, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Failed to create todo item:", err)
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Failed to create note:", err)
		return
	}

	fmt.Println("Todo item created successfully!")

	err = outputData(todoItem)
	if err != nil {
		fmt.Println("Failed to save todo item:", err)
		return
	}

	fmt.Println("Todo item saved successfully!")

	fmt.Println("Note created successfully!")

	err = outputData(note)
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

func outputData(d outputable) error {
	d.Display()
	return saveData(d)
}

func saveData(s saver) error {
	err := s.Save()

	if err != nil {
		fmt.Println("Failed to save data:", err)
		return err
	}
	fmt.Println("Data saved successfully!")
	return nil
}

func getTodoData() string {
	return getUserInput("Enter the todo item: ")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of your note: ")
	content := getUserInput("Enter the content of your note: ")
	return title, content
}
