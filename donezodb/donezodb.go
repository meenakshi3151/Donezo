package donezodb

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	const donezodb string = "doneZo.db"
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

func CheckEmail(db *sql.DB, email string) bool {
	var count int
	// this value 2  will be assigned to the count variable via the Scan method.
	err := db.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", email).Scan(&count)
	if err != nil {
		fmt.Println(count)
		fmt.Println("Error checking for duplicate email:", err)
		return false
	}
	if count > 0 {
	return false
	}
	return true
}

func InsertUser(db *sql.DB, firstName string, lastName string, email string, password string) {
	_,err := db.Exec("INSERT INTO user (First_name, Last_name, Email, Password) VALUES (?, ?, ?, ?)", firstName, lastName, email, password)
	if err != nil {
		fmt.Println("Error inserting user:", err)
	} else {
	fmt.Println("User inserted successfully. ")
	}
}

func LoginUser(db *sql.DB, email string, password string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user WHERE Email = ? AND Password = ?", email, password).Scan(&count)
	if err != nil {
		fmt.Println("Error checking for user:", err)
		return false
	}
	if count > 0 {
		fmt.Print("User found. ")
	} else {
		fmt.Println("User not found. Provide correct email and password")
		return false
	}
	return true
}

func GetTasksAndStatus(db *sql.DB, task *[][]string, email string) {
	err := db.QueryRow("SELECT Name, Status FROM task WHERE Email = ?", email).Scan(&task)
	if err != nil {
		fmt.Print("Not getting tasks and status:", err," ")
	}
}

func InsertTask(db *sql.DB, tasks [][]string) {
	for _, task := range tasks { 
		_, err := db.Exec("INSERT INTO task (Name, Email) VALUES (?,?)", task)
		if err!= nil {
			fmt.Println("Error inserting task:", err)
		}
	}
}

