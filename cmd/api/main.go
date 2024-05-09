package main

import (
	"github.com/eduardoraider/go-todo-gin/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/todos", handlers.GetTodosHandler)
	r.GET("/todos/:id", handlers.GetTodoByIdHandler)
	r.POST("/todos", handlers.CreateTodoHandler)
	r.PUT("/todos/:id", handlers.UpdateTodoHandler)
	r.DELETE("/todos/:id", handlers.DeleteTodoHandler)

	r.Run(":8080")
}
