package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toDoList/internal/domain/models"
	"toDoList/internal/service"
)

// обрабатываем для вывода, возвращаем респонсы с ошибками и проч.

func getTasks(ctx *gin.Context) {
	tasks := service.GetAllTasksInMap()
	if len(tasks) != 0 {
		ctx.JSON(http.StatusOK, tasks)
	} else {
		ctx.JSON(http.StatusOK, "Task list is empty")
	}
}

func getTaskByID(ctx *gin.Context) {
	taskID := ctx.Param("id")
	foundedTask, err := service.GetTaskByID(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, foundedTask)
}

func createTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	taskID, err := service.CreateNewTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"TaskID": taskID})
}

func updateTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	var newTask models.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newTask.ID = taskID
	taskID, err := service.UpdateTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"TaskID": taskID})
}

func deleteTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	if err := service.DeleteTaskByID(taskID); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Task was deleted")
}
