package server

import (
	"net/http"
	"toDoList/internal/domain/task/task_models"
	"toDoList/internal/service/task_service"

	"github.com/gin-gonic/gin"
)

// обрабатываем для вывода, возвращаем респонсы с ошибками и проч.

func (srv *ToDoListApi) getTasks(ctx *gin.Context) {
	tasks := task_service.GetAllTasksInMap()
	if len(tasks) != 0 {
		ctx.JSON(http.StatusOK, tasks)
	} else {
		ctx.JSON(http.StatusOK, "Task list is empty")
	}
}

func (srv *ToDoListApi) getTaskByID(ctx *gin.Context) {
	taskID := ctx.Param("id")
	foundedTask, err := task_service.GetTaskByID(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, foundedTask)
}

func (srv *ToDoListApi) createTask(ctx *gin.Context) {
	var newTask task_models.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	taskID, err := task_service.CreateNewTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"TaskID": taskID})
}

func (srv *ToDoListApi) updateTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	var newTask task_models.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newTask.ID = taskID
	taskID, err := task_service.UpdateTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"TaskID": taskID})
}

func (srv *ToDoListApi) deleteTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	if err := task_service.DeleteTaskByID(taskID); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Task was deleted")
}
