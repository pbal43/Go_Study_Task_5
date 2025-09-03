package server

import (
	"github.com/gin-gonic/gin"
)

func ToDoListServer() {
	router := gin.Default()
	toDoList := router.Group("/todolist")
	{
		toDoList.GET("/tasks", getTasks)
		toDoList.GET("/tasks/:id", getTaskByID)
		toDoList.POST("/tasks", createTask)
		toDoList.PUT("/tasks/:id", updateTask)
		toDoList.DELETE("/tasks/:id", deleteTask)
	}
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
