package task_service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	task2 "toDoList/internal/domain/task/task_errors"
	"toDoList/internal/domain/task/task_models"
	"toDoList/internal/repository"
)

func GetAllTasksInMap() map[string]any {
	allTasks := repository.GetAllTasksFromDB()
	allTasksMapped := make(map[string]any)
	for _, task := range allTasks {
		allTasksMapped[task.ID] = struct {
			Title       string          `json:"title"`
			Description string          `json:"description"`
			Status      task.TaskStatus `json:"status"`
		}{
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		}
	}
	return allTasksMapped
}

func GetTaskByID(taskID string) (task_models.Task, error) {
	if taskID == "" {
		return task_models.Task{}, task2.EpmtyStringErr
	}
	task, _, err := repository.GetOneTaskByID(taskID)
	if err != nil {
		return task.Task{}, err
	}
	return *task, nil
}

func CreateNewTask(task task_models.Task) (string, error) {
	task.ID = uuid.New().String()
	validatorForTask := validator.New()
	if err := validatorForTask.Struct(task); err != nil {
		return "", err
	}
	if !task.Status.IsValid() {
		return "", task2.WrongStatusErr
	}
	repository.AddTask(task)
	return task.ID, nil
}

func UpdateTask(task task_models.Task) (string, error) {
	validatorForTask := validator.New()
	if err := validatorForTask.Struct(task); err != nil {
		return "", err
	}
	if !task.Status.IsValid() {
		return "", task2.WrongStatusErr
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
