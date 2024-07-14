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
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return itemId, nil
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todolist.TodoItem, error) {
	var items []todolist.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2", todoItemsTable, listItemsTable, userListsTable)
	err := r.db.Select(&items, query, listId, userId)
	if err != nil {
		return nil, err
	}
	return items, nil
}
