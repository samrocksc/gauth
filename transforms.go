package main

type TransformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TransformedUsers struct {
	ID          int `json:"id"`
	IsValidated bool
	Version     int
	Email       string
	Username    string
}
