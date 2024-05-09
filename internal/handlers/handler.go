package handlers

import (
	"github.com/eduardoraider/go-todo-gin/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTodosHandler(c *gin.Context) {
	todos := entity.GetTodos()
	c.JSON(http.StatusOK, todos)
}

func GetTodoByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := entity.GetTodoByID(id)
	c.JSON(http.StatusOK, todo)
}

func CreateTodoHandler(c *gin.Context) {
	var todo entity.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity.CreateTodo(todo)

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo entity.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	entity.UpdateTodo(id, todo)

	c.JSON(http.StatusOK, todo)
}

func DeleteTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity.DeleteTodo(id)

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})

}
