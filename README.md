# Task Tracker Application

This is a Task Tracker application built using Go, following the principles of Hexagonal Architecture. 
The application allows users to create, update, delete, and fetch all tasks.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
- [Running Locally](#running-locally)
- [API Documentation](#api-documentation)
- [Future Scope](#future-scope)

## Features

- Create new tasks
- Update existing tasks
- Delete tasks
- Fetch all tasks
- User management (create and login users)

## Technologies

- Go (Golang)
- Gin (HTTP web framework)
- GORM (ORM for database interactions)
- PostgreSQL (or any relational database supported by GORM)

## Getting Started

### Prerequisites

Make sure you have the following installed:

- Go (version 1.16 or higher)
- A relational database (e.g., PostgreSQL) set up and running
- Any necessary environment variables configured (e.g., database connection strings)

### Running Locally

To run the application locally, follow these steps:

1. Clone the repository: git clone https://github.com/singhamritpalAP/Task-Tracker.git
2. Navigate to the application directory: cd cmd/task-tracker
3. Run the application: go run .


The application should now be running locally on the specified port (defined in `constants.ApplicationPort`).

## API Documentation

### Base URL
http://localhost:8080

### Endpoints

#### 1. Create Task

**POST** `/tasktracker/tasks`

**Request Body**:

{
"title": "Task Tracker CRUD",
"description": "Create crud for task Tracker",
"status": "COMPLETED"
}


**Response**:
- **201 Created**: Task created successfully.
- **400 Bad Request**: Missing required fields.

#### 2. Fetch All Tasks

**GET** `/tasktracker/tasks`

**Response**:
- **200 OK**: Returns a list of tasks.

{
"tasks": [
{
"TaskId": 1,
"Title": "Task Tracker CRUD",
"Description": "Create crud for task Tracker",
"Status": "PENDING",
"CreatedAt": "2025-01-27T00:24:01.676295+05:30",
"UpdatedAt": "2025-01-27T00:24:31.137888+05:30"
},
{
"TaskId": 2,
"Title": "Use JWT",
"Description": "Use JWT for validation",
"Status": "COMPLETED",
"CreatedAt": "2025-01-27T00:27:29.745479+05:30",
"UpdatedAt": "2025-01-27T00:27:29.745479+05:30"
}
]
}

- **204 No Content**: Returns nothing.


#### 3. Update Task

**PATCH** `/tasktracker/tasks`

**Request Body**:

{
"task_id": 1,
"title": "Updated Task Title",
"description": "Updated Description"
}

**Response**:
- **200 OK**: Task updated successfully.
- **400 Bad Request**: Missing task ID or required fields.

#### 4. Delete Task

**DELETE** `/tasktracker/tasks?taskId=1`

**Response**:
- **200 OK**: Task deleted successfully.
- **400 Bad Request**: Missing task ID.

#### 5. Create User

**POST** `/tasktracker/user`

**Request Body**:

{
"username": "newuser",
"password": "securepassword"
}

**Response**:
- **201 Created**: User created successfully.
- **400 Bad Request**: Missing required fields.

#### 6. Login User

**POST** `/tasktracker/login`

**Request Body**:

{
"username": "existinguser",
"password": "securepassword"
}

**Response**:
- **200 OK**: User validated successfully.

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzgxNzgxODksImlkIjoxfQ.TbMCt2ME2VvUQr_zOrqWJkSLxWtowR7S4QiMBZPeZgo"
}
- **401 Unauthorized**: Invalid credentials.


## Future Scope

The following features are planned for future development:

1. **Store Hashed Passwords:** Implement password hashing to securely store user passwords in the database.
2. **Validate Username and Password:** Enhance user validation logic to ensure usernames are unique and passwords meet security standards.
3. **Create Dockerfile and Compose:** Develop a Dockerfile and Docker Compose configuration for containerized deployment of the application.
4. **User-Specific Tasks:** Implement functionality so that users can only perform CRUD operations on tasks they created.
5. **Set Database Column Length:** Define appropriate lengths for database columns to enforce data integrity.
6. **Fetch Configs from Config Map:** Integrate configuration management to fetch settings from a config map for better environment handling.
7. **Add Validation for User Input:** Implement input validation for user details during creation and authentication processes.