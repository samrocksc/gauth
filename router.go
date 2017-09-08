package main

import (
	yaag_gin "github.com/betacraft/yaag/gin"
	"github.com/betacraft/yaag/yaag"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gin-gonic/gin.v1"
)

func Router() {
	yaag.Init(&yaag.Config{On: true, DocTitle: "Authentication", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})

	router := gin.Default()
	router.Use(yaag_gin.Document())

	v1 := router.Group("/auth")
	{
		v1.POST("/todo", CreateTodo)
		v1.GET("/todo", FetchAllTodo)
		v1.GET("/todo/:id", FetchSingleTodo)
		v1.PUT("/todo/:id", UpdateTodo)
		v1.DELETE("/todo/:id", DeleteTodo)
		v1.POST("/user", CreateUser)
		v1.GET("/user", GetAllUsers)
	}
	router.Run()
}
