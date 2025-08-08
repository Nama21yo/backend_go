package mocks

import (
	"github.com/stretchr/testify/mock"
)

type PasswordServiceMock struct {
	mock.Mock
}

func (m *PasswordServiceMock) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *PasswordServiceMock) Compare(hashed, plain string) error {
	args := m.Called(hashed, plain)
	return args.Error(0)
}
