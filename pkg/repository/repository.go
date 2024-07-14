package repository

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todolist.User) (int, error)
	GetUser(username, password string) (todolist.User, error)
}
type TodoList interface {
	Create(userId int, list todolist.TodoList) (int, error)
	GetAll(userId int) ([]todolist.TodoList, error)
	GetById(userId int, listId int) (todolist.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todolist.UpdateListInput) error
}
type TodoItem interface {
	Create(listId int, item todolist.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todolist.TodoItem, error)
	GetById(userId, itemId int) (todolist.TodoItem, error)
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
