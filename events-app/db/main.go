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

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id int(12) PRIMARY KEY AUTO_INCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			price DECIMAL(10, 2) NOT NULL,
			date_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			user_id int(11) NOT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	createTables()

	return DB, nil
}
