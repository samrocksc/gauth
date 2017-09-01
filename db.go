package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gomain sslmode=disable password=postgres")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
