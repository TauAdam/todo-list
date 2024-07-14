package service

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/repository"
)

type TodoItemService struct {
	todoRepo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(todoRepo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{todoRepo, listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todolist.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// if list does not exist or does not belong to the user
		return 0, err
	}
	return s.todoRepo.Create(listId, item)
}
func (s *TodoItemService) GetAll(userId, listId int) ([]todolist.TodoItem, error) {
	return s.todoRepo.GetAll(userId, listId)
}
func (s *TodoItemService) GetById(userId, itemId int) (todolist.TodoItem, error) {
	return s.todoRepo.GetById(userId, itemId)
}
