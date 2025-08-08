package mocks

import (
	infrastructure "task-manager/Infrastructure"
	"time"

	"github.com/stretchr/testify/mock"
)

type JWTServiceMock struct {
	mock.Mock
}

func (m *JWTServiceMock) GenerateToken(userID string, role string, ttl time.Duration) (string, error) {
	args := m.Called(userID, role, ttl)
	return args.String(0), args.Error(1)
}

func (m *JWTServiceMock) ValidateToken(tokenStr string) (*infrastructure.TokenClaims, error) {
	args := m.Called(tokenStr)
	var claims *infrastructure.TokenClaims
	if rf := args.Get(0); rf != nil {
		claims = rf.(*infrastructure.TokenClaims)
	}
	return claims, args.Error(1)
}
