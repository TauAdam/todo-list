package service

import (
	"crypto/sha1"
	"fmt"
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user todolist.User) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

const (
	salt         = "asdasd"
	tokenTTL     = 12 * time.Hour
	mySigningKey = "AllYourBase" // Changed to a string constant
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func hashPassword(password string) string {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		logrus.Error(err)
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, hashPassword(password))
	if err != nil {
		return "", err
	}
	claims := &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(mySigningKey))
}
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(mySigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}
