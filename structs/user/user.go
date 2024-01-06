package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	birthDate string
	createdAt time.Time
}

// Constructor function Pattern
// its's just a normal function nothing special
// It just uses new the Constructor Pattern to create new User
// Can place validations and formatting on the data
func New(firstName, lastName, birthDate string, createdAt time.Time) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("please enter a valid data(Firstname, Lastname, and Birthdate are required)")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		birthDate: birthDate,
		createdAt: createdAt,
	}, nil
}

// Receiver for type "User"
func (user User) DisplayDetails() {

	fmt.Printf("Hey! %v %v, I know you you were born on %v. You told me about this on %v\n", user.FirstName, user.LastName, user.birthDate, user.createdAt.Local())
}

func (user *User) ClearUsername() {
	user.FirstName = ""
	user.LastName = ""
}

func (user User) compareUsers(userB User) {
	fmt.Println("Comparison Result", user == userB)
}
