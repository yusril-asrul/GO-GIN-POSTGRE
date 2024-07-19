package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	connStr := "user=postgres dbname=belajar password=12345678 sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        code VARCHAR(50) NOT NULL,
        title VARCHAR(100) NOT NULL,
        author VARCHAR(100) NOT NULL
    );`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
