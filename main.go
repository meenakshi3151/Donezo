package main
import (
	"fmt"
)

var countOfDetails int = 0

func greetUsers() {
	fmt.Println("Welcome to one stop solution of your tasks")
}

func verification(firstName string, lastName string, email string) {
	if len(firstName) > 2 && len(lastName) > 2 && len(email) > 2 {
		fmt.Println("You are verified")
	} else {
		fmt.Println("Please enter a valid name with at least 2 characters.")
		countOfDetails++
		if countOfDetails == 2 {
			fmt.Println("You have reached the limit of entering the details")
			return 
		}
		getUserDetails()
	}
}

func getUserDetails() {
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
	verification(firstName, lastName, email)
}


func main() {
	greetUsers();
	getUserDetails();

   
}