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
	GetById(userId, listId int) (todolist.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todolist.UpdateListInput) error
}
type TodoItem interface {
	Create(userId, listId int, item todolist.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]todolist.TodoItem, error)
	GetById(userId, itemId int) (todolist.TodoItem, error)
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
		TodoItem: NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
