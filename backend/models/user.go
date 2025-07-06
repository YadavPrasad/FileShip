package models

type User struct {
	ID unit `json : "id"`
	Username string `json : "username"`
	Password string `json : "-"`
}