package main

import (
	"time"
)

type Todo struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

type User struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsValidated bool
	IsDeleted   bool
	Version     string
	Email       string
	Password    string
	Username    string
}

type TimeToken struct {
	ID        string
	UserId    string
	TokenType string
	CreatedAt time.Time
	UsedAt    time.Time
}
