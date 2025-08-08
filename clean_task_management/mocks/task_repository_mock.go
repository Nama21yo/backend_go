package mocks

import (
	"github.com/stretchr/testify/mock"
	"task-manager/Domain"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) Create(t *domain.Task) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *TaskRepositoryMock) Update(t *domain.Task) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *TaskRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TaskRepositoryMock) GetByID(id string) (*domain.Task, error) {
	args := m.Called(id)
	var t *domain.Task
	if rf := args.Get(0); rf != nil {
		t = rf.(*domain.Task)
	}
	return t, args.Error(1)
}

func (m *TaskRepositoryMock) ListByOwner(ownerID string) ([]*domain.Task, error) {
	args := m.Called(ownerID)
	var ts []*domain.Task
	if rf := args.Get(0); rf != nil {
		ts = rf.([]*domain.Task)
	}
	return ts, args.Error(1)
}

func (m *TaskRepositoryMock) ListAll() ([]*domain.Task, error) {
	args := m.Called()
	var ts []*domain.Task
	if rf := args.Get(0); rf != nil {
		ts = rf.([]*domain.Task)
	}
	return ts, args.Error(1)
}
