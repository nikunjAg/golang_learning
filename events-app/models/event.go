package models

import (
	"fmt"

	"example.com/events-app/db"
)

type Event struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Location    string  `json:"location" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	DateTime    string  `json:"date_time"`
	UserId      int64   `json:"user_id"`
}

func NewEvent(name, description, location string, price float64, userId int64) *Event {
	return &Event{
		Id:          1,
		Name:        name,
		Description: description,
		Location:    location,
		Price:       price,
		UserId:      userId,
		DateTime:    "",
	}
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events = []Event{}

	for rows.Next() {
		var event Event
		var date_time []uint8
		err = rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Price, &date_time, &event.UserId)

		event.DateTime = string(date_time)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (event *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, price, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	sql_res, err := stmt.Exec(event.Name, event.Description, event.Location, event.Price, event.UserId)

	if err != nil {
		return err
	}

	id, err := sql_res.LastInsertId()
	fmt.Println(id)
	event.Id = id

	return err
}
