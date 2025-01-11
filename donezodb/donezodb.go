package donezodb

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	const donezodb string = "donezo.db"
	db, err := sql.Open("sqlite3", donezodb)
	if err != nil {
		fmt.Println("Error opening the database:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}

	fmt.Println("DB connection successful")
	return db, nil
}

func CheckEmail (db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		fmt.Println("Error checking for duplicate email:", err)
		return false
	}
	return count == 0
}
