package main

import (
	"net/http"
	"strconv"
	"todoApp/models"

	"github.com/gin-gonic/gin"
)

var todos = make(map[int]models.Todo)

func main() {
	r := gin.Default()

	r.Static("/static", "/template")
	r.LoadHTMLGlob("template/*")
	r.GET("/", getTodos)
	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.POST("/update/:id", updateTodo)
	r.GET("/complete/:id", complete)
	r.GET("/delete/:id", deleteTodo)
	r.GET("/update/:id", todoUpdate1)

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
	todos[todo.ID] = todo
	c.HTML(http.StatusOK, "home.html", gin.H{"todo": todos})
}

func todoUpdate1(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}
	c.HTML(http.StatusAccepted, "update.html", gin.H{"todo": todos[index]})

}

func updateTodo(c *gin.Context) {
	var newTodo models.Todo
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}
	newTodo.ID = index
	newTodo.Title = c.PostForm("title")
	newTodo.Description = c.PostForm("description")
	newTodo.Status = c.PostForm("status")

	todos[index] = newTodo
	c.HTML(http.StatusAccepted, "home.html", gin.H{"message": "Task details Updated", "todo": todos})
}

func complete(c *gin.Context) {
	var newTodo models.Todo
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}

	newTodo.ID = index
	newTodo.Title = todos[index].Title
	newTodo.Description = todos[index].Description
	newTodo.Status = "Done"
	todos[index] = newTodo
	c.HTML(http.StatusAccepted, "home.html", gin.H{"message": "Task details Updated", "todo": todos})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}
	delete(todos, index)
	c.HTML(http.StatusAccepted, "home.html", gin.H{"message": "Task deleted", "todo": todos})
}
