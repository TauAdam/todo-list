package service

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/repository"
)

type Auth interface {
	CreateUser(user todolist.User) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Auth
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Authorization),
	}
}
