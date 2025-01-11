package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	const donezodb string = "donezo.db"
	db, err := sql.Open("sqlite3", donezodb)
	if err != nil {
		fmt.Println("error opening the database",err)
	}
	defer db.Close() 
	if err := db.Ping(); err != nil {
		fmt.Println("Error connecting the database",err)
	}
	fmt.Println("DB connection successful")
}
