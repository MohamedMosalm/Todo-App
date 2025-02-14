# To-Do List Application

This is a To-Do List application built with Go, Gin, and GORM. It provides user authentication and task management functionalities.

## Table of Contents

- [Setup Instructions](#setup-instructions)
- [API Documentation](#api-documentation)
  - [Authentication](#authentication)
  - [Tasks](#tasks)
- [Usage Examples](#usage-examples)
- [Running Tests](#running-tests)

## Setup Instructions

1. **Clone the repository:**

   ```sh
   git clone https://github.com/MohamedMosalm/Todo-App.git
   cd Todo-App/Mohamed-Mosalm-todo-app
   ```

2. **Create a `.env` file:**

   ```sh
   touch .env
   ```

   Add the following environment variables to the `.env` file:

   ```env
   HTTP_PORT=9090
   JWT_SECRET=your_jwt_secret
   DB_HOST=localhost
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_PORT=5432
   DB_SSLMODE=disable
   ```

3. **Run the application using Docker:**

   ```sh
   docker-compose up --build
   ```

4. **Run the application locally:**

   Ensure you have Go installed and run the following commands:

   ```sh
   go mod download
   go run main.go
   ```

## Running Tests

To run tests, use the following command:

```sh
go test ./...
```

## API Documentation

### Authentication

- **Register**

  ```http
  POST /api/auth/register
  ```

  Request Body:

  ```json
  {
    "first_name": "Mohamed",
    "last_name": "Mosalm",
    "email": "MohamedMosalm@example.com",
    "phone": "1234567890",
    "password": "password123"
  }
  ```

  Response:

  ```json
  {
    "status": "success",
    "message": "User registered successfully"
  }
  ```

- **Login**

  ```http
  POST /api/auth/login
  ```

  Request Body:

  ```json
  {
    "email": "MoahmedMosalm@example.com",
    "password": "password123"
  }
  ```

  Response:

  ```json
  {
    "status": "success",
    "message": "Login successful",
    "data": {
      "access_token": "your_jwt_token",
      "token_type": "Bearer",
      "user": {
        "id": "user_id",
        "email": "MohamedMosalm@example.com",
        "first_name": "Mohamed",
        "last_name": "Mosalm"
      }
    }
  }
  ```

### Tasks

- **Create Task**

  ```http
  POST /api/tasks
  ```

  Request Body:

  ```json
  {
    "title": "New Task",
    "description": "Task description"
  }
  ```

  Response:

  ```json
  {
    "status": "success",
    "message": "Task created successfully",
    "data": {
      "id": "task_id",
      "title": "New Task",
      "description": "Task description",
      "status": false,
      "user_id": "user_id",
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
  }
  ```

- **Get Tasks**

  ```http
  GET /api/tasks
  ```

  Response:

  ```json
  {
    "status": "success",
    "message": "Tasks retrieved successfully",
    "data": [
      {
        "id": "task_id",
        "title": "Task Title",
        "description": "Task description",
        "status": false,
        "user_id": "user_id",
        "created_at": "timestamp",
        "updated_at": "timestamp"
      }
    ]
  }
  ```

- **Update Task**

  ```http
  PUT /api/tasks/:id
  ```

  Request Body:

  ```json
  {
    "title": "Updated Task Title",
    "description": "Updated Task description",
    "status": true
  }
  ```

  Response:

  ```json
  {
    "status": "success",
    "message": "Task updated successfully",
    "data": {
      "id": "task_id",
      "title": "Updated Task Title",
      "description": "Updated Task description",
      "status": true,
      "user_id": "user_id",
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
  }
  ```

- **Delete Task**

  ```http
  DELETE /api/tasks/:id
  ```

  Response:

  ```json
  {
    "status": "success",
    "message": "Task deleted successfully"
  }
  ```

## Usage Examples

### Register a New User

```sh
curl -X POST http://localhost:9090/api/auth/register \
    -H "Content-Type: application/json" \
    -d '{
        "first_name": "Mohamed",
        "last_name": "Mosalm",
        "email": "MohamedMosalm@example.com",
        "phone": "1234567890",
        "password": "password123"
    }'
```

### Login

```sh
curl -X POST http://localhost:9090/api/auth/login \
    -H "Content-Type: application/json" \
    -d '{
        "email": "MohamedMosalm@example.com",
        "password": "password123"
    }'
```

### Create a Task

```sh
curl -X POST http://localhost:9090/api/tasks \
    -H "Authorization: Bearer your_jwt_token" \
    -H "Content-Type: application/json" \
    -d '{
        "title": "New Task",
        "description": "Task description"
    }'
```

### Get All Tasks

```sh
curl -X GET http://localhost:9090/api/tasks \
    -H "Authorization: Bearer your_jwt_token"
```

### Update a Task

```sh
curl -X PUT http://localhost:9090/api/tasks/task_id \
    -H "Authorization: Bearer your_jwt_token" \
    -H "Content-Type: application/json" \
    -d '{
        "title": "Updated Task Title",
        "description": "Updated Task description",
        "status": true
    }'
```

### Delete a Task

```sh
curl -X DELETE http://localhost:9090/api/tasks/task_id \
    -H "Authorization: Bearer your_jwt_token"
```
