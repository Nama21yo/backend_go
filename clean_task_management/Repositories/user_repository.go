package repositories

import (
	"errors"
	"sync"
	"time"
	"task-manager/Domain"
)

// UserRepository defines user persistence behavior
type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	store map[string]*domain.User
	// email index
	emailIndex map[string]string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		store:      make(map[string]*domain.User),
		emailIndex: make(map[string]string),
	}
}

func (r *InMemoryUserRepository) Create(user *domain.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now
	r.store[user.ID] = user
	r.emailIndex[user.Email] = user.ID
	return nil
}

func (r *InMemoryUserRepository) GetByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.store[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *InMemoryUserRepository) GetByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.emailIndex[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	u, ok := r.store[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}
