package main

import (
	yaag_gin "github.com/betacraft/yaag/gin"
	"github.com/betacraft/yaag/yaag"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	yaag.Init(&yaag.Config{On: true, DocTitle: "Authentication", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})

	//Migrate the schema
	db := Database()
	db.AutoMigrate(&Todo{}, &User{})

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
