package repository

import (
	"slices"
	task2 "toDoList/internal/domain/task/task_errors"
	"toDoList/internal/domain/task/task_models"
)

var tasksRepository []task_models.Task

func GetAllTasksFromDB() []task_models.Task {
	return tasksRepository
}

func GetOneTaskByID(id string) (*task_models.Task, int, error) { // как возвращать опционалы или не возвращать инт?
	for i := range tasksRepository {
		if tasksRepository[i].ID == id {
			return &tasksRepository[i], i, nil
		}
	}
	return nil, -1, task2.FoundNothingErr
}

func AddTask(task task_models.Task) {
	tasksRepository = append(tasksRepository, task)
}

func updateTaskToNewTask(newTask task_models.Task, oldTask *task_models.Task) {
	oldTask.Status = newTask.Status
	oldTask.Description = newTask.Description
	oldTask.Status = newTask.Status
}

func UpdateExistedTask(task task_models.Task) error {
	foundedTask, _, err := GetOneTaskByID(task.ID)
	if err != nil {
		return task2.FoundNothingErr
	}
	updateTaskToNewTask(task, foundedTask)
	return nil
}

func DeleteExistedTask(id string) error {
	_, index, err := GetOneTaskByID(id)
	if err != nil {
		return task2.FoundNothingErr
	}
	tasksRepository = slices.Delete(tasksRepository, index, index+1)
	return nil
}
