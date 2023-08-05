package main

import (
	"net/http"
	"todoApp/models"

	"github.com/gin-gonic/gin"
)

var todos []models.Todo

func main() {
	r := gin.Default()

	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	/*r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)*/

	r.Run(":8080")
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}
