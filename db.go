package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gomain sslmode=disable password=")
	if err != nil {
		panic("failed to connect database")
	}
	// Make sure to not handle structs as plural names
	db.SingularTable(true)

	// Comment to disable logger
	db.LogMode(true)
	return db
}
