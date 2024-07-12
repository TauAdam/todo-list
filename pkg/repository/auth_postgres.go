package repository

import (
	"fmt"
	todolist "github.com/TauAdam/todo-list"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user todolist.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AuthPostgres) GetUser(username, password string) (todolist.User, error) {
	var user todolist.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
