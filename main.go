package main

import (
	yaag_gin "github.com/betacraft/yaag/gin"
	"github.com/betacraft/yaag/yaag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strconv"
)

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gomain sslmode=disable password=")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {
	yaag.Init(&yaag.Config{On: true, DocTitle: "Authentication", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})

	//Migrate the schema
	db := Database()
	db.AutoMigrate(&Todo{})

	router := gin.Default()
	router.Use(yaag_gin.Document())

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", CreateTodo)
		v1.GET("/", FetchAllTodo)
		v1.GET("/:id", FetchSingleTodo)
		v1.PUT("/:id", UpdateTodo)
		v1.DELETE("/:id", DeleteTodo)
	}
	router.Run()

}

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

type TransformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := Todo{Title: c.PostForm("title"), Completed: completed}
	db := Database()
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

func FetchAllTodo(c *gin.Context) {
	var todos []Todo
	var _todos []TransformedTodo

	db := Database()
	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

func FetchSingleTodo(c *gin.Context) {
	var todo Todo
	todoId := c.Param("id")

	db := Database()
	db.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

func UpdateTodo(c *gin.Context) {
	var todo Todo
	todoId := c.Param("id")
	db := Database()
	db.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

func DeleteTodo(c *gin.Context) {
	var todo Todo
	todoId := c.Param("id")
	db := Database()
	db.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
