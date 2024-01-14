package models

import (
	"errors"

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

func (user *User) Login() error {

	query := `
		SELECT id, password
		FROM users
		WHERE email=?
	`

	sql_row := db.DB.QueryRow(query, user.Email)

	var userId int64
	var hashedPassword string
	err := sql_row.Scan(&userId, &hashedPassword)

	if err != nil {
		return err
	}

	res := utils.ComparePassword(hashedPassword, user.Password)

	if !res {
		return errors.New("credentails mismatched")
	}

	user.Id = userId

	return nil
}
