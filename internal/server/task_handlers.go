package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toDoList/internal/domain/task/task_models"
	"toDoList/internal/service/task"
)

// обрабатываем для вывода, возвращаем респонсы с ошибками и проч.

func GetTasks(ctx *gin.Context) {
	tasks := task.GetAllTasksInMap()
	if len(tasks) != 0 {
		ctx.JSON(http.StatusOK, tasks)
	} else {
		ctx.JSON(http.StatusOK, "Task list is empty")
	}
}

func GetTaskByID(ctx *gin.Context) {
	taskID := ctx.Param("id")
	foundedTask, err := task.GetTaskByID(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, foundedTask)
}

func CreateTask(ctx *gin.Context) {
	var newTask task_models.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	taskID, err := task.CreateNewTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"TaskID": taskID})
}

func UpdateTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	var newTask task_models.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newTask.ID = taskID
	taskID, err := task.UpdateTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"TaskID": taskID})
}

func DeleteTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	if err := task.DeleteTaskByID(taskID); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Task was deleted")
}
