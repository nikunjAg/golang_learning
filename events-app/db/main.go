package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type DB_ROW interface {
	Scan(...any) error
}

const (
	DRIVER_NAME = "mysql"
	DB_USER     = "nikunj"
	DB_PASS     = "ttn123"
	DB_PROTOCOL = "tcp(127.0.0.1:3306)"
	DB_NAME     = "my_events"
)

func createTables() error {

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id int(12) AUTO_INCREMENT,
			email varchar(300) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			CONSTRAINT uk_users_email UNIQUE KEY(email)
		) ENGINE=InnoDB DEFAULT CHARSET=latin1;
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		return err
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id int(12) AUTO_INCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			price DECIMAL(10, 2) NOT NULL,
			date_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			user_id int(11),
			PRIMARY KEY (id),
			CONSTRAINT fk_events_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=latin1;
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		return err
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id int(12) AUTO_INCREMENT NOT NULL,
			user_id int(12) NOT NULL,
			event_id int(12) NOT NULL,
			created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY(id),
			CONSTRAINT uk_registrations_user_id_event_id UNIQUE KEY (user_id, event_id),
			CONSTRAINT fk_registrations_user_id FOREIGN KEY (user_id) REFERENCES users(id),
			CONSTRAINT fk_registrations_event_id FOREIGN KEY (event_id) REFERENCES events(id)
		) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=15000;
	`
	_, err = DB.Exec(createRegistrationsTable)

	return err
}

func InitDB() (*sql.DB, error) {

	if DB != nil {
		return DB, nil
	}

	DATA_SOURCE_NAME := fmt.Sprintf("%v:%v@%v/%v", DB_USER, DB_PASS, DB_PROTOCOL, DB_NAME)

	db, err := sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)

	if err != nil {
		return nil, err
	}

	DB = db

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(time.Minute * 3)

	err = createTables()

	return DB, err
}
