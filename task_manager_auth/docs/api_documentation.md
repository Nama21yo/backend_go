# API Documentation - Task Management API with JWT Auth

## 🔐 Authentication & Authorization

### Register a User

`POST /register`

```json
{
  "username": "john",
  "password": "secret"
}
```

**Response:**

```json
{
  "message": "User registered successfully",
  "role": "admin" // if first user
}
```

### Login a User

`POST /login`

```json
{
  "username": "john",
  "password": "secret"
}
```

**Response:**

```json
{
  "token": "<JWT_TOKEN>"
}
```

Use the token in the `Authorization` header for all protected endpoints:

```
Authorization: Bearer <JWT_TOKEN>
```

---

## 🧑‍🎓 User Management

### Promote to Admin

`POST /promote/:username`

* Requires `admin` role
* Promotes user to admin

**Response:**

```json
{
  "message": "User promoted to admin"
}
```

---

## 📋 Tasks Endpoints

### Create Task

`POST /tasks`

```json
{
  "title": "Finish API",
  "description": "Complete by Sunday",
  "due_date": "2025-07-25T00:00:00Z",
  "status": "pending"
}
```

**Authorization:** Admin only

### Get All Tasks

`GET /tasks`

* Public (authenticated users)

### Get Task by ID

`GET /tasks/:id`

* Public (authenticated users)

### Update Task

`PUT /tasks/:id`

* Admin only

### Delete Task

`DELETE /tasks/:id`

* Admin only

---

## ❗ Errors

```json
{
  "error": "Unauthorized access"
}
```

```json
{
  "error": "Invalid credentials"
}
```

```json
{
  "error": "Task not found"
}
```

---

## ⚙️ Environment Variables

Create a `.env` file:

```
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=supersecretkey
PORT=8080
```

---

