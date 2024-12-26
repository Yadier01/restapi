package models

import (
	"github.com/Yadier01/rest-api/db"
	"github.com/Yadier01/rest-api/utils"
)

type Users struct {
	ID       int64
	UserName string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) Save() error {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.UserName, u.Email, hashedPass)
	return err
}

func GetUserByEmail(email string) (*Users, error) {
	query := `SELECT * FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, email)

	var u Users
	err := row.Scan(&u.ID, &u.UserName, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
