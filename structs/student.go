package main

import (
	"time"

	"example.com/app/user"
)

type Student struct {
	id   int
	User user.User
}

func NewStudent(id int) (*Student, error) {

	user, err := user.New("StudentF", "StudentL", "StudentB", time.Now())

	if err != nil {
		return nil, err
	}

	return &Student{
		id:   id,
		User: *user,
	}, nil
}
