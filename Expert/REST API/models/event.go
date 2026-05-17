package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Failed to prepare statement: " + err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		panic("Failed to execute statement: " + err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic("Failed to retrieve last insert ID: " + err.Error())
	}

	e.ID = int(id)
	return err
}

func GetAllEvents() []Event {
	return events
}
