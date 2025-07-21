# Task Management API Documentation

## Base URL

```
[http://localhost:8080](http://localhost:8080)
```

## Endpoints

---

### GET /tasks

Get a list of all tasks.

**Request:**

```http
GET /tasks HTTP/1.1
Host: localhost:8080
```

**Response:**

```json
[
  {
    "id": 1,
    "title": "Example Task",
    "description": "This is a sample task",
    "due_date": "2025-07-22T12:00:00Z",
    "status": "Pending"
  }
]
```

---

### GET /tasks/\:id

Get the details of a specific task by ID.

**Request:**

```http
GET /tasks/1 HTTP/1.1
Host: localhost:8080
```

**Response:**

```json
{
  "id": 1,
  "title": "Example Task",
  "description": "This is a sample task",
  "due_date": "2025-07-22T12:00:00Z",
  "status": "Pending"
}
```

**Error Response (if not found):**

```json
{
  "error": "task not found"
}
```

---

### POST /tasks

Create a new task.

**Request:**

```http
POST /tasks HTTP/1.1
Content-Type: application/json

{
  "title": "Write Docs",
  "description": "Complete the documentation",
  "due_date": "2025-07-25T15:00:00Z",
  "status": "Pending"
}
```

**Response:**

HTTP/1.1 201 Created

```json
{
  "id": 2,
  "title": "Write Docs",
  "description": "Complete the documentation",
  "due_date": "2025-07-25T15:00:00Z",
  "status": "Pending"
}
```

**Error Response (invalid input):**

```json
{
  "error": "invalid input"
}
```

---

### PUT /tasks/\:id

Update an existing task by ID.

**Request:**

```http
PUT /tasks/2 HTTP/1.1
Content-Type: application/json

{
  "title": "Update Task",
  "description": "Updated task details",
  "due_date": "2025-07-27T12:00:00Z",
  "status": "Completed"
}
```

**Response:**

HTTP/1.1 200 OK

```json
{
  "id": 2,
  "title": "Update Task",
  "description": "Updated task details",
  "due_date": "2025-07-27T12:00:00Z",
  "status": "Completed"
}
```

**Error Response (task not found):**

```json
{
  "error": "task not found"
}
```

---

### DELETE /tasks/\:id

Delete a task by ID.

**Request:**

```http
DELETE /tasks/2 HTTP/1.1
Host: localhost:8080
```

**Response:**

HTTP/1.1 204 No Content

**Error Response (if not found):**

```json
{
  "error": "task not found"
}
```

---

## Notes

- All date formats are expected in [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) format.
- Status values can be: `"Pending"`, `"In Progress"`, `"Completed"`.
- All endpoints return JSON responses.
