package infrastructure
package infrastructure

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTService abstracts token generation and validation
type JWTService interface {
	GenerateToken(userID string, role string, ttl time.Duration) (string, error)
	ValidateToken(tokenStr string) (*TokenClaims, error)
}

type jwtService struct {
	secret []byte
}

type TokenClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTService(secret string) JWTService {
	return &jwtService{
		secret: []byte(secret),
	}
}

func (s *jwtService) GenerateToken(userID string, role string, ttl time.Duration) (string, error) {
	claims := &TokenClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(s.secret)
}

func (s *jwtService) ValidateToken(tokenStr string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
