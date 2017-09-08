// models.go - models of structures directly comparable to their counterparts in the migration files
package main

import (
	"time"
)

// Users is a struct that is directly comparable to its migration counterpart
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

// TimeToken is a struct that is directly comparable to its migration counterpart
type TimeToken struct {
	ID        string
	UserId    string
	TokenType string
	CreatedAt time.Time
	UsedAt    *time.Time
}
