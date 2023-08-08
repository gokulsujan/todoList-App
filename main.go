package main

import (
	"net/http"
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

	r.Run(":8080")
}

func getTodos(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"todo": todos})
	//c.JSON(http.StatusOK, gin.H{"todo": todos})
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
