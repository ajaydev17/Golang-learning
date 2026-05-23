package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	passwordHash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, passwordHash)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = int64(id)

	return err
}

func (u *User) ValidateCredentials() (bool, error) {
	query := `SELECT id, email, password FROM users WHERE email = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var storedUser User
	err = stmt.QueryRow(u.Email).Scan(&storedUser.ID, &storedUser.Email, &storedUser.Password)
	if err != nil {
		return false, err
	}

	isMatch, err := utils.CheckPasswordHash(u.Password, storedUser.Password)
	if err != nil {
		return false, err
	}

	u.ID = storedUser.ID
	return isMatch, nil
}
