package service

import "github.com/TauAdam/todo-list/pkg/repository"

type Auth interface {
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
	return &Service{}
}
