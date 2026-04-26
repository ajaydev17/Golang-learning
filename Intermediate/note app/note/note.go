package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Note) Display() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nCreated At: %s", n.Title, n.Content, n.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	jsonData, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error encoding note to JSON:", err)
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s.json", fileName), jsonData, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("title and content cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
