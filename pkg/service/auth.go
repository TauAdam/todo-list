package service

import (
	"crypto/sha1"
	"fmt"
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/repository"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user todolist.User) (int, error) {
	user.Password = s.hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

const salt = "asdasd"

func (s *AuthService) hashPassword(password string) string {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		logrus.Error(err)
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
