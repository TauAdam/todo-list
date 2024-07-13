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
	GetById(userId int, id int) (todolist.TodoList, error)
}
type TodoItem interface {
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
	}
}
