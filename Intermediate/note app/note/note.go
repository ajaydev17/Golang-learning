package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (n *Note) Display() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nCreated At: %s", n.title, n.content, n.createdAt.Format("2006-01-02 15:04:05"))
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("title and content cannot be empty")
	}

	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
