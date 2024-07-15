package repository

import (
	"fmt"
	todolist "github.com/TauAdam/todo-list"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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
func (r *TodoItemPostgres) GetById(userId, itemId int) (todolist.TodoItem, error) {
	var item todolist.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2", todoItemsTable, listItemsTable, userListsTable)
	err := r.db.Get(&item, query, itemId, userId)
	if err != nil {
		return item, err
	}
	return item, nil
}
func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2", todoItemsTable, listItemsTable, userListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}
func (r *TodoItemPostgres) Update(userId, itemId int, input todolist.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d", todoItemsTable, setQuery, listItemsTable, userListsTable, argId, argId+1)
	args = append(args, userId, itemId)
	logrus.Debug("update query: ", query)
	logrus.Debug("update args: ", args)
	_, err := r.db.Exec(query, args...)
	return err
}
