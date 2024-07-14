package service

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func (s *TodoListService) GetById(userId, id int) (todolist.TodoList, error) {
	return s.repo.GetById(userId, id)
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todolist.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
func (s *TodoListService) GetAll(userId int) ([]todolist.TodoList, error) {
	return s.repo.GetAll(userId)
}
func (s *TodoListService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
func (s *TodoListService) Update(userId, id int, input todolist.UpdateListInput) error {
	return s.repo.Update(userId, id, input)
}
