package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"toDoList/internal/domain/errors"
	"toDoList/internal/domain/models"
	"toDoList/internal/repository"
)

func GetAllTasksInMap() map[string]any {
	allTasks := repository.GetAllTasksFromDB()
	allTasksMapped := make(map[string]any)
	for _, task := range allTasks {
		allTasksMapped[task.ID] = struct {
			Title       string            `json:"title"`
			Description string            `json:"description"`
			Status      models.TaskStatus `json:"status"`
		}{
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		}
	}
	return allTasksMapped
}

func GetTaskByID(taskID string) (models.Task, error) {
	if taskID == "" {
		return models.Task{}, errors.EpmtyStringErr
	}
	task, _, err := repository.GetOneTaskByID(taskID)
	if err != nil {
		return models.Task{}, err
	}
	return *task, nil
}

func CreateNewTask(task models.Task) (string, error) {
	task.ID = uuid.New().String()
	validatorForTask := validator.New()
	if err := validatorForTask.Struct(task); err != nil {
		return "", err
	}
	if !task.Status.IsValid() {
		return "", errors.WrongStatusErr
	}
	repository.AddTask(task)
	return task.ID, nil
}

func UpdateTask(task models.Task) (string, error) {
	validatorForTask := validator.New()
	if err := validatorForTask.Struct(task); err != nil {
		return "", err
	}
	if !task.Status.IsValid() {
		return "", errors.WrongStatusErr
	}
	if err := repository.UpdateExistedTask(task); err != nil {
		return "", err
	}
	return task.ID, nil
}

func DeleteTaskByID(taskID string) error {
	if err := repository.DeleteExistedTask(taskID); err != nil {
		return err
	}
	return nil
}
