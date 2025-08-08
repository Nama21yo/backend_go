package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordService abstracts password hashing functions
type PasswordService interface {
	Hash(password string) (string, error)
	Compare(hashed, plain string) error
}

type BcryptPasswordService struct {
	cost int
}

func NewBcryptPasswordService(cost int) *BcryptPasswordService {
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}
	return &BcryptPasswordService{cost: cost}
}

func (s *BcryptPasswordService) Hash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s *BcryptPasswordService) Compare(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
