# ManagingTasks - Back-End (Go + SQLite)

This back-end application provides a RESTful API for managing tasks, built with Go (Golang) and SQLite. It supports creating, reading, updating, and deleting tasks (CRUD operations).

## Table of Contents

1. [Features](#features)
2. [Tech Stack](#tech-stack)
3. [Project Structure](#project-structure)
4. [Getting Started](#getting-started)
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

---

## Project Structure

