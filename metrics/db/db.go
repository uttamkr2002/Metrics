// db/db.go
package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Initialize DB connection
func InitDB() {
	var err error
	connStr := "user=postgres password=1234 dbname=metrics_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}
	fmt.Println("Connected to the PostgreSQL database!")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
