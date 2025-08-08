package mocks

import (
	"github.com/stretchr/testify/mock"
	"task-manager/Domain"
)

// UserRepositoryMock is a testify/mock implementation of repositories.UserRepository
type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(u *domain.User) error {
	args := m.Called(u)
	if args.Get(0) == nil {
		return args.Error(0)
	}
	return args.Error(0)
}

func (m *UserRepositoryMock) GetByID(id string) (*domain.User, error) {
	args := m.Called(id)
	var u *domain.User
	if rf := args.Get(0); rf != nil {
		u = rf.(*domain.User)
	}
	return u, args.Error(1)
}

func (m *UserRepositoryMock) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	var u *domain.User
	if rf := args.Get(0); rf != nil {
		u = rf.(*domain.User)
	}
	return u, args.Error(1)
}
