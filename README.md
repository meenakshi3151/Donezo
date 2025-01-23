# To-Do List Application (Terminal-Based)

This is a terminal-based To-Do List application built with Go and SQLite. The application provides users with a simple way to manage tasks, including adding, updating, and displaying tasks. It also supports user registration and login functionality to maintain personalized task lists.

## Features

- **User Management**:
  - New users can register by providing their first name, last name, email, and password.
  - Existing users can log in using their email and password.
  - Duplicate email checks are enforced during registration.

- **Task Management**:
  - View tasks and their statuses.
  - Add new tasks.
  - Update the status of existing tasks.

- **Database Integration**:
  - SQLite is used as the backend database to store user details and tasks.

- **Error Handling**:
  - The application handles common errors like invalid email, duplicate entries, and database errors gracefully.

## Technologies Used

- **Go (Golang)**: For building the application backend and managing logic.
- **SQLite**: As the database for persisting user and task data.
- **SQL**: For executing database queries to interact with the SQLite database.

## Prerequisites

- Go (Golang) installed on your system. [Download Go](https://go.dev/dl/)
- SQLite installed for database management. [Download SQLite](https://www.sqlite.org/download.html)

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/your-repo/todo-terminal-app.git
cd todo-terminal-app
