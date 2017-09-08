package main

type TransformedUsers struct {
	ID          int `json:"id"`
	IsValidated bool
	Version     int
	Email       string
	Username    string
}
