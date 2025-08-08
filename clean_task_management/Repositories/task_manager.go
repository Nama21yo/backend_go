package repositories

import (
	"errors"
	"sync"
	"time"
	"task-manager/Domain"
)

// TaskRepository defines behavior for task persistence
type TaskRepository interface {
	Create(task *domain.Task) error
	Update(task *domain.Task) error
	Delete(id string) error
	GetByID(id string) (*domain.Task, error)
	ListByOwner(ownerID string) ([]*domain.Task, error)
	ListAll() ([]*domain.Task, error)
}

// InMemoryTaskRepository is a simple thread-safe in-memory storage for tasks
type InMemoryTaskRepository struct {
	mu    sync.RWMutex
	store map[string]*domain.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		store: make(map[string]*domain.Task),
	}
}

func (r *InMemoryTaskRepository) Create(task *domain.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now().UTC()
	task.CreatedAt = now
	task.UpdatedAt = now
	r.store[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) Update(task *domain.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	existing, ok := r.store[task.ID]
	if !ok {
		return errors.New("task not found")
	}
	task.CreatedAt = existing.CreatedAt
	task.UpdatedAt = time.Now().UTC()
	r.store[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.store[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(r.store, id)
	return nil
}

func (r *InMemoryTaskRepository) GetByID(id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.store[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return t, nil
}

func (r *InMemoryTaskRepository) ListByOwner(ownerID string) ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	res := make([]*domain.Task, 0)
	for _, t := range r.store {
		if t.OwnerID == ownerID {
			res = append(res, t)
		}
	}
	return res, nil
}

func (r *InMemoryTaskRepository) ListAll() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	res := make([]*domain.Task, 0, len(r.store))
	for _, t := range r.store {
		res = append(res, t)
	}
	return res, nil
}
