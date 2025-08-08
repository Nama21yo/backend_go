package usecases
package usecases

import (
	"errors"
	"time"
	"task-manager/Domain"
	"task-manager/Infrastructure"
	"task-manager/Repositories"
	"github.com/google/uuid"
)

// UserUsecase defines user related business operations
type UserUsecase interface {
	Register(email, password, displayName string) (*domain.User, error)
	Login(email, password string) (token string, user *domain.User, err error)
}

// userUsecase is concrete implementation
type userUsecase struct {
	userRepo        repositories.UserRepository
	passwordService infrastructure.PasswordService
	jwtService      infrastructure.JWTService
	tokenTTL        time.Duration
}

func NewUserUsecase(userRepo repositories.UserRepository, ps infrastructure.PasswordService, jwtSvc infrastructure.JWTService, ttl time.Duration) UserUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		passwordService: ps,
		jwtService:      jwtSvc,
		tokenTTL:        ttl,
	}
}

func (u *userUsecase) Register(email, password, displayName string) (*domain.User, error) {
	// check existing
	_, err := u.userRepo.GetByEmail(email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	hashed, err := u.passwordService.Hash(password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:           uuid.NewString(),
		Email:        email,
		PasswordHash: hashed,
		DisplayName:  displayName,
		Role:         "user",
	}
	if err := user.ValidateForRegister(); err != nil {
		return nil, err
	}
	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Login(email, password string) (string, *domain.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}
	if err := u.passwordService.Compare(user.PasswordHash, password); err != nil {
		return "", nil, errors.New("invalid credentials")
	}
	token, err := u.jwtService.GenerateToken(user.ID, user.Role, u.tokenTTL)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}
