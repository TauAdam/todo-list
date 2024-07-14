package repository

import (
	"fmt"
	todolist "github.com/TauAdam/todo-list"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db}
}
func (r *TodoItemPostgres) Create(listId int, item todolist.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listItemsTable)

	if _, err := tx.Exec(createListItemQuery, listId, itemId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}
