# Task Management API with MongoDB - API Documentation

## Installation
# Task Management API with MongoDB - API Documentation

## Installation

### Prerequisites

* Go 1.18 or later
* MongoDB (local or remote instance)

### Install Dependencies

```bash
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
```

## MongoDB Configuration

Set the MongoDB URI in the code (`data/task_service.go`) or use an environment variable for better security:

```go
uri := "mongodb://localhost:27017" // Replace with your MongoDB URI
```

---

## Base URL

```
http://localhost:8080
```

---

## Endpoints

### GET /tasks

**Description**: Retrieve all tasks.

**Response:**

```json
[
  {
    "id": "64ec34f8af12b9a82237a940",
    "title": "Learn Go",
    "description": "Complete Go tutorials",
    "due_date": "2025-08-01T00:00:00Z",
    "status": "Pending"
  }
]
```

---

### GET /tasks/\:id

**Description**: Retrieve a task by its ID.

**Response:**

```json
{
  "id": "64ec34f8af12b9a82237a940",
  "title": "Learn Go",
  "description": "Complete Go tutorials",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "Pending"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

### POST /tasks

**Description**: Create a new task.

**Request Body:**

```json
{
  "title": "Build API",
  "description": "Create CRUD endpoints",
  "due_date": "2025-08-05T12:00:00Z",
  "status": "Pending"
}
```

**Response:**

```json
{
  "id": "64ec360b4a12b9a82237a945",
  "title": "Build API",
  "description": "Create CRUD endpoints",
  "due_date": "2025-08-05T12:00:00Z",
  "status": "Pending"
}
```

---

### PUT /tasks/\:id

**Description**: Update an existing task.

**Request Body:**

```json
{
  "title": "Build REST API",
  "description": "Implement all endpoints",
  "due_date": "2025-08-06T12:00:00Z",
  "status": "Completed"
}
```

**Response:**

```json
{
  "id": "64ec360b4a12b9a82237a945",
  "title": "Build REST API",
  "description": "Implement all endpoints",
  "due_date": "2025-08-06T12:00:00Z",
  "status": "Completed"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

### DELETE /tasks/\:id

**Description**: Delete a task by ID.

**Response:**

```json
{
  "message": "task deleted"
}
```

**Error:**
# Task Management API with MongoDB - API Documentation

## Installation

### Prerequisites

* Go 1.18 or later
* MongoDB (local or remote instance)

### Install Dependencies

```bash
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
```

## MongoDB Configuration

Set the MongoDB URI in the code (`data/task_service.go`) or use an environment variable for better security:

```go
uri := "mongodb://localhost:27017" // Replace with your MongoDB URI
```

---

## Base URL

```
http://localhost:8080
```

---

## Endpoints

### GET /tasks

**Description**: Retrieve all tasks.

**Response:**

```json
[
  {
    "id": "64ec34f8af12b9a82237a940",
    "title": "Learn Go",
    "description": "Complete Go tutorials",
    "due_date": "2025-08-01T00:00:00Z",
    "status": "Pending"
  }
]
```

---

### GET /tasks/\:id

**Description**: Retrieve a task by its ID.

**Response:**

```json
{
  "id": "64ec34f8af12b9a82237a940",
  "title": "Learn Go",
  "description": "Complete Go tutorials",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "Pending"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

### POST /tasks

**Description**: Create a new task.

**Request Body:**

```json
{
  "title": "Build API",
  "description": "Create CRUD endpoints",
  "due_date": "2025-08-05T12:00:00Z",
  "status": "Pending"
}
```

**Response:**

```json
{
  "id": "64ec360b4a12b9a82237a945",
  "title": "Build API",
  "description": "Create CRUD endpoints",
  "due_date": "2025-08-05T12:00:00Z",
  "status": "Pending"
}
```

---

### PUT /tasks/\:id

**Description**: Update an existing task.

**Request Body:**

```json
{
  "title": "Build REST API",
  "description": "Implement all endpoints",
  "due_date": "2025-08-06T12:00:00Z",
  "status": "Completed"
}
```

**Response:**

```json
{
  "id": "64ec360b4a12b9a82237a945",
  "title": "Build REST API",
  "description": "Implement all endpoints",
  "due_date": "2025-08-06T12:00:00Z",
  "status": "Completed"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

### DELETE /tasks/\:id

**Description**: Delete a task by ID.

**Response:**

```json
{
  "message": "task deleted"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

## Status Values

* `Pending`
* `In Progress`
* `Completed`

---

## Notes

* All dates should follow RFC3339 format.
* ID field returned is MongoDB ObjectID string.
* Responses are JSON formatted.
# Task Management API with MongoDB - API Documentation

## Installation

### Prerequisites

* Go 1.18 or later
* MongoDB (local or remote instance)

### Install Dependencies

```bash
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
```

## MongoDB Configuration

Set the MongoDB URI in the code (`data/task_service.go`) or use an environment variable for better security:

```go
uri := "mongodb://localhost:27017" // Replace with your MongoDB URI
```

---

## Base URL

```
http://localhost:8080
```

---

## Endpoints

### GET /tasks

**Description**: Retrieve all tasks.

**Response:**

```json
[
  {
    "id": "64ec34f8af12b9a82237a940",
    "title": "Learn Go",
    "description": "Complete Go tutorials",
    "due_date": "2025-08-01T00:00:00Z",
    "status": "Pending"
  }
]
```

---

### GET /tasks/\:id

**Description**: Retrieve a task by its ID.

**Response:**

```json
{
  "id": "64ec34f8af12b9a82237a940",
  "title": "Learn Go",
  "description": "Complete Go tutorials",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "Pending"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

### POST /tasks

**Description**: Create a new task.

**Request Body:**

```json
{
  "title": "Build API",
  "description": "Create CRUD endpoints",
  "due_date": "2025-08-05T12:00:00Z",
  "status": "Pending"
}
```

**Response:**

```json
{
  "id": "64ec360b4a12b9a82237a945",
  "title": "Build API",
  "description": "Create CRUD endpoints",
  "due_date": "2025-08-05T12:00:00Z",
  "status": "Pending"
}
```

---

### PUT /tasks/\:id

**Description**: Update an existing task.

**Request Body:**

```json
{
  "title": "Build REST API",
  "description": "Implement all endpoints",
  "due_date": "2025-08-06T12:00:00Z",
  "status": "Completed"
}
```

**Response:**

```json
{
  "id": "64ec360b4a12b9a82237a945",
  "title": "Build REST API",
  "description": "Implement all endpoints",
  "due_date": "2025-08-06T12:00:00Z",
  "status": "Completed"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

### DELETE /tasks/\:id

**Description**: Delete a task by ID.

**Response:**

```json
{
  "message": "task deleted"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

## Status Values

* `Pending`
* `In Progress`
* `Completed`

---

## Notes

* All dates should follow RFC3339 format.
* ID field returned is MongoDB ObjectID string.
* Responses are JSON formatted.

```json
{
  "error": "task not found"
}
```

---

## Status Values

* `Pending`
* `In Progress`
* `Completed`

---

## Notes

* All dates should follow RFC3339 format.
* ID field returned is MongoDB ObjectID string.
* Responses are JSON formatted.

### Prerequisites

* Go 1.18 or later
* MongoDB (local or remote instance)

### Install Dependencies

```bash
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
```

## MongoDB Configuration

Set the MongoDB URI in the code (`data/task_service.go`) or use an environment variable for better security:

```go
uri := "mongodb://localhost:27017" // Replace with your MongoDB URI
```

---

## Base URL

```
http://localhost:8080
```

---

## Endpoints

### GET /tasks

**Description**: Retrieve all tasks.

**Response:**

```json
[
  {
    "id": "64ec34f8af12b9a82237a940",
    "title": "Learn Go",
    "description": "Complete Go tutorials",
    "due_date": "2025-08-01T00:00:00Z",
    "status": "Pending"
  }
]
```

---

### GET /tasks/\:id

**Description**: Retrieve a task by its ID.

**Response:**

```json
{
  "id": "64ec34f8af12b9a82237a940",
  "title": "Learn Go",
  "description": "Complete Go tutorials",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "Pending"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

### POST /tasks

**Description**: Create a new task.

**Request Body:**

```json
{
  "title": "Build API",
  "description": "Create CRUD endpoints",
  "due_date": "2025-08-05T12:00:00Z",
  "status": "Pending"
}
```

**Response:**

```json
{
  "id": "64ec360b4a12b9a82237a945",
  "title": "Build API",
  "description": "Create CRUD endpoints",
  "due_date": "2025-08-05T12:00:00Z",
  "status": "Pending"
}
```

---

### PUT /tasks/\:id

**Description**: Update an existing task.

**Request Body:**

```json
{
  "title": "Build REST API",
  "description": "Implement all endpoints",
  "due_date": "2025-08-06T12:00:00Z",
  "status": "Completed"
}
```

**Response:**

```json
{
  "id": "64ec360b4a12b9a82237a945",
  "title": "Build REST API",
  "description": "Implement all endpoints",
  "due_date": "2025-08-06T12:00:00Z",
  "status": "Completed"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

### DELETE /tasks/\:id

**Description**: Delete a task by ID.

**Response:**

```json
{
  "message": "task deleted"
}
```

**Error:**

```json
{
  "error": "task not found"
}
```

---

## Status Values

* `Pending`
* `In Progress`
* `Completed`

---

## Notes

* All dates should follow RFC3339 format.
* ID field returned is MongoDB ObjectID string.
* Responses are JSON formatted.
