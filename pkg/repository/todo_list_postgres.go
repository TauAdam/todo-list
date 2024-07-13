package repository

import (
	"fmt"
	todolist "github.com/TauAdam/todo-list"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}
func (r *TodoListPostgres) Create(userId int, list todolist.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2)", userListsTable)
	_, err = tx.Exec(createUserListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *TodoListPostgres) GetAll(userId int) ([]todolist.TodoList, error) {
	var lists []todolist.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, userListsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListPostgres) GetById(userId int, id int) (todolist.TodoList, error) {
	var list todolist.TodoList
	// also check that the list belongs to the user
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND tl.id = $2", todoListsTable, userListsTable)
	err := r.db.Get(&list, query, userId, id)
	return list, err
}
func (r *TodoListPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND tl.id = $2", todoListsTable, userListsTable)
	_, err := r.db.Exec(query, userId, id)
	return err
}
