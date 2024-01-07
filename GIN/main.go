package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json;"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// todos is a slice of todo
var todos = []todo{
	{ID: "1", Item: "Learn Golang with gin", Completed: false},
	{ID: "2", Item: "Learn Gin", Completed: false},
	{ID: "3", Item: "Learn React", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodos(c *gin.Context) {
	var newTodo todo

	// check id should be unique
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	for _, todo := range todos {
		if todo.ID == newTodo.ID || todo.Item == newTodo.Item {
			c.JSON(http.StatusBadRequest, gin.H{"message": "ID or Item already exists"})
			return
		}
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func findTodoById(id string) (todo, bool) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, true
		}
	}
	return todo{}, false
}

func deleteTodoById(id string) bool {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}

func updateTodoById(id string, newTodo todo) bool {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = newTodo
			return true
		}
	}
	return false
}

func findTodoByIdWrapper(c *gin.Context) {
	id := c.Param("id")
	todo, ok := findTodoById(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func deleteTodoByIdWrapper(c *gin.Context) {
	id := c.Param("id")
	ok := deleteTodoById(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func updateTodoByIdWrapper(c *gin.Context) {
	id := c.Param("id")
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	ok := updateTodoById(id, newTodo)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func wrongRoutes(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func main() {
	router := gin.Default()

	router.NoRoute(wrongRoutes)

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.GET("/todos/:id", findTodoByIdWrapper)
	router.DELETE("/todos/:id", deleteTodoByIdWrapper)
	router.PUT("/todos/:id", updateTodoByIdWrapper)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
