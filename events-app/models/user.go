package models

import (
	"example.com/events-app/db"
	"example.com/events-app/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUser(Email, Password string) *User {
	return &User{
		Email:    Email,
		Password: Password,
	}
}

func (user *User) Save() error {

	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	sql_res, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := sql_res.LastInsertId()

	user.Id = id

	return err
}
