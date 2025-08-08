package domain_test

import (
	domain "task-manager/Domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskValidate(t *testing.T) {
	t1 := &domain.Task{
		ID:      "1",
		Title:   "",
		OwnerID: "user1",
	}
	err := t1.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title is required")

	t2 := &domain.Task{
		ID:      "2",
		Title:   "ok",
		OwnerID: "",
	}
	err = t2.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "owner_id is required")

	t3 := &domain.Task{
		ID:      "3",
		Title:   "ok",
		OwnerID: "user1",
	}
	assert.NoError(t, t3.Validate())
}

func TestUserValidateForRegister(t *testing.T) {
	u := &domain.User{
		ID:           "u1",
		Email:        "",
		PasswordHash: "",
	}
	err := u.ValidateForRegister()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email is required")

	u.Email = "a@b.com"
	u.PasswordHash = "hash"
	assert.NoError(t, u.ValidateForRegister())
}
