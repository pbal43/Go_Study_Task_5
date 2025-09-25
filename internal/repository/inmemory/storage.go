package inmemory

import (
	"toDoList/internal/domain/task/task_models"
	"toDoList/internal/domain/user/user_models"
)

type Storage struct {
	users map[string]user_models.User
	tasks map[string]task_models.Task
}

func NewInMemoryStorage() *Storage {
	return &Storage{
		users: make(map[string]user_models.User),
		tasks: make(map[string]task_models.Task),
	}
}
