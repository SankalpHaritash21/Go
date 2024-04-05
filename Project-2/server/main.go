package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	// Routes
	r.GET("/todos", getTodos)
	r.POST("/todos", addTodo)

	// Run server
	r.Run(":5000")
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}
