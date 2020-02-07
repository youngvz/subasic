package models

type User struct {
	UserID int `json:"userId"`
	Name string `json:"name"`
	Email string `json:"email"`
	Admin bool `json:"admin"`
	GroupID *int `json:"groupId"`
}