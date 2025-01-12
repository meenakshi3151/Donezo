package main

import (
	"fmt"
	"strings"
	"ToDoList/donezodb"
	"database/sql"
)

var countOfDetails int = 0
var countOfLogin int = 0
var tasksExisting [][]string

func greetUsers() {
	fmt.Println("Welcome to one stop solution of your tasks")
}

func verifyEmail(email string) bool {
	var specialChar = []string{"!", "#", "$", "%", "^", "&", "*", "(", ")", "_", "-", "+", "=", "{", "}", "[", "]", ":", ";", "'", "<", ">", ",",  "?", "/", "|", "`", "~"}
	for _, char := range specialChar {
		if strings.Contains(email, char) {
			return false
		}
	}
	return strings.Contains(email, "@gmail.com")
}

func verification(firstName string, lastName string, email string, password string, db *sql.DB) {
	if len(firstName) > 2 && len(lastName) > 2 && verifyEmail(email) && len(password) >= 8 {
		fmt.Println("You are verified")
		checkDuplicateEmail(db, email, firstName, lastName, password)
	} else {
		fmt.Println("Please enter a valid name and email")
		countOfDetails++
		if countOfDetails == 2 {
			fmt.Println("You have reached the limit of entering the details")
			return 
		}
		getUserDetails(db)
	}
}

func checkDuplicateEmail(db *sql.DB,email string, firstName string, lastName string, password string) {
   if(donezodb.CheckEmail(db, email)) {
	   donezodb.InsertUser(db, firstName, lastName, email, password)
   } else {
	   fmt.Println("Email is duplicate. Use another email")
	   getUserDetails(db)
   }
}

func getUserDetails(db *sql.DB) {
	var firstName string
	var lastName string
	var email string
	var password string
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print(("Enter your password: "))
	fmt.Scan(&password)
	fmt.Println("First Name: ", firstName)
	fmt.Println("Last Name: ", lastName)
	fmt.Println("Email: ", email)

	verification(firstName, lastName, email, password, db)
}

func askForExistingUser(db *sql.DB) bool {
	var choice string
	fmt.Print("Are you an existing user? (y/n): ")
	fmt.Scan(&choice)
	return choice == "y"
}

func loginUser(db *sql.DB) string {
	var email string
	var password string
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter your password: ")
	fmt.Scan(&password)
	for countOfLogin < 2 {
		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Print("Enter your password: ")
		fmt.Scan(&password)
		if donezodb.LoginUser(db, email, password) {
			fmt.Println("Login successful")
			return email
		}
		fmt.Println("Login failed")
		countOfLogin++
	}
	fmt.Println("You have reached the limit of entering the login details")
	return ""
}

func showTasks(db *sql.DB, email string) {
	donezodb.GetTasksAndStatus(db, &tasksExisting, email)
	if len(tasksExisting) == 0 {
		fmt.Println("No tasks to show")
		return
	}
	fmt.Println("Your tasks along with the status are:")
	for _, task := range tasksExisting {
		fmt.Printf("Task Id: %s, Task Name: %s, Status: %s\n", task[0], task[1], task[2])
	}
}

func askForUpdationOfTasks(db *sql.DB, email string ) {
   var existingTask = len (tasksExisting)
   var noOfTasksUpdate int 
   fmt.Println("Enter the number of tasks you want to update out of: ", existingTask)
   fmt.Scan(&noOfTasksUpdate)
   for i := 0; i < noOfTasksUpdate; i++ {
	   fmt.Print("Enter the task id you want to update: ")
	   var taskId int
	   fmt.Scan(&taskId)
	   fmt.Print("Enter the task status: ")
	   var taskStatus string
	   fmt.Scan(&taskStatus)
	   _, err := db.Exec("UPDATE task SET Status = ? WHERE Id = ? AND Email = ?", taskStatus, taskId, email)
	   if err != nil {
		   fmt.Println("Error updating the task:", err)
	   }
   }
}

func AddNewTasks(db *sql.DB, email string) {
	var noOfTasks int
	fmt.Print("Enter the number of tasks you want to add: ")
	fmt.Scan(&noOfTasks)
	var tasks []string
	for i := 0; i < noOfTasks; i++ {
		fmt.Printf("Enter the task %d: ", i+1)
		var taskName string
		fmt.Scanln(&taskName) 
		tasks = append(tasks, taskName)
	}
	donezodb.InsertTask(db, tasks, email)
	fmt.Println("Tasks added successfully")
}

func askChoiceOfUser(db *sql.DB, email string) {
	fmt.Println("Do you want to add new tasks or update the existing tasks?")
	fmt.Println("1. Add new tasks")
	fmt.Println("2. Update the existing tasks")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		AddNewTasks(db, email)
	case 2:
		askForUpdationOfTasks(db, email)
	default:
		fmt.Println("Invalid choice")
		return
	}
}

func exit() bool {
    fmt.Println("Do you want to close the session? (y/n): ")
	var choice string
	fmt.Scan(&choice)
	if choice == "y" {
	   fmt.Println("I hope you liked our service")
	   return true
	}
	return false 
}

func main() {
	db, err := donezodb.Connect()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()
	greetUsers()
	if askForExistingUser(db) {
		fmt.Println("Welcome back")
	} else {
		getUserDetails(db)
	} 
	var response_email = loginUser(db)
	if response_email == "" {
		fmt.Println("Unable to log in. Exiting...")
		return
	}	
	fmt.Println("---------------------------------")
	showTasks(db, response_email)
	fmt.Println("---------------------------------")
	for {
		askChoiceOfUser(db, response_email)
		showTasks(db, response_email)
		if exit() {
			break
		}		
	}
}
