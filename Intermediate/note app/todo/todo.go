package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("Todo: %s\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding todo to JSON:", err)
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s.json", fileName), jsonData, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("text cannot be empty")
	}

	return Todo{
		Text: text,
	}, nil
}
