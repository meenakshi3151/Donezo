package main

import (
	"fmt"
	"strings"
	"ToDoList/donezodb"
	"database/sql"
)

var countOfDetails int = 0

func greetUsers() {
	fmt.Println("Welcome to one stop solution of your tasks")
}

func verifyEmail(email string) bool {
	var specialChar = []string{"!", "#", "$", "%", "^", "&", "*", "(", ")", "_", "-", "+", "=", "{", "}", "[", "]", ":", ";", "'", "<", ">", ",", ".", "?", "/", "|", "`", "~"}
	for _, char := range specialChar {
		if strings.Contains(email, char) {
			return false
		}
	}
	return strings.Contains(email, "@gmail.com")
}

func verification(firstName string, lastName string, email string, db *sql.DB) {
	if len(firstName) > 2 && len(lastName) > 2 && verifyEmail(email) {
		fmt.Println("You are verified")
	} else {
		fmt.Println("Please enter a valid name with at least 2 characters.")
		countOfDetails++
		if countOfDetails == 2 {
			fmt.Println("You have reached the limit of entering the details")
			return
		}
		getUserDetails(db)
	}
}

func checkDuplicateEmail(db *sql.DB,email string) {
   if(donezodb.CheckEmail(db, email)) {
	   
   } else {
	   fmt.Println("Email is duplicate. Use another email")
	   getUserDetails(db)
   }
}

func getUserDetails(db *sql.DB) {
	var firstName string
	var lastName string
	var email string
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("First Name: ", firstName)
	fmt.Println("Last Name: ", lastName)
	fmt.Println("Email: ", email)
	verification(firstName, lastName, email, db)
	checkDuplicateEmail(db, email)
}

func main() {
	db, err := donezodb.Connect()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()
	greetUsers()
	getUserDetails(db) 
}
