package main

import (
	"time"
)

type Todo struct {
	ID        uint       `db:"id"`
	Title     string     `db:"title"`
	Completed int        `db:"completed"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type Users struct {
	ID          int
	IsValidated bool
	IsDeleted   bool
	Version     int
	Email       string
	Password    string
	Username    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type TimeToken struct {
	ID        string
	UserId    string
	TokenType string
	CreatedAt time.Time
	UsedAt    *time.Time
}
