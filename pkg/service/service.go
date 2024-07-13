package service

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/repository"
)

type Auth interface {
	CreateUser(user todolist.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}
type TodoList interface {
	Create(userId int, list todolist.TodoList) (int, error)
	GetAll(userId int) ([]todolist.TodoList, error)
	GetById(userId, id int) (todolist.TodoList, error)
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
		Auth:     NewAuthService(repo.Authorization),
		TodoList: NewTodoListService(repo.TodoList),
	}
}
