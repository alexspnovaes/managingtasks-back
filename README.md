# ManagingTasks - Back-End (Go + SQLite)

This back-end application provides a RESTful API for managing tasks, built with Go (Golang) and SQLite. It supports creating, reading, updating, and deleting tasks (CRUD operations).

## Table of Contents

1. [Features](#features)
2. [Tech Stack](#tech-stack)
3. [Getting Started](#getting-started)
5. [Running the Application](#running-the-application)
6. [API Endpoints](#api-endpoints)
7. [Database Setup](#database-setup)
8. [License](#license)

---

## Features

- **CRUD for Tasks**:
  - Create a new task (title, completed=false by default).
  - Retrieve all tasks.
  - Update a task (change title or completion status).
  - Delete a task by ID.
- **CORS Middleware** for cross-origin requests from the Angular front-end.
- **SQLite** used as the database.

---

## Tech Stack

- **Go** (Golang)  
- **SQLite** (local database)  
- **Gorilla Mux** or the built-in net/http for routing  
- **Go Modules** for dependency management

## Getting Started

1. **Install Go** (>= 1.19 recommended).
2. **Clone** this repository:
```bash
git clone [managingtasks](https://github.com/alexspnovaes/managingtasks-back.git)
```    
Navigate to the backend folder:
   ```bash
cd managingtasks/backend
```
## Running the Application
  Install dependencies:
   ```bash
go mod tidy
```
Run the Go server:
```bash
go run main.go
```
The server will start on http://localhost:8080.
By default, a SQLite database file is created in the project folder if it doesn’t exist.

## API Endpoints
Method	Endpoint	     Description	
GET	    /tasks	       Get all tasks	
POST	  /tasks	       Create a new task	{"title": "Learn Go"}
PUT	    /tasks/{id}	   Update a task (title/status)	{"title": "Learn Go + Angular", "completed":true}
DELETE	/tasks/{id}	   Delete a task by ID	-

## Database Setup
The application uses SQLite by default.
No special setup is needed; the code automatically creates a .db file (or uses an existing one) if configured that way.
Ensure the code in infrastructure/database or wherever you manage DB is correct.
## License
This project is open-sourced under the MIT License. Feel free to modify and distribute it.
