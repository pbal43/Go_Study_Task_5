package repository

import (
	"slices"
	"toDoList/internal/domain/errors"
	"toDoList/internal/domain/models"
)

var tasksRepository []models.Task

func GetAllTasksFromDB() []models.Task {
	return tasksRepository
}

func GetOneTaskByID(id string) (*models.Task, int, error) { // как возвращать опционалы или не возвращать инт?
	for i := range tasksRepository {
		if tasksRepository[i].ID == id {
			return &tasksRepository[i], i, nil
		}
	}
	return nil, -1, errors.FoundNothingErr
}

func AddTask(task models.Task) {
	tasksRepository = append(tasksRepository, task)
}

func updateTaskToNewTask(newTask models.Task, oldTask *models.Task) {
	oldTask.Status = newTask.Status
	oldTask.Description = newTask.Description
	oldTask.Status = newTask.Status
}

func UpdateExistedTask(task models.Task) error {
	foundedTask, _, err := GetOneTaskByID(task.ID)
	if err != nil {
		return errors.FoundNothingErr
	}
	updateTaskToNewTask(task, foundedTask)
	return nil
}

func DeleteExistedTask(id string) error {
	_, index, err := GetOneTaskByID(id)
	if err != nil {
		return errors.FoundNothingErr
	}
	tasksRepository = slices.Delete(tasksRepository, index, index+1)
	return nil
}
