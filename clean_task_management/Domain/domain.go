package domain

import (
	"errors"
	"time"
)

type TaskStatus string

const (
	StatusTodo     TaskStatus = "todo"
	StatusDoing    TaskStatus = "doing"
	StatusComplete TaskStatus = "complete"
)

// Task is the core domain entity
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	OwnerID     string     `json:"owner_id"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Validate ensures required fields
func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.OwnerID == "" {
		return errors.New("owner_id is required")
	}
	return nil
}

// User is the core domain entity for users
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	DisplayName  string    `json:"display_name"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *User) ValidateForRegister() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.PasswordHash == "" {
		return errors.New("password is required")
	}
	return nil
}
