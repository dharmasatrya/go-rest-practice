package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresStorage(connStr string) (*sql.DB, error) {
	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	// Check if the connection is actually working
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	// Connection was successful
	fmt.Println("Successfully connected to the database!")

	// Return the DB object for further use
	return db, nil
}
