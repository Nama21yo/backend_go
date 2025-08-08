package usecases_test

import (
	"errors"
	"testing"
	"time"
	"task-manager/Domain"
	"task-manager/Usecases"
	"task-manager/mocks"

	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	userRepo *mocks.UserRepositoryMock
	pwdSvc   *mocks.PasswordServiceMock
	jwtSvc   *mocks.JWTServiceMock
	uuc      usecases.UserUsecase
}

func (s *UserUsecaseTestSuite) SetupTest() {
	s.userRepo = &mocks.UserRepositoryMock{}
	s.pwdSvc = &mocks.PasswordServiceMock{}
	s.jwtSvc = &mocks.JWTServiceMock{}
	s.uuc = usecases.NewUserUsecase(s.userRepo, s.pwdSvc, s.jwtSvc, time.Hour)
}

func (s *UserUsecaseTestSuite) TestRegister_Success() {
	email := "a@a.com"
	password := "pass123"

	// userRepo.GetByEmail returns error -> not found
	s.userRepo.On("GetByEmail", email).Return(nil, errors.New("not found"))
	// password hash
	s.pwdSvc.On("Hash", password).Return("hashed", nil)
	// userRepo.Create should be called
	s.userRepo.On("Create", mockAnyUser()).Return(nil)

	u, err := s.uuc.Register(email, password, "Alice")
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), email, u.Email)
	s.userRepo.AssertExpectations(s.T())
	s.pwdSvc.AssertExpectations(s.T())
}

func (s *UserUsecaseTestSuite) TestRegister_EmailExists() {
	email := "a@a.com"
	password := "pass123"
	existing := &domain.User{ID: "u1", Email: email}
	s.userRepo.On("GetByEmail", email).Return(existing, nil)

	_, err := s.uuc.Register(email, password, "Alice")
	assert.Error(s.T(), err)
	s.userRepo.AssertExpectations(s.T())
}

func (s *UserUsecaseTestSuite) TestLogin_Success() {
	email := "a@b.com"
	password := "pass123"
	user := &domain.User{ID: "u1", Email: email, PasswordHash: "h"}
	s.userRepo.On("GetByEmail", email).Return(user, nil)
	s.pwdSvc.On("Compare", user.PasswordHash, password).Return(nil)
	s.jwtSvc.On("GenerateToken", user.ID, user.Role, mockAnyTTL()).Return("tok", nil)

	token, gotUser, err := s.uuc.Login(email, password)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "tok", token)
	assert.Equal(s.T(), user, gotUser)
}

func (s *UserUsecaseTestSuite) TestLogin_BadPassword() {
	email := "a@b.com"
	password := "bad"
	user := &domain.User{ID: "u1", Email: email, PasswordHash: "h"}
	s.userRepo.On("GetByEmail", email).Return(user, nil)
	s.pwdSvc.On("Compare", user.PasswordHash, password).Return(errors.New("mismatch"))

	_, _, err := s.uuc.Login(email, password)
	assert.Error(s.T(), err)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

// helper checkers for mocks: use testify/mock's Run or Matcher
// we create simple matchers below

// mockAnyUser returns a function matcher for *domain.User (used in On())
func mockAnyUser() func(interface{}) bool {
	return func(u interface{}) bool {
		_, ok := u.(*domain.User)
		return ok
	}
}

// mockAnyTTL returns a matcher for time.Duration arg (we accept any)
func mockAnyTTL() func(interface{}) bool {
	return func(d interface{}) bool {
		_, ok := d.(time.Duration)
		return ok
	}
}
