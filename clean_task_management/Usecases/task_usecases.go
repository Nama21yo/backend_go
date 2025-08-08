package usecases

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// TaskUsecase defines business operations for tasks
type TaskUsecase interface {
	CreateTask(title, description, ownerID string) (*domain.Task, error)
	UpdateTask(task *domain.Task, requesterID string) (*domain.Task, error)
	GetTaskByID(id string, requesterID string) (*domain.Task, error)
	ListTasks(ownerID string) ([]*domain.Task, error)
	DeleteTask(id string, requesterID string) error
	ListAllTasks() ([]*domain.Task, error) // admin helper
}

type taskUsecase struct {
	taskRepo repositories.TaskRepository
}

func NewTaskUsecase(tr repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: tr}
}

func (t *taskUsecase) CreateTask(title, description, ownerID string) (*domain.Task, error) {
	task := &domain.Task{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		OwnerID:     ownerID,
		Status:      domain.StatusTodo,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	if err := task.Validate(); err != nil {
		return nil, err
	}
	if err := t.taskRepo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskUsecase) UpdateTask(task *domain.Task, requesterID string) (*domain.Task, error) {
	// only owner can update in this simple model
	existing, err := t.taskRepo.GetByID(task.ID)
	if err != nil {
		return nil, err
	}
	if existing.OwnerID != requesterID {
		return nil, errors.New("unauthorized")
	}
	task.UpdatedAt = time.Now().UTC()
	if err := task.Validate(); err != nil {
		return nil, err
	}
	if err := t.taskRepo.Update(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskUsecase) GetTaskByID(id string, requesterID string) (*domain.Task, error) {
	task, err := t.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if task.OwnerID != requesterID {
		return nil, errors.New("unauthorized")
	}
	return task, nil
}

func (t *taskUsecase) ListTasks(ownerID string) ([]*domain.Task, error) {
	return t.taskRepo.ListByOwner(ownerID)
}

func (t *taskUsecase) DeleteTask(id string, requesterID string) error {
	task, err := t.taskRepo.GetByID(id)
	if err != nil {
		return err
	}
	if task.OwnerID != requesterID {
		return errors.New("unauthorized")
	}
	return t.taskRepo.Delete(id)
}

func (t *taskUsecase) ListAllTasks() ([]*domain.Task, error) {
	return t.taskRepo.ListAll()
}
