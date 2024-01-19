package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"example.com/events-app/db"
)

type Registration struct {
	Id        int64      `binding:"required" json:"id"`
	EventId   int64      `binding:"required" json:"event_id"`
	UserId    int64      `binding:"required" json:"user_id"`
	Event     *Event     `json:"event,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedOn *time.Time `json:"created_on,omitempty"`
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
}

type Event struct {
	Id            int64          `json:"id"`
	Name          string         `json:"name" binding:"required"`
	Description   string         `json:"description" binding:"required"`
	Location      string         `json:"location" binding:"required"`
	Price         float64        `json:"price" binding:"required"`
	DateTime      string         `json:"date_time"`
	UserId        int64          `json:"user_id"`
	Registrations []Registration `bindings:"required" json:"registrations"`
}

func NewEvent(name, description, location string, price float64, userId int64) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		Price:       price,
		UserId:      userId,
		DateTime:    "",
	}
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

func ScanRegistrationFromDBRow(row db.DB_ROW) (*Registration, error) {

	var registration Registration
	var userData string
	var eventData string

	err := row.Scan(&registration.Id, &registration.EventId, &eventData, &registration.UserId, &userData)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(userData), &registration.User)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(eventData), &registration.Event)

	return &registration, err
}

func ScanEventFromDBRow(row db.DB_ROW, scan_registrations bool) (*Event, error) {

	var event Event
	var date_time []uint8
	var registartions string

	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Price, &date_time, &event.UserId)

	if err != nil {
		return nil, err
	}

	if scan_registrations {
		err := row.Scan(&registartions)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(registartions), &event.Registrations)
		if err != nil {
			return nil, err
		}
	}

	event.DateTime = string(date_time)

	return &event, nil
}

func GetAllEvents() ([]*Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events = []*Event{}

	for rows.Next() {
		event, err := ScanEventFromDBRow(rows, false)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error) {

	query := "SELECT * FROM events WHERE id=?"

	row := db.DB.QueryRow(query, eventId)

	event, err := ScanEventFromDBRow(row, false)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func UpdateEventById(eventId int64, created_by int64, event *Event) error {

	query := `
		UPDATE events
		SET name=?, description=?, location=?, price=?, user_id=?
		WHERE id=? and user_id=?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	sql_res, err := stmt.Exec(event.Name, event.Description, event.Location, event.Price, 3, eventId, created_by)

	if err != nil {
		return err
	}

	rows, err := sql_res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no such event found")
	}

	return nil
}

func DeleteEventById(eventId int64) error {
	query := "DELETE FROM events WHERE id=?"

	sql_res, err := db.DB.Exec(query, eventId)

	if err != nil {
		return err
	}

	rows, err := sql_res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no such event found")
	}

	return nil

}

func (event *Event) GetAllRegistrations() error {
	query := `
		SELECT
			JSON_ARRAYAGG(
				JSON_OBJECT(
					'id', r.id,
					'user_id', u.id,
					'event_id', e.id,
					'user', JSON_OBJECT(
						'id', u.id,
						'email', u.email
					)
				)
			) as registrations
		FROM events e
		LEFT JOIN registrations r ON r.event_id=e.id
		LEFT JOIN users u ON r.user_id=u.id
		WHERE e.id=?
		GROUP BY e.id;
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	sql_row := stmt.QueryRow(event.Id)

	var registrations string
	err = sql_row.Scan(&registrations)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(registrations), &event.Registrations)

	var updatedRegistrations []Registration

	for _, registartion := range event.Registrations {
		if registartion.Id != 0 {
			updatedRegistrations = append(updatedRegistrations, registartion)
		}
	}

	event.Registrations = updatedRegistrations

	return err
}

func (event *Event) RegisterUser(user_id int64) (int64, error) {
	query := `
		INSERT INTO registrations(user_id, event_id)
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	sql_res, err := stmt.Exec(user_id, event.Id)

	if err != nil {
		return 0, err
	}

	registration_id, err := sql_res.LastInsertId()

	return registration_id, err
}

func (event *Event) DeleteUserRegistration(user_id int64) error {
	query := `
		DELETE FROM registrations
		WHERE user_id=? and event_id=?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	sql_res, err := stmt.Exec(user_id, event.Id)

	if err != nil {
		return err
	}

	rows_deleted, err := sql_res.RowsAffected()

	if err != nil {
		return err
	}

	if rows_deleted == 0 {
		return errors.New("no such registartion exists")
	}

	return nil
}
