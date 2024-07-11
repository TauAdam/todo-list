package repository

import "github.com/jmoiron/sqlx"

type Auth interface {
}
type TodoList interface {
}
type TodoItem interface {
}
type Repository struct {
	Auth
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
