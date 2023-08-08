package main

import (
	"net/http"
	"strconv"
	"todoApp/models"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{}

func main() {
	r := gin.Default()

	r.Static("/static", "/template")
	r.LoadHTMLGlob("template/*")
	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.GET("/complete/:id", complete)
	r.GET("/delete/:id", deleteTodo)

	r.Run(":8080")
}

func getTodos(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"todo": todos})
	c.Redirect(http.StatusPermanentRedirect, "/todos")
}

func createTodo(c *gin.Context) {
	var todo models.Todo
	todo.Title = c.PostForm("title")
	todo.Description = c.PostForm("description")
	todo.Status = c.PostForm("status")

	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	c.HTML(http.StatusOK, "home.html", nil)
}

func updateTodo(c *gin.Context) {
	var newTodo models.Todo
	c.ShouldBindJSON(&newTodo)
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}

	todos[index].Title = newTodo.Title
	todos[index].Description = newTodo.Description
	todos[index].Status = newTodo.Status
	c.JSON(http.StatusAccepted, gin.H{"message": "updated"})
}

func complete(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}

	todos[index].Status = "Completed"
	c.JSON(http.StatusAccepted, gin.H{"message": "updated"})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}
	todos = append(todos[:index], todos[index+1:]...)
	c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
}
