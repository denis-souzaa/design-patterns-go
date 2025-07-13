package config

import (
	"database/sql"
	"fmt"
	"log"
)

func New() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_DATABASE)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
